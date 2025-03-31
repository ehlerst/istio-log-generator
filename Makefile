otel:
	docker run -v $(pwd)/config.yaml:/etc/otelcol-contrib/config.yaml otel/opentelemetry-collector-contrib:0.121.0

run:
	go run main.go

start-loki:
	-docker volume create --name=grafana-data
	docker-compose up -d

clean-loki:
	docker-compose down
	sudo rm -rf data/*
	-docker volume rm grafana-data


