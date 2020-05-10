build-client:
	docker-compose build grpc-client

build-server:
	docker-compose build grpc-server

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

kube-run-client:
	kubectl apply -f deploy/client-deployment.yml

kube-run-server:
	kubectl apply -f deploy/server-deployment.yml

kube-create-namespace:
	kubectl apply -f deploy/grpc-arithmetic-service-namespace.yml

kube-create-ingress:
	kubectl apply -f deploy/grpc-arithmetic-ingress.yml

kube-delete-namespace:
	kubectl delete namespace grpc-arithmetic-service

kube-rollout-client:
	kubectl rollout restart deployment grpc-client-deployment

kube-rollout-server:
	kubectl rollout restart deployment grpc-server-deployment

kube-get-namespace:
	kubectl config view | grep namespace

# this target lists all the make targets in this Makefile
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null \
		| awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' \
		| sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

.PHONY: list build-client build-server docker-up docker-down kube-run-client kube-run-server kube-create-namespace kube-create-ingress kube-delete-namespace kube-rollout-client kube-rollout-server kube-get-namespace
