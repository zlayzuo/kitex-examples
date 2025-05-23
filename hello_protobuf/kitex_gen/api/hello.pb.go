// Code generated by Kitex v0.13.1. DO NOT EDIT.

package api

import (
	"context"

	"github.com/cloudwego/prutal"
)

type Request struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (x *Request) Reset() { *x = Request{} }

func (x *Request) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *Request) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (x *Response) Reset() { *x = Response{} }

func (x *Response) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *Response) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Hello interface {
	Echo(ctx context.Context, req *Request) (res *Response, err error)
}
