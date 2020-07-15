package main

import (
	"google.golang.org/grpc"
	"net"
	"studyGrpc/services"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis,_ := net.Listen("tcp", ":8081")
	_ = rpcServer.Serve(lis)
}