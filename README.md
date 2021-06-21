#GRPC调用demo
## 该demo 掩饰了go 写的服务器，分别使用go、python client进行调用
### 步骤：
1. 编写example目录下的order.proto文件
2. 终端执行命令生成go 代码  
"protoc --go_out=plugins=grpc:. example/order.proto" 
3. 编写go server
4. 编写go client
5. 终端执行命令生成python 代码  
"python3 -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. example/order.proto"
6. 编写python client 请求 
