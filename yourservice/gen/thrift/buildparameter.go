package main

import (
	g "github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"reflect"
	"fmt"
	"flag"
)

var methodName = flag.String("n", "", "")

func main() {
	flag.Parse()
	str := buildParameterStr(*methodName)
	fmt.Print(str)
}

func buildParameterStr(methodName string) string {
	switch methodName { 
	case "SayHello":
		var result string
		args := g.YourServiceSayHelloArgs{}
		at := reflect.TypeOf(args)
		num := at.NumField()
		for i := 0; i < num; i++ {
			result += fmt.Sprintf(
			"\n\t\t\tparams[%d].Interface().(%s),",
			i, at.Field(i).Type.String())
		}
		return result
	case "EatApple":
		var result string
		args := g.YourServiceEatAppleArgs{}
		at := reflect.TypeOf(args)
		num := at.NumField()
		for i := 0; i < num; i++ {
			result += fmt.Sprintf(
			"\n\t\t\tparams[%d].Interface().(%s),",
			i, at.Field(i).Type.String())
		}
		return result
	case "TestProto":
		var result string
		args := g.YourServiceTestProtoArgs{}
		at := reflect.TypeOf(args)
		num := at.NumField()
		for i := 0; i < num; i++ {
			result += fmt.Sprintf(
			"\n\t\t\tparams[%d].Interface().(%s),",
			i, at.Field(i).Type.String())
		}
		return result
	default:
		return "error"
	}
}
