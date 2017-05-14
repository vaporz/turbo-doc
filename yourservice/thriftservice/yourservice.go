package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"os"
	"turbo-example/yourservice/gen/thrift/gen-go/gen"
	"turbo-example/yourservice/thriftservice/impl"
)

func main() {
	transport, err := thrift.NewTServerSocket(":50052")
	if err != nil {
		log.Println("socket error")
		os.Exit(1)
	}

	server := thrift.NewTSimpleServer4(gen.NewYourServiceProcessor(impl.YourService{}), transport,
		thrift.NewTTransportFactory(), thrift.NewTBinaryProtocolFactoryDefault())
	server.Serve()
}
