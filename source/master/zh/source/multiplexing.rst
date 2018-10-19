.. _multiplexing:

启动多个service
================

Server端有时需要定义多个Service，并在服务启动时在一个进程同时启动，共用端口。

gRPC
----
1. 在.proto文件中定义新的service

.. code-block:: diff

 +message MeowRequest {
 +}

 +message MeowResponse {
 +  string message = 1;
 +}

 +service CatService {
 +  rpc meow (MeowRequest) returns (MeowResponse) {}
 +}

2. 修改service.yaml，添加新API

.. code-block:: diff

 config:
 -  grpc_service_name: TestService
 +  grpc_service_name: TestService,CatService
 urlmapping:
 + - GET /cat/meow CatService Meow

3. 重新生成代码

.. code-block:: shell

 $ turbo generate package/path/to/yourservice -r grpc -I $GOPATH/src/package/path/to/yourservice

4. 在Server端，编写实现新service

.. code-block:: diff

 func RegisterServer(s *grpc.Server) {
 	proto.RegisterTestServiceServer(s, &TestService{})
 +	proto.RegisterCatServiceServer(s, &CatService{})
 }

 +type CatService struct {
 +}

 +func (c *CatService) Meow(ctx context.Context, req *proto.MeowRequest) (*proto.MeowResponse, error) {
 +	return &proto.MeowResponse{Message: "Meow!"}, nil
 +}

5. 在HTTP端，注册新service

.. code-block:: diff

 func GrpcClient(conn *grpc.ClientConn) map[string]interface{} {
 	return map[string]interface{}{
 		"TestService": proto.NewTestServiceClient(conn),
 +		"CatService": proto.NewCatServiceClient(conn),
 	}
 }

完成！

Thrift
------
1. 在.thrift文件中定义新的service

.. code-block:: diff

 +struct MeowResponse {
 + 1: string message,
 +}

 +service CatService {
 +  MeowResponse Meow (1:string msg)
 +}

2. 修改service.yaml，添加新API

.. code-block:: diff

 config:
 -  thrift_service_name: TestService
 +  thrift_service_name: TestService,CatService
 urlmapping:
 + - GET /cat/meow CatService Meow

3. 重新生成代码

.. code-block:: shell

 $ turbo generate package/path/to/yourservice -r thrift -I $GOPATH/src/package/path/to/yourservice

4. 在Server端，编写实现新service

.. code-block:: diff

 func TProcessor() map[string]thrift.TProcessor {
	return map[string]thrift.TProcessor{
		"TestService": gen.NewTestServiceProcessor(TestService{}),
 +		"CatService": gen.NewCatServiceProcessor(CatService{}),
	}
 }

 +type CatService struct {
 +}

 +func (c CatService) Meow(msg string) (r *gen.MeowResponse, err error) {
 +	return &gen.MeowResponse{Message: msg}, nil
 +}

5. 在HTTP端，注册新service

.. code-block:: diff

 func ThriftClient(trans thrift.TTransport, f thrift.TProtocolFactory) map[string]interface{} {
	iprot := f.GetProtocol(trans)
	return map[string]interface{}{
		"TestService": t.NewTestServiceClientProtocol(trans, iprot, thrift.NewTMultiplexedProtocol(iprot, "TestService")),
 +		"CatService": t.NewCatServiceClientProtocol(trans, iprot, thrift.NewTMultiplexedProtocol(iprot, "CatService")),
	}
 }

完成！
