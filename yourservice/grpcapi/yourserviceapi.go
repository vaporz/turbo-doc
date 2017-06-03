package main

import (
	"github.com/vaporz/turbo"
	"github.com/vaporz/turbo-example/yourservice/gen"
	"github.com/vaporz/turbo-example/yourservice/grpcapi/component"
)

func main() {
	component.InitComponents()
	turbo.StartGrpcHTTPServer("github.com/vaporz/turbo-example/yourservice",
		"service", component.GrpcClient, gen.GrpcSwitcher)
}
