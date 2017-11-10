package main

import (
	"github.com/bketelsen/talks/codemotion/greeting/handler"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	greeting "github.com/bketelsen/talks/codemotion/greeting/proto/greeting"
)
//import _ "github.com/micro/go-plugins/registry/kubernetes"

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.greeting"),
		micro.Version("0.1"),
	)

	// Register Handler
	greeting.RegisterGreetingHandler(service.Server(), new(handler.Greeting))
	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
