package main

import (
	"github.com/vaporz/turbo"
	"github.com/vaporz/turbo-example/yourservice/gen"
	tg "github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"google.golang.org/grpc"
	"github.com/vaporz/turbo-example/yourservice/gen/proto"
	"github.com/vaporz/turbo-example/yourservice/grpcservice/impl"
	timpl "github.com/vaporz/turbo-example/yourservice/thriftservice/impl"
	"git.apache.org/thrift.git/lib/go/thrift"
	t "github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
)

func main() {
	//turbo.StartGRPC("github.com/vaporz/turbo-example/yourservice", "service",
	//	50051, grpcClient, gen.GrpcSwitcher, registerServer)

	turbo.StartTHRIFT("github.com/vaporz/turbo-example/yourservice", "service",
		50052, thriftClient, gen.ThriftSwitcher, _TProcessor)
}

func grpcClient(conn *grpc.ClientConn) interface{} {
	return proto.NewYourServiceClient(conn)
}

func registerServer(s *grpc.Server) {
	proto.RegisterYourServiceServer(s, &impl.YourService{})
}

func thriftClient(trans thrift.TTransport, f thrift.TProtocolFactory) interface{} {
	return t.NewYourServiceClientFactory(trans, f)
}

func _TProcessor() thrift.TProcessor {
	return tg.NewYourServiceProcessor(timpl.YourService{})
}
