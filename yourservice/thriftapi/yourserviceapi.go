package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/vaporz/turbo"
	"net/http"
	"reflect"
	"github.com/vaporz/turbo-example/yourservice/gen"
	t "github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	i "github.com/vaporz/turbo-example/yourservice/interceptor"
)

func main() {
	turbo.Intercept([]string{"GET"}, "/hello", i.LogInterceptor{})
	turbo.RegisterMessageFieldConvertor(new(t.HelloValues), convertHelloValues)
	turbo.StartThriftHTTPServer("github.com/vaporz/turbo-example/yourservice", thriftClient, gen.ThriftSwitcher)
}

func thriftClient(trans thrift.TTransport, f thrift.TProtocolFactory) interface{} {
	return t.NewYourServiceClientFactory(trans, f)
}

func convertHelloValues(req *http.Request) reflect.Value {
	result := &t.HelloValues{}
	result.Message = "a message from convertor"
	return reflect.ValueOf(result)
}
