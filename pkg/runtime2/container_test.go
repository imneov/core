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

package runtime3

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainer_HandleEvent(t *testing.T) {
	ctx := context.Background()
	ev := ContainerEvent{
		ID:   "entity1",
		Type: OpEntity,
		Value: &EntityEvent{
			JSONPath: "a.b.c",
			OP:       "replace",
			Value:    []byte(`"abc"`),
		},
	}
	cc := NewContainer(ctx, "core-1")
	ret, err := cc.HandleEvent(ctx, &ev)
	assert.NoError(t, err, "err is %v", err)
	byt, err := cc.entities["entity1"].Raw()
	t.Log(cc, string(byt))
	t.Log(string(ret.State))
}

func TestContainer_OpEntity_HandleEvent(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name       string
		typ        EntityEventType
		path       string
		eventValue []byte
		wantErr    bool
		want       string
	}{
		{"1", OpEntityPropsUpdata, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{"a":{"b":{"c":"abc"}}},"Scheme":{}}`},
		//{"2", OpEntityPropsPatch, "a.b.c", []byte(`"abc"`), false, `{"Property":{"a":{"b":{"c":"abc"}}},"Scheme":{}}`},
		{"3", OpEntityPropsGet, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{}}`},
		{"4", OpEntityConfigsUpdata, "a.b.c", []byte(`"abc"`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{"a":{"b":{"c":"abc"}}}}`},
		//{"5", OpEntityConfigsPatch, "a.b.c", []byte(`"abc"`), false, `{"Property":{},"Scheme":{"a":{"b":{"c":"abc"}}}}`},
		{"6", OpEntityConfigsGet, "a.b.c", []byte(`{
			Id:     "device123",
			Type:   "DEVICE",
			Owner:  "tomas",
			Source: "CORE-SDK",
			Properties: map[string]interface{}{
			"temp": 20,
		},`), false, `{"ID":"","Type":"","Owner":"","Source":"","Version":0,"LastTime":0,"TemplateID":"","Property":{},"Scheme":{}}`},
	}
	ctx := context.Background()
	for _, tt := range tests {
		cc := newContainer()
		t.Run(tt.name, func(t *testing.T) {
			ev := &ContainerEvent{ID: "123", Type: OpEntity,
				Callback: "Http://127.0.0.1:8088/callback/8888",
				Value: &EntityEvent{
					JSONPath: tt.path,
					OP:       tt.typ,
					Value:    []byte(`"abc"`),
				}}
			got, err := cc.HandleEvent(ctx, ev)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got.State), tt.want) {
				t.Errorf("HandleEvent() \ngot = %v, \nwant  %v", string(got.State), tt.want)
			}
		})
	}
}

func newContainer() *Container {
	cc := NewContainer(context.Background(), "core-1")
	return cc
}