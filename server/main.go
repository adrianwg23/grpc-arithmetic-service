package main

import (
	"context"
	"fmt"
	"github.com/adrianwg23/grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterArithmeticServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("Request served!")
	ip := getIP()
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result, Ip: ip}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("Request served!")
	ip := getIP()
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result, Ip: ip}, nil
}

func getIP() []byte {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			return ipv4
		}
	}

	return []byte("no ip")
}
