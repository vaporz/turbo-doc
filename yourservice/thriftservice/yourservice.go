package main

import (
	"github.com/vaporz/turbo-example/yourservice/thriftservice/impl"
	"github.com/vaporz/turbo"
)

func main() {
	turbo.StartThriftService(50052, impl.TProcessor)
}
