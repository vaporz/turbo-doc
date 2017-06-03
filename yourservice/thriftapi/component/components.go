package component

import (
	"github.com/vaporz/turbo"
	"git.apache.org/thrift.git/lib/go/thrift"
	"reflect"
	t "github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"net/http"
	i "github.com/vaporz/turbo-example/yourservice/interceptor"
)

func ThriftClient(trans thrift.TTransport, f thrift.TProtocolFactory) interface{} {
	return t.NewYourServiceClientFactory(trans, f)
}

func InitComponents() {
	turbo.Intercept([]string{"GET"}, "/hello", i.LogInterceptor{})
	turbo.RegisterMessageFieldConvertor(new(t.HelloValues), convertHelloValues)
}

func convertHelloValues(req *http.Request) reflect.Value {
	result := &t.HelloValues{}
	result.Message = "a message from convertor"
	return reflect.ValueOf(result)
}
