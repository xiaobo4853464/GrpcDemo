package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcDemo2/example"
	"net"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrder(ctx context.Context, req *example.OrderRequest) (*example.OrderResponse, error) {
	orders := map[string]example.OrderResponse{
		"no1": {OrderId: "orderId1", OrderName: "cloth", OrderStatus: "normal"},
		"no2": {OrderId: "orderId2", OrderName: "shoes", OrderStatus: "normal"},
		"no3": {OrderId: "orderId3", OrderName: "hats", OrderStatus: "unnormal"},
	}
	if v, ok := orders[req.OrderId]; ok {
		//return &example.OrderResponse{OrderId: v.OrderId, OrderName: v.OrderName, OrderStatus: v.OrderStatus}, nil
		return &v, nil
	}
	return nil, nil
}

func main() {
	server := grpc.NewServer()
	example.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	listen, err := net.Listen("tcp", ":8071")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
