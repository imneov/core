/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runtime

import (
	"context"
	"github.com/stretchr/testify/assert"
	v1 "github.com/tkeel-io/core/api/core/v1"
	"github.com/tkeel-io/core/pkg/placement"
	"github.com/tkeel-io/core/pkg/repository"
	"github.com/tkeel-io/core/pkg/util/json"
	"github.com/tkeel-io/core/pkg/util/path"
	"testing"

	"github.com/tkeel-io/tdtl"
)

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	proto "github.com/tkeel-io/core/api/core/v1"
// 	"github.com/tkeel-io/core/pkg/util"
// )

// func TestRuntime_HandleEvent(t *testing.T) {
// 	ctx := context.Background()
// 	ev := &proto.ProtoEvent{
// 		Id:        util.UUID("ev"),
// 		RawData:   []byte(`{}`),
// 		Timestamp: time.Now().UnixNano(),
// 		Metadata:  map[string]string{},
// 	}

// 	cc := NewRuntime(ctx, "core-1")
// 	ret, err := cc.HandleEvent(ctx, ev)
// 	assert.NoError(t, err, "err is %v", err)
// 	byt, err := cc.entities["entity1"].Raw()
// 	t.Log(cc, string(byt))
// 	t.Log(string(ret.State))
// }

// func TestRuntime_OpEntity_HandleEvent(t *testing.T) {
// 	type args struct {
// 	}
// 	tests := []struct {
// 		name       string
// 		typ        EntityEventType
// 		path       string
// 		eventValue []byte
// 		wantErr    bool
// 		want       string
// 	}{
// 		{"1", OpEntityPropsUpdata, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{"a":{"b":{"c":"abc"}}},"Scheme":{}}`},
// 		//{"2", OpEntityPropsPatch, "a.b.c", []byte(`"abc"`), false, `{"Property":{"a":{"b":{"c":"abc"}}},"Scheme":{}}`},
// 		{"3", OpEntityPropsGet, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{}}`},
// 		{"4", OpEntityConfigsUpdata, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{"a":{"b":{"c":"abc"}}}}`},
// 		//{"5", OpEntityConfigsPatch, "a.b.c", []byte(`"abc"`), false, `{"Property":{},"Scheme":{"a":{"b":{"c":"abc"}}}}`},
// 		{"6", OpEntityConfigsGet, "a.b.c", []byte(`{
// 			Id:     "device123",
// 			Type:   "DEVICE",
// 			Owner:  "tomas",
// 			Source: "CORE-SDK",
// 			Properties: map[string]interface{}{
// 			"temp": 20,
// 		},`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{}}`},
// 	}
// 	ctx := context.Background()
// 	for _, tt := range tests {
// 		cc := newRuntime()
// 		t.Run(tt.name, func(t *testing.T) {
// 			ev := &proto.ProtoEvent{}
// 			got, err := cc.HandleEvent(ctx, ev)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("HandleEvent() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(string(got.State), tt.want) {
// 				t.Errorf("HandleEvent() \ngot = %v, \nwant  %v", string(got.State), tt.want)
// 			}
// 		})
// 	}
// }

// func newRuntime() *Runtime {
// 	cc := NewRuntime(context.Background(), "core-1")
// 	return cc
// }

func TestTTDL(t *testing.T) {
	cc := tdtl.New([]byte(`{"a": 20}`))
	cc.Set("", tdtl.New("{}"))
	t.Log(cc.Error())
	t.Log(string(cc.Raw()))
	t.Log(cc.Get("").String())
}

func Test_adjustTSData(t *testing.T) {
	in := []byte(`{"subOffline":334,"a":"abc"}`)
	out := adjustTSData(in)
	t.Log(string(out))
	in = []byte(`{"ts":1646954803319,"values":{"humidity5":83.0,"temperature5":43.6}}`)
	out = adjustTSData(in)
	t.Log(string(out))

	in = []byte(`{"ModBus-TCP":{"ts":1649215733364,"values":{"wet":42,"temperature":"abc"}},"OPC-UA":{"ts":1649215733364,"values":{"counter":15}}}`)
	out = adjustTSData(in)
	t.Log(string(out))
}

var expr1 = `{"ID":"/core/v1/expressions/usr-57bea3a2d74e21ebbedde8268610/iotd-b10bcaa1-ba98-4e03-bece-6f852feb6edf/properties.telemetry.yc1",
	"Path":"properties.telemetry.yc1",
	"Name":"","Type":"eval","Owner":"usr-57bea3a2d74e21ebbedde8268610",
	"EntityID":"iotd-b10bcaa1-ba98-4e03-bece-6f852feb6edf",
	"Expression":"iotd-06a96c8d-c166-447c-afd1-63010636b362.properties.telemetry.src1",
	"Description":"iotd-06a96c8d-c166-447c-afd1-63010636b362=映射3,yc1=遥测1"}`
var expr2 = `{"ID":"/core/v1/expressions/usr-57bea3a2d74e21ebbedde8268610/iotd-b10bcaa1-ba98-4e03-bece-6f852feb6edf/properties.telemetry.yc2",
	"Path":"properties.telemetry.yc2",
	"Name":"","Type":"eval","Owner":"usr-57bea3a2d74e21ebbedde8268610",
	"EntityID":"iotd-b10bcaa1-ba98-4e03-bece-6f852feb6edf",
	"Expression":"iotd-06a96c8d-c166-447c-afd1-63010636b362.properties.telemetry.src2",
	"Description":"iotd-06a96c8d-c166-447c-afd1-63010636b362=映射3,yc1=遥测1"}`
var state = `{
	"properties": {
        "_version": {"type": "number"},
        "ts": {"type": "number"},
        "telemetry": {"src1": 123,"src2": 123}
    }
}`

type dispatcher_mock struct {
}

func (d *dispatcher_mock) DispatchToLog(ctx context.Context, bytes []byte) error {
	return nil
}

func (d *dispatcher_mock) Dispatch(ctx context.Context, event v1.Event) error {
	return nil
}

//iotd-b10bcaa1-ba98-4e03-bece-6f852feb6edf //目标
//- yc2 = iotd-06a96c8d-c166-447c-afd1-63010636b362.properties.telemetry.yc1  //e
//- yc1 = iotd-06a96c8d-c166-447c-afd1-63010636b362.properties.telemetry.yc1
func TestRuntime_handleComputed(t *testing.T) {
	placement.Initialize()
	placement.Global().Append(placement.Info{
		ID:   "core/1234",
		Flag: true,
	})
	en, err := NewEntity("iotd-06a96c8d-c166-447c-afd1-63010636b362", []byte(state))
	assert.Nil(t, err)

	rt := Runtime{
		dispatcher: &dispatcher_mock{},
		enCache: NewCacheMock(map[string]Entity{
			"iotd-06a96c8d-c166-447c-afd1-63010636b362": en,
		}),
		expressions: map[string]ExpressionInfo{},
		subTree:     path.NewRefTree(),
		evalTree:    path.New(),
	}

	var exprs = make([]repository.Expression, 0)
	exprs = append(exprs, updateExpr(t, rt, expr1))
	exprs = append(exprs, updateExpr(t, rt, expr2))
	feed := Feed{
		TTL:      0,
		Err:      nil,
		Event:    &v1.ProtoEvent{},
		State:    nil,
		EntityID: "iotd-06a96c8d-c166-447c-afd1-63010636b362",
		Patches:  nil,
		Changes: []Patch{
			{
				Op:    json.OpReplace,
				Path:  "properties.telemetry.src1",
				Value: tdtl.New(123),
			},
			{
				Op:    json.OpReplace,
				Path:  "properties.telemetry.src2",
				Value: tdtl.New(123),
			},
		},
	}
	got := rt.handleTentacle(context.Background(), &feed)
	t.Log(got)
	got = rt.handleComputed(context.Background(), &feed)
	t.Log(got)

	removeExpr(t, rt, expr1)
	removeExpr(t, rt, expr2)
	got = rt.handleTentacle(context.Background(), &feed)
	t.Log(got)
	got = rt.handleComputed(context.Background(), &feed)
	t.Log(got)
}

func updateExpr(t *testing.T, rt Runtime, exprRaw string) repository.Expression {
	var expr = repository.Expression{}
	expr.Decode([]byte(""), []byte(exprRaw))
	exprInfo1, err := parseExpression(expr, 1)
	assert.Nil(t, err)
	for runtimeID, exprIns := range exprInfo1 {
		t.Log(runtimeID)
		rt.AppendExpression(*exprIns)
	}
	return expr
}

func removeExpr(t *testing.T, rt Runtime, exprRaw string) repository.Expression {
	var expr = repository.Expression{}
	expr.Decode([]byte(""), []byte(exprRaw))
	exprInfo1, err := parseExpression(expr, 1)
	assert.Nil(t, err)
	for runtimeID, exprIns := range exprInfo1 {
		t.Log(runtimeID)
		rt.AppendExpression(*exprIns)
	}
	exprInfo := newExprInfo(&expr)
	rt.RemoveExpression(exprInfo.ID)
	return expr
}

func Test_mergePath(t *testing.T) {
	tests := []struct {
		name       string
		subPath    string
		changePath string
		want       string
	}{
		{"1", "dev1.a.b.*", "dev1.a.b.c", "a.b"},
		{"1", "dev1.a.b.c.d", "dev1.a.b.c", "a.b.c.d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergePath(tt.subPath, tt.changePath); got != tt.want {
				t.Errorf("mergePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
