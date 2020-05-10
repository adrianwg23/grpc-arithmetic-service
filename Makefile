build-client:
	docker-compose build grpc-client

build-server:
	docker-compose build grpc-server

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

kube-delete-service-client:
	kubectl delete service grpc-client-service

kube-delete-deployment-client:
	kubectl delete deploy grpc-client-deployment

kube-delete-service-server:
	kubectl delete service grpc-server-service

kube-delete-deployment-server:
	kubectl delete deploy grpc-server-deployment

kube-run-client:
	kubectl create -f deploy/client-deployment.yml

kube-run-server:
	kubectl create -f deploy/server-deployment.yml

kube-stop-client: kube-delete-deployment-client kube-delete-service-client

kube-stop-server: kube-delete-deployment-server kube-delete-service-server

.PHONY: build-client build-server docker-up docker-down kube-delete-service-client kube-delete-deployment-client kube-delete-service-server kube-delete-deployment-server kube-run-client kube-run-server kube-stop-client kube-stop-server
