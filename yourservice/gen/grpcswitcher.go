package gen

import (
	g "github.com/vaporz/turbo-example/yourservice/gen/proto"
	"github.com/vaporz/turbo"
	"reflect"
	"net/http"
	"errors"
)

/*
this is a generated file, DO NOT EDIT!
 */
var GrpcSwitcher = func(methodName string, resp http.ResponseWriter, req *http.Request) (serviceResponse interface{}, err error) {
	switch methodName { 
	case "SayHello":
		request := &g.SayHelloRequest{ Values: &g.CommonValues{}, }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return nil, err
		}
		return turbo.GrpcService().(g.YourServiceClient).SayHello(req.Context(), request)
	case "EatApple":
		request := &g.EatAppleRequest{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return nil, err
		}
		return turbo.GrpcService().(g.YourServiceClient).EatApple(req.Context(), request)
	case "TestProto":
		request := &g.TestProtoRequest{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return nil, err
		}
		return turbo.GrpcService().(g.YourServiceClient).TestProto(req.Context(), request)
	default:
		return nil, errors.New("No such method[" + methodName + "]")
	}
}
