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
		params := turbo.MakeParams(req, reflect.ValueOf(request))
		return turbo.ParseResult(callGrpcMethod(methodName, params))
	case "EatApple":
		request := &g.EatAppleRequest{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return nil, err
		}
		params := turbo.MakeParams(req, reflect.ValueOf(request))
		return turbo.ParseResult(callGrpcMethod(methodName, params))
	case "TestProto":
		request := &g.TestProtoRequest{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return nil, err
		}
		params := turbo.MakeParams(req, reflect.ValueOf(request))
		return turbo.ParseResult(callGrpcMethod(methodName, params))
	default:
		return nil, errors.New("No such method[" + methodName + "]")
	}
}

func callGrpcMethod(methodName string, params []reflect.Value) []reflect.Value {
	return reflect.ValueOf(turbo.GrpcService().(g.YourServiceClient)).MethodByName(methodName).Call(params)
}
