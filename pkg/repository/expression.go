package repository

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"github.com/tkeel-io/core/pkg/repository/dao"
	"github.com/tkeel-io/core/pkg/util"
	"github.com/tkeel-io/kit/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

const (
	ExprTypeSub  = "sub"
	ExprTypeEval = "eval"
)

type ListExprReq struct {
	Owner    string
	EntityID string
}

type defaultExprCodec struct{}

func (ec *defaultExprCodec) Key() dao.Codec {
	return &defaultExprKeyCodec{}
}

func (ec *defaultExprCodec) Value() dao.Codec {
	return &defaultExprValueCodec{}
}

type defaultExprKeyCodec struct{}
type defaultExprValueCodec struct{}

func (dec *defaultExprKeyCodec) Encode(v interface{}) ([]byte, error) {
	panic("implement me")
}

func (dec *defaultExprKeyCodec) Decode(raw []byte, v interface{}) error {
	panic("implement me")
}

func (dec *defaultExprValueCodec) Encode(v interface{}) ([]byte, error) {
	panic("implement me")
}

func (dec *defaultExprValueCodec) Decode(raw []byte, v interface{}) error {
	panic("implement me")
}

type Expression struct {
	// expression identifier.
	ID string
	// target path.
	Path string
	// expression name.
	Name string
	// expression type.
	Type string
	// expression owner.
	Owner string
	// entity id.
	EntityID string
	// expression.
	Expression string
	// description.
	Description string
}

func NewExpression(owner, entityID, name, path, expr, desc string) *Expression {
	escapePath := url.PathEscape(path)
	typ := ExprTypeEval
	if escapePath == "" {
		path = util.UUID("exprsub")
		escapePath = url.PathEscape(path)
		typ = ExprTypeSub
	}

	identifier := fmt.Sprintf("", owner, entityID, escapePath)
	return &Expression{
		ID:          identifier,
		Name:        name,
		Path:        path,
		Type:        typ,
		Owner:       owner,
		EntityID:    entityID,
		Expression:  expr,
		Description: desc,
	}
}

func (e *Expression) Codec() dao.KVCodec {
	return &defaultExprCodec{}
}

func (r *repo) PutExpression(ctx context.Context, expr Expression) error {
	err := r.dao.PutResource(ctx, &expr)
	return errors.Wrap(err, "put expression repository")
}

func (r *repo) GetExpression(ctx context.Context, expr Expression) (Expression, error) {
	_, err := r.dao.GetResource(ctx, &expr)
	return expr, errors.Wrap(err, "get expression repository")
}

func (r *repo) DelExpression(ctx context.Context, expr Expression) error {
	err := r.dao.DelResource(ctx, &expr)
	return errors.Wrap(err, "del expression repository")
}

func (r *repo) DelExprByEnity(ctx context.Context, expr Expression) error {
	// construct prefix key.
	var prefix string

	err := r.dao.DelResources(ctx, prefix)
	return errors.Wrap(err, "del expressions repository")
}

func (r *repo) HasExpression(ctx context.Context, expr Expression) (bool, error) {
	has, err := r.dao.HasResource(ctx, &expr)
	return has, errors.Wrap(err, "exists expression repository")
}

func (r *repo) ListExpression(ctx context.Context, rev int64, req *ListExprReq) ([]*Expression, error) {
	// construct prefix.
	prefix := ""
	ress, err := r.dao.ListResource(ctx, rev, prefix,
		func(raw []byte) (dao.Resource, error) {
			var res Expression // escape.
			valCodec := &defaultExprValueCodec{}
			err := valCodec.Decode(raw, &res)
			return &res, errors.Wrap(err, "decode expression")
		})

	var exprs []*Expression
	for index := range ress {
		if expr, ok := ress[index].(*Expression); ok {
			exprs = append(exprs, expr)
			continue
		}
		// panic.
	}
	return exprs, errors.Wrap(err, "list expression repository")
}

func (r *repo) RangeExpression(ctx context.Context, rev int64, handler RangeExpressionFunc) {
	var prefix string
	r.dao.RangeResource(ctx, rev, prefix, func(kvs []*mvccpb.KeyValue) {
		var exprs []*Expression
		valCodec := &defaultExprValueCodec{}
		for index := range kvs {
			var expr Expression
			err := valCodec.Decode(kvs[index].Value, &expr)
			if nil != err {
				log.L().Error("")
				continue
			}
			exprs = append(exprs, &expr)
		}
		handler(exprs)
	})
}

func (r *repo) WatchExpression(ctx context.Context, rev int64, handler WatchExpressionFunc) {
	var prefix string
	r.dao.RangeResource(ctx, rev, prefix, func(kvs []*mvccpb.KeyValue) {
		valCodec := &defaultExprValueCodec{}
		for index := range kvs {
			var expr Expression
			err := valCodec.Decode(kvs[index].Value, &expr)
			if nil != err {
				log.L().Error("")
				continue
			}
			handler(expr)
		}
	})
}

type RangeExpressionFunc func([]*Expression)
type WatchExpressionFunc func(Expression)