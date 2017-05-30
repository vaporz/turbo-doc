package main

import (
	"google.golang.org/grpc"
	"github.com/vaporz/turbo-example/yourservice/gen/proto"
	"github.com/vaporz/turbo-example/yourservice/grpcservice/impl"
	"github.com/vaporz/turbo"
)

func main() {
	turbo.StartGrpcService(50051, registerServer)
}

func registerServer(s *grpc.Server) {
	proto.RegisterYourServiceServer(s, &impl.YourService{})
}
