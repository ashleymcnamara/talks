package handler

import (
	"github.com/bketelsen/talks/codemotion/greeting/proto/greeting"
	"github.com/micro/go-log"
	"golang.org/x/net/context"
)

type Greeting struct{}

func (e *Greeting) Hello(ctx context.Context, req *greeting.HelloRequest, rsp *greeting.HelloResponse) error {
	log.Log("Received Greeting.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Greeting) Goodbye(ctx context.Context, req *greeting.GoodbyeRequest, rsp *greeting.GoodbyeResponse) error {
	log.Log("Received Greeting.Goodbye request")
	rsp.Msg = "Goodbye " + req.Name
	return nil
}
