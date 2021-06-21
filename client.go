package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcDemo2/example"
)

func main() {
	conn, err := grpc.Dial("localhost:8071", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := example.NewOrderServiceClient(conn)
	req := &example.OrderRequest{OrderId: "no1"}
	resp, err := client.GetOrder(context.Background(), req)
	fmt.Println(resp)

}
