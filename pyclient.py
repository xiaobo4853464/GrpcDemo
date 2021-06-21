import grpc

from example.order_pb2 import OrderRequest

from example.order_pb2_grpc import OrderServiceStub


# # 通过protoc 生成客户端
# pip install grpcio
# pip install grpcio-tools
# python3 -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. example/order.proto

def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel('localhost:8071')
    # 调用 rpc 服务
    stub = OrderServiceStub(channel)
    response = stub.GetOrder(OrderRequest(order_id="no23"))

    print(response)


if __name__ == '__main__':
    run()
