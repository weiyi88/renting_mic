package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	helloMicro "helloMicro/proto"
)

type HelloMicro struct{}

// Return a new handler
func New() *HelloMicro {
	return &HelloMicro{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *HelloMicro) Call(ctx context.Context, req *helloMicro.Request, rsp *helloMicro.Response) error {
	log.Info("Received HelloMicro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *HelloMicro) Stream(ctx context.Context, req *helloMicro.StreamingRequest, stream helloMicro.HelloMicro_StreamStream) error {
	log.Infof("Received HelloMicro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&helloMicro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *HelloMicro) PingPong(ctx context.Context, stream helloMicro.HelloMicro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&helloMicro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
