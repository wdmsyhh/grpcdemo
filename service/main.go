package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	hello "grpcdemo/proto"
	hello2 "grpcdemo/service/hello"
	"log"
	"net"
	"net/http"
)

func main() {
	go startGRPCGateway()

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(hello2.HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(grpcServer.Serve(lis))
	log.Println("Server end")
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}))
	err := hello.RegisterHelloServiceHandlerFromEndpoint(c, mux, ":1234", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cann't start grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8080", mux) // grpc gateway 的
	if err != nil {
		log.Fatalf("cann't listen and serve: %v", err)
	}
	log.Println("Gateway end")
}
