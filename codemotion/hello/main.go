package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/bketelsen/talks/codemotion/hello/handler"
	"github.com/bketelsen/talks/codemotion/hello/subscriber"

	example "github.com/bketelsen/talks/codemotion/hello/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.hello", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.hello", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
