package main

import (
	"github.com/vaporz/turbo-example/yourservice/grpcservice/impl"
	"github.com/vaporz/turbo"
)

func main() {
	turbo.StartGrpcService(50051, impl.RegisterServer)
}
