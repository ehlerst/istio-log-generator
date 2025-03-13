package main

import (
	"fmt"
	"context"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/ehlerst/istio-log-generator/pkg"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"

)

var 
(
	logger			 = otelslog.NewLogger("istio-log-generator")
	serviceName      = "istio-log-generator"
	otelCollectorURL = "localhost:4317"
)

func main() {
	loggerProvider, err := newLoggerProvider()
	if err != nil {
		fmt.Printf("Failed to create logger provider: %v", err)
		return
	}

	defer func() {
		if err := loggerProvider.Shutdown(context.Background()); err != nil {
			fmt.Println("Error shutting down logger provider")
		}
	}()

	global.SetLoggerProvider(loggerProvider)
	totalSent := 0
	for {
		log := generate.GenerateIstioLog()
		log.StartTime = time.Now()
		log.TraceID = fmt.Sprintf("%x", rand.Int63())
		log.ResponseCode = generate.GenerateRandomCode()
		log.UpstreamLocalAddress = generate.GenerateRandomIP()
		log.XDatadogTraceID = fmt.Sprintf("%x", rand.Int63())
		log.BytesSent = rand.Intn(8096) + 512
		log.DownstreamRemoteAddress = generate.GenerateRandomIP()
		log.DownstreamLocalAddress = generate.GenerateRandomIP()
		log.Duration = rand.Intn(500) + 100
		log.UpstreamServiceTime = fmt.Sprintf("%d", rand.Intn(500) + 100)
		log.UpstreamCluster = fmt.Sprintf("%x", rand.Int63())
		log.Authority = fmt.Sprintf("%x", rand.Int63())
		json, err := json.MarshalIndent(log, "", "  ")
		b := len(json)
		totalSent += b
		fmt.Printf("Total bytes sent: %d\n", totalSent)
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}
		
		logger.InfoContext(context.Background(), string(json))

		if totalSent > 1024*1024 * 1 { // GB
			break
		}
	}

	loggerProvider.ForceFlush(context.TODO())

}


func newLoggerProvider() (*log.LoggerProvider, error) {
	logExporter, err := otlploggrpc.New(
		context.Background(),
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithEndpoint(otelCollectorURL),
	)

	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}