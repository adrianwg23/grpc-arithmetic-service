# gRPC Arithmetic Service
gRPC Arithmetic Service is a simple service I've created to better understand Golang, gRPC, and Kubernetes. This service
is a simple service that performs simple Addition and Multiplication on 2 integers. This repo contains both the `grpc-client`
and `grpc-server`. The `grpc-client` is a simple HTTP server that consumes two GET endpoints and forwards traffic via grpc 
to the `grpc-server`. The grpc-server is simply a grpc server configured with two rpc's, `Add` and `Multiply`.

## Dependencies
1. Gin (client HTTP server library)
2. gRPC 
3. Linkerd (Kubernetes service mesh/sidecar proxy solution)
4. nginx-ingress-controller (NGINX implementation of k8s Ingress Controller)
## How to Run  

### Via minikube
This service has been configured to run flawlessly on a single node kubernetes cluster via minikube. To run this service 
on minikube, ensure that you have minikube and docker installed on your computer. With minikube installed, run:
```
minikube start
```

With minikube started, you will first need to configure minikube to use the Docker daemon:
```
eval $(minikube docker-env)
```

You will then need to enable the ingress addon in minikube. Enabling this will automatically create an `nginx-ingress-controller`
in the `kube-system` namespace. You can enable the ingress addon by running:
```
minikube addons enable ingress
``` 

Finally, you will need to install `Linkerd`. Linkerd is a service mesh solution that allows us to load balance on the L7 layer.
This is essential since Kubernetes' `Service` object load balances pods on the L4 layer while gRPC leverages HTTP2 technology. 
With HTTP2, requests are multiplexed by the same TCP connection, so L4 layer loadbalancing will not suffice. `Linkerd` solves this
by injecting a sidecar proxy container `linkerd2-proxy`. `linkerd2-proxy` lives in your pod and listens to outbound requests 
to other services. These requests are proxied and load balanced by the Linkerd proxy. This allows us to load balance amongst
are gRPC pods in the cluster! Visit [Linkerd's getting started](https://linkerd.io/2/getting-started/) for more details. To install 
`Linkerd` on mac, run:

```
brew install linkerd
```

With the Linkerd cli installed, we need to install the necessary Linkerd containers in our k8s cluster for the service mesh to work.
We can do this by running:
```
linkerd install | kubectl apply -f -
```

Our cluster is now setup! We now need to build are grpc-client and grpc-server images and get them running.
We build our client and server images by running:
```
make build-client
make build-server
```

With the images created in our k8s cluster, we can now run the necessary resources in our k8s cluster!
```
make kube-create-namespace
make kube-run-client
make kube-run-server
```
Our services are now running! But how do we access it from the outside world? To do so, we will need to slightly modify the Ingress rules.

#### Configuring Ingress
In order to access our cluster from the outside, there is some configuration that will need to be made to `grpc-arithmetic-ingress.yml`. 
The first thing you will need to do is decide whether you want TLS enabled or not. If not, go ahead and remove the `tls` tag altogether.
If you want TLS enabled, you will need to generate a signed certificate and key and create a `grpc-arithmetic-tls` Kubernetes secret. 
You can see view the template of the TLS Secret object [here](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls).

Next, you will need to configure the host name you want exposed outside of the cluster. I've chosen `arithmeticgrpc.com`. You are free
to change this to any hostname you want/own. If you want to choose a hostname that you do not own, simply fetch the IP of your minikube cluster
by running `minikube ip` and modifying your `hosts` file. You can modify your hosts file by running

```
sudo vim /etc/hosts
```

and adding your `Hostname MINIKUBE_IP` in the config.

With your Ingress rules now setup, we can now run: 

```
make kube-create-ingress
```
to create our Ingress object!

## Testing
To test, we can do a simple curl to our Ingress Controller. Using HTTP, we can curl:
```
curl -X GET "http://<your_hostname>.com/add/2/3"
curl -X GET "http://<your_hostname>.com/mult/9/7"
```

With HTTPS, we can curl:
```
curl -k -X GET "http://<your_hostname>.com/add/2/3"
curl -k -X GET "http://<your_hostname>.com/mult/9/7"
```

Alternatively, you can run the python script `makerequests.py` in the `/scripts` directory. This script will make an HTTP
curl to the two endpoints every `0.5s`. 

## Running via docker-compose
To run the service via docker-compose, simply run:
```
docker-compose up -d
```
The grpc-client service is configured to run on port `8080`, so you can run a simple curl to test:
```
curl -X GET "http://localhost:8080/add/2/3"
```

## Running via CLI
To run the service locally on your command line, ensure you have Golang 1.14 or greater installed.
You can install Golang by running `brew install golang`.

With Go installed, from the root directory, simply run:
```
go run server/main.go
go run client/main.go
```

The grpc-client service is configured to run on port `8080`
```
curl -X GET "http://localhost:8080/add/2/3"
```