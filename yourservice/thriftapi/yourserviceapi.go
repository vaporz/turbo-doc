package main

import (
	"github.com/vaporz/turbo"
	"github.com/vaporz/turbo-example/yourservice/gen"
	"github.com/vaporz/turbo-example/yourservice/thriftapi/component"
)

func main() {
	component.InitComponents()
	turbo.StartThriftHTTPServer("github.com/vaporz/turbo-example/yourservice",
		"service", component.ThriftClient, gen.ThriftSwitcher)
}
