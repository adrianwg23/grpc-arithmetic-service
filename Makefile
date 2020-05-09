build-client:
	docker-compose build grpc-client

build-server:
	docker-compose build grpc-server

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

.PHONY: build-client build-server docker-up docker-down
