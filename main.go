package main

import (
	"fmt"
	"context"

	"encoding/json"

	"github.com/ehlerst/istio-log-generator/pkg"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"

)

const (
	GB = 1024*1024*1024 * 1 
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
		json, err := json.MarshalIndent(log, "", "  ")
		b := len(json)
		totalSent += b
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}
		
		logger.InfoContext(context.Background(), string(json))

		if totalSent > GB { 
			fmt.Printf("Total bytes sent %d\n", totalSent)
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