package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"github.com/vaporz/turbo-example/yourservice/thriftservice/impl"
	"github.com/vaporz/turbo"
)

func main() {
	turbo.StartThriftService(50052, _TProcessor)
}

func _TProcessor() thrift.TProcessor {
	return gen.NewYourServiceProcessor(impl.YourService{})
}
