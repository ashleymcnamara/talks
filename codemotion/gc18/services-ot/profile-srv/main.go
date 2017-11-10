package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bketelsen/talks/codemotion/gc18/services-ot/profile-srv/db"
	"github.com/bketelsen/talks/codemotion/gc18/services-ot/profile-srv/handler"
	"github.com/bketelsen/talks/codemotion/gc18/services-ot/profile-srv/proto/record"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	mot "github.com/micro/go-plugins/wrapper/trace/opentracing"

	opentracing "github.com/opentracing/opentracing-go"

	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const ServiceName = "gophercon.srv.profileot"

var (
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// Host + port of our service.
	hostPort = "0.0.0.0:0"
	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

func main() {
	tracer := initTracer()
	opentracing.SetGlobalTracer(tracer)

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/profile",
			},
		),

		micro.WrapClient(mot.NewClientWrapper(tracer)),
		micro.WrapHandler(mot.NewHandlerWrapper(tracer)),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()

	db.Init()

	record.RegisterRecordHandler(service.Server(), new(handler.Record))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func initTracer() opentracing.Tracer {

	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, debug, hostPort, ServiceName)

	// Create our tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	// Explicitly set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)
	return tracer
}
