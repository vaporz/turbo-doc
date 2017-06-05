package gen

import (
	"github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"github.com/vaporz/turbo"
	"reflect"
	"net/http"
	"errors"
)

/*
this is a generated file, DO NOT EDIT!
 */
var ThriftSwitcher = func(methodName string, resp http.ResponseWriter, req *http.Request) (serviceResponse interface{}, err error) {
	switch methodName { 
	case "SayHello":
		args := gen.YourServiceSayHelloArgs{}
		params, err := turbo.BuildArgs(reflect.TypeOf(args), reflect.ValueOf(args), req, buildStructArg)
		if err != nil {
			return nil, err
		}
		return turbo.ThriftService().(*gen.YourServiceClient).SayHello(
			params[0].Interface().(string),
			params[1].Interface().(*gen.CommonValues),
			params[2].Interface().(*gen.HelloValues), )
	case "EatApple":
		args := gen.YourServiceEatAppleArgs{}
		params, err := turbo.BuildArgs(reflect.TypeOf(args), reflect.ValueOf(args), req, buildStructArg)
		if err != nil {
			return nil, err
		}
		return turbo.ThriftService().(*gen.YourServiceClient).EatApple(
			params[0].Interface().(int32),
			params[1].Interface().(string),
			params[2].Interface().(bool), )
	case "TestProto":
		return turbo.ThriftService().(*gen.YourServiceClient).TestProto( )
	default:
		return nil, errors.New("No such method[" + methodName + "]")
	}
}

func buildStructArg(typeName string, req *http.Request) (v reflect.Value, err error) {
	switch typeName { 
	case "CommonValues":
		request := &gen.CommonValues{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return v, err
		}
		return reflect.ValueOf(request), nil
	case "HelloValues":
		request := &gen.HelloValues{  }
		err = turbo.BuildStruct(reflect.TypeOf(request).Elem(), reflect.ValueOf(request).Elem(), req)
		if err != nil {
			return v, err
		}
		return reflect.ValueOf(request), nil
	default:
		return v, errors.New("unknown typeName[" + typeName + "]")
	}
}
