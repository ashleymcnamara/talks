package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/envy"

	"github.com/bketelsen/talks/codemotion/gc18/gophercon-ot/actions"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

// ServiceName is the string name of the service, since
// it is used in multiple places, it's an exported Constant
const ServiceName = "gophercon.web"

var (
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// Host + port of our service.
	hostPort = "0.0.0.0:3000"
	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

func main() {

	tracer := initTracer()
	opentracing.SetGlobalTracer(tracer)

	port := envy.Get("PORT", "3001")
	actions.Tracer = tracer
	app := actions.App()

	log.Fatal(app.Start(port))
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
