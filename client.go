package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"studyGrpc/services"
)

func main() {
	conn,err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn);
	resp,_ := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 1})
	fmt.Println(resp.GetProdStock())
}