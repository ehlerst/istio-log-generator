otel:
	docker run -v $(pwd)/config.yaml:/etc/otelcol-contrib/config.yaml otel/opentelemetry-collector-contrib:0.121.0

run:
	go run main.go