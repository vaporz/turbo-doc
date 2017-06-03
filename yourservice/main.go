package main

import (
	"github.com/vaporz/turbo"
	"github.com/vaporz/turbo-example/yourservice/gen"
	gcomponent "github.com/vaporz/turbo-example/yourservice/grpcapi/component"
	gimpl "github.com/vaporz/turbo-example/yourservice/grpcservice/impl"
	//tcompoent "github.com/vaporz/turbo-example/yourservice/thriftapi/component"
	//timpl "github.com/vaporz/turbo-example/yourservice/thriftservice/impl"
)

func main() {
	turbo.StartGRPC("github.com/vaporz/turbo-example/yourservice", "service",
		50051, gcomponent.GrpcClient, gen.GrpcSwitcher, gimpl.RegisterServer)

	//turbo.StartTHRIFT("github.com/vaporz/turbo-example/yourservice", "service",
	//	50052, tcompoent.ThriftClient, gen.ThriftSwitcher, timpl.TProcessor)
}
