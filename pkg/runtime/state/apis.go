package state

import (
	"context"
	"encoding/json"
	"sort"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/pkg/errors"
	"github.com/tkeel-io/collectjs"
	"github.com/tkeel-io/collectjs/pkg/json/jsonparser"
	"github.com/tkeel-io/core/pkg/constraint"
	xerrors "github.com/tkeel-io/core/pkg/errors"
	zfield "github.com/tkeel-io/core/pkg/logger"
	"github.com/tkeel-io/core/pkg/mapper"
	"github.com/tkeel-io/core/pkg/repository/dao"
	"github.com/tkeel-io/core/pkg/runtime/message"
	"github.com/tkeel-io/core/pkg/types"
	"github.com/tkeel-io/core/pkg/util"
	xjson "github.com/tkeel-io/core/pkg/util/json"
	"github.com/tkeel-io/kit/log"
	"go.uber.org/zap"
)

type APIID string

func (a APIID) String() string {
	return string(a)
}

const (
	APICreateEntity        APIID = "core.apis.Entity.Create"
	APIUpdateEntity        APIID = "core.apis.Entity.Update"
	APIGetEntity           APIID = "core.apis.Entity.Get"
	APIDeleteEntity        APIID = "core.apis.Entity.Delete"
	APIUpdataEntityProps   APIID = "core.apis.Entity.Props.Update"
	APIPatchEntityProps    APIID = "core.apis.Entity.Props.Patch"
	APIGetEntityProps      APIID = "core.apis.Entity.Props.Get"
	APIUpdataEntityConfigs APIID = "core.apis.Entity.Configs.Update"
	APIPatchEntityConfigs  APIID = "core.apis.Entity.Configs.Patch"
	APIGetEntityConfigs    APIID = "core.apis.Entity.Configs.Get"
)

type APIHandler func(context.Context, message.Context) ([]WatchKey, error)

// call core.APIs.
func (s *statem) callAPIs(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	log.Debug("call core.APIs", zfield.Header(msgCtx.Attributes()))

	apiID := APIID(msgCtx.Get(message.ExtAPIIdentify))
	switch apiID {
	case APICreateEntity:
		return s.cbCreateEntity(ctx, msgCtx)
	case APIUpdateEntity:
		return s.cbUpdateEntity(ctx, msgCtx)
	case APIGetEntity:
		return s.cbGetEntity(ctx, msgCtx)
	case APIDeleteEntity:
		return s.cbDeleteEntity(ctx, msgCtx)
	case APIUpdataEntityProps:
		return s.cbUpdateEntityProps(ctx, msgCtx)
	case APIPatchEntityProps:
		return s.cbPatchEntityProps(ctx, msgCtx)
	case APIGetEntityProps:
		return s.cbGetEntityProps(ctx, msgCtx)
	case APIUpdataEntityConfigs:
		return s.cbUpdateEntityConfigs(ctx, msgCtx)
	case APIPatchEntityConfigs:
		return s.cbPatchEntityConfigs(ctx, msgCtx)
	case APIGetEntityConfigs:
		return s.cbGetEntityConfigs(ctx, msgCtx)
	default:
		// nerver.
	}

	return []WatchKey{}, nil
}

func (s *statem) makeEvent() cloudevents.Event {
	ev := cloudevents.NewEvent()
	ev.SetID(util.UUID())
	ev.SetSource("core.runtime")
	ev.SetType(message.MessageTypeAPIRespond.String())
	ev.SetExtension(message.ExtEntityID, s.ID)
	ev.SetExtension(message.ExtEntityType, s.Type)
	ev.SetExtension(message.ExtEntityOwner, s.Owner)
	ev.SetExtension(message.ExtEntitySource, s.Source)
	ev.SetExtension(message.ExtAPIRespStatus, types.StatusOK.String())
	return ev
}

func (s *statem) setEventPayload(ev *cloudevents.Event, reqID string) error {
	var (
		err   error
		bytes []byte
	)

	if bytes, err = dao.GetEntityCodec().Encode(&s.Entity); nil != err {
		log.Error("set response", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
		ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		return nil
	}

	if err = ev.SetData(bytes); nil != err {
		log.Error("set response event payload", zfield.Eid(s.ID), zap.Error(err))
		ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
		ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
	} else if err = ev.Validate(); nil != err {
		log.Error("validate response", zfield.Eid(s.ID), zap.Error(err))
		ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
		ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
	}

	return errors.Wrap(err, "set event payload")
}

func (s *statem) cbCreateEntity(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var (
		err   error
		reqEn dao.Entity
	)

	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version > 0 {
		err = xerrors.ErrEntityAleadyExists
		log.Error("state machine already exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine already exists")
	}

	// decode request.
	if err = dao.GetEntityCodec().Decode(msgCtx.Message(), &reqEn); nil != err {
		log.Error("decode core api request", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "decode core api request")
	}

	// initialize.
	s.LastTime = util.UnixMilli()
	s.Properties = reqEn.Properties
	s.TemplateID = reqEn.TemplateID

	// sync template.
	if reqEn.TemplateID != "" {
		var tempEn = &dao.Entity{ID: reqEn.TemplateID}
		if tempEn, err = s.Repo().GetEntity(ctx, tempEn); nil != err {
			log.Error("pull template entity", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
			return nil, errors.Wrap(err, "pull template entity")
		}

		s.ConfigBytes = tempEn.ConfigBytes
	}

	// parse configs.
	s.reparseConfig()

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set response event payload")
	}

	s.Version++

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	s.flush(ctx)

	return nil, nil
}

func (s *statem) cbUpdateEntity(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	panic("implement me")
}

func (s *statem) cbGetEntity(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var err error
	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}
	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	return []WatchKey{}, errors.Wrap(err, "")
}

func (s *statem) cbDeleteEntity(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	s.status = SMStatusDeleted

	var err error
	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	return []WatchKey{}, errors.Wrap(err, "")
}

func (s *statem) cbUpdateEntityProps(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var (
		err   error
		reqEn dao.Entity
	)

	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}

	// decode request.
	if err = dao.GetEntityCodec().Decode(msgCtx.Message(), &reqEn); nil != err {
		log.Error("decode core api request", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "decode core api request")
	}

	// update state properties.
	stateIns := State{ID: s.ID, Props: s.Properties}

	watchKeys := make([]mapper.WatchKey, 0)
	for key, val := range reqEn.Properties {
		if _, err = stateIns.Patch(xjson.OpReplace, key, []byte(val.String())); nil != err {
			log.Error("upsert state property", zfield.ID(s.ID), zfield.PK(key), zap.Error(err))
		} else {
			watchKeys = append(watchKeys, mapper.WatchKey{EntityID: s.ID, PropertyKey: key})
		}
	}

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	s.flush(ctx)

	return watchKeys, nil
}

func (s *statem) cbPatchEntityProps(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var (
		err error
		pds []PatchData
	)

	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}

	// decode request.
	if pds, err = GetPatchCodec().Decode(msgCtx.Message()); nil != err {
		log.Error("decode core api request", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "decode core api request")
	}

	// update state properties.
	stateIns := State{ID: s.ID, Props: s.Properties}
	watchKeys := make([]mapper.WatchKey, 0)
	for index := range pds {
		valBytes, _ := pds[index].Value.([]byte)
		op := xjson.NewPatchOp(pds[index].Operator)
		if _, err = stateIns.Patch(op, pds[index].Path, valBytes); nil != err {
			log.Error("upsert state property", zfield.ID(s.ID), zfield.PK(pds[index].Path), zap.Error(err))
		} else {
			watchKeys = append(watchKeys, mapper.WatchKey{EntityID: s.ID, PropertyKey: pds[index].Path})
		}
	}

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	s.flush(ctx)
	return watchKeys, nil
}

func (s *statem) cbGetEntityProps(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var err error
	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}
	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	return []WatchKey{}, nil
}

func (s *statem) cbUpdateEntityConfigs(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var (
		err   error
		reqEn dao.Entity
	)

	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}

	// decode request.
	if err = dao.GetEntityCodec().Decode(msgCtx.Message(), &reqEn); nil != err {
		log.Error("decode core api request", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "decode core api request")
	}

	// update configs.
	var configBytes = []byte("{}")
	collectjs.ForEach(reqEn.ConfigBytes, jsonparser.Object,
		func(key, value []byte, dataType jsonparser.ValueType) {
			propertyKey := string(key)
			if configBytes, err = collectjs.Set(configBytes, propertyKey, value); nil != err {
				log.Error("call core.APIs.PatchConfigs patch add", zap.Error(err))
				err = errors.Wrap(err, "patch config")
			}
		})

	if nil != err {
		log.Error("call core.APIs.PatchConfigs patch configs",
			zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "update state configs")
	}

	s.ConfigBytes = configBytes
	// reparse state configs.
	s.reparseConfig()

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	s.flush(ctx)
	return []WatchKey{}, nil
}

func (s *statem) cbPatchEntityConfigs(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	var (
		err   error
		bytes []byte
		pds   []PatchData
	)

	// create event.
	ev := s.makeEvent()
	reqID := msgCtx.Get(message.ExtAPIRequestID)
	ev.SetExtension(message.ExtAPIRequestID, reqID)
	ev.SetExtension(message.ExtCallback, msgCtx.Get(message.ExtCallback))

	defer func() {
		if nil != err {
			ev.SetExtension(message.ExtAPIRespStatus, types.StatusError.String())
			ev.SetExtension(message.ExtAPIRespErrCode, err.Error())
		}
		if innerErr := s.dispatcher.Dispatch(ctx, ev); nil != err {
			log.Error("diispatch event", zap.Error(innerErr),
				zfield.Eid(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))
		}
	}()

	// check version.
	if s.Version == 0 {
		err = xerrors.ErrEntityNotFound
		log.Error("state machine not exists", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "state machine not exists")
	}

	// decode request.
	if pds, err = GetPatchCodec().Decode(msgCtx.Message()); nil != err {
		log.Error("decode core api request", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "set event payload")
	}

	for _, pd := range pds {
		switch xjson.NewPatchOp(pd.Operator) {
		case xjson.OpAdd:
			if bytes, err = collectjs.Append(s.ConfigBytes, pd.Path, pd.Value.([]byte)); nil != err {
				log.Error("call core.APIs.PatchConfigs patch add",
					zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
				continue
			}
			s.ConfigBytes = bytes
		case xjson.OpRemove:
			s.ConfigBytes = collectjs.Del(s.ConfigBytes, pd.Path)
		case xjson.OpReplace:
			if valBytes, ok := pd.Value.([]byte); ok {
				if bytes, err = collectjs.Set(s.ConfigBytes, pd.Path, valBytes); nil != err {
					log.Error("call core.APIs.PatchConfigs patch replace", zap.Error(err))
					return nil, errors.Wrap(err, "patch replace")
				}
				s.ConfigBytes = bytes
			}
			log.Error("assert patch data value", zfield.Eid(s.ID), zfield.ReqID(reqID))
		}
	}

	s.reparseConfig()

	// set response.
	if err = s.setEventPayload(&ev, reqID); nil != err {
		log.Error("set event payload", zap.Error(err), zfield.Eid(s.ID), zfield.ReqID(reqID))
		return nil, errors.Wrap(err, "patch replace")
	}

	log.Debug("core.APIS callback", zfield.ID(s.ID), zfield.ReqID(msgCtx.Get(message.ExtAPIRequestID)))

	s.flush(ctx)
	return []WatchKey{}, nil
}

func (s *statem) cbGetEntityConfigs(ctx context.Context, msgCtx message.Context) ([]WatchKey, error) {
	panic("implement me")
}

func (s *statem) reparseConfig() {
	var err error
	// parse state config again.
	configs := make(map[string]interface{})
	if err = json.Unmarshal(s.ConfigBytes, &configs); nil != err {
		log.Error("json unmarshal", zap.Error(err), zap.String("configs", string(s.ConfigBytes)))
	}

	var cfg constraint.Config
	cfgs := make(map[string]constraint.Config)
	for key, val := range configs {
		if cfg, err = constraint.ParseConfigFrom(val); nil != err {
			log.Error("parse configs", zap.Error(err))
			continue
		}
		cfgs[key] = cfg
	}

	// reset state machine configs.
	s.constraints = make(map[string]*constraint.Constraint)
	s.searchConstraints = make(sort.StringSlice, 0)
	s.tseriesConstraints = make(sort.StringSlice, 0)

	for _, cfg = range cfgs {
		if ct := constraint.NewConstraintsFrom(cfg); nil != ct {
			s.constraints[ct.ID] = ct
			// generate search indexes.
			if searchIndexes := ct.GenEnabledIndexes(constraint.EnabledFlagSearch); len(searchIndexes) > 0 {
				s.searchConstraints = util.SliceAppend(s.searchConstraints, searchIndexes)
			}

			// generate time-series indexes.
			if tseriesIndexes := ct.GenEnabledIndexes(constraint.EnabledFlagTimeSeries); len(tseriesIndexes) > 0 {
				s.tseriesConstraints = util.SliceAppend(s.tseriesConstraints, tseriesIndexes)
			}
		}
	}
}
