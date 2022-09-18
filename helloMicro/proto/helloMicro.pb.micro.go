// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/helloMicro.proto

package helloMicro

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HelloMicro service

func NewHelloMicroEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HelloMicro service

type HelloMicroService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (HelloMicro_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (HelloMicro_PingPongService, error)
}

type helloMicroService struct {
	c    client.Client
	name string
}

func NewHelloMicroService(name string, c client.Client) HelloMicroService {
	return &helloMicroService{
		c:    c,
		name: name,
	}
}

func (c *helloMicroService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "HelloMicro.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloMicroService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (HelloMicro_StreamService, error) {
	req := c.c.NewRequest(c.name, "HelloMicro.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &helloMicroServiceStream{stream}, nil
}

type HelloMicro_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type helloMicroServiceStream struct {
	stream client.Stream
}

func (x *helloMicroServiceStream) Close() error {
	return x.stream.Close()
}

func (x *helloMicroServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloMicroServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloMicroServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloMicroServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloMicroService) PingPong(ctx context.Context, opts ...client.CallOption) (HelloMicro_PingPongService, error) {
	req := c.c.NewRequest(c.name, "HelloMicro.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &helloMicroServicePingPong{stream}, nil
}

type HelloMicro_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type helloMicroServicePingPong struct {
	stream client.Stream
}

func (x *helloMicroServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *helloMicroServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *helloMicroServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloMicroServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloMicroServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *helloMicroServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HelloMicro service

type HelloMicroHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, HelloMicro_StreamStream) error
	PingPong(context.Context, HelloMicro_PingPongStream) error
}

func RegisterHelloMicroHandler(s server.Server, hdlr HelloMicroHandler, opts ...server.HandlerOption) error {
	type helloMicro interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type HelloMicro struct {
		helloMicro
	}
	h := &helloMicroHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloMicro{h}, opts...))
}

type helloMicroHandler struct {
	HelloMicroHandler
}

func (h *helloMicroHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.HelloMicroHandler.Call(ctx, in, out)
}

func (h *helloMicroHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HelloMicroHandler.Stream(ctx, m, &helloMicroStreamStream{stream})
}

type HelloMicro_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type helloMicroStreamStream struct {
	stream server.Stream
}

func (x *helloMicroStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloMicroStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloMicroStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloMicroStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloMicroStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *helloMicroHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.HelloMicroHandler.PingPong(ctx, &helloMicroPingPongStream{stream})
}

type HelloMicro_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type helloMicroPingPongStream struct {
	stream server.Stream
}

func (x *helloMicroPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *helloMicroPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloMicroPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloMicroPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloMicroPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *helloMicroPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}