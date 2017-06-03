package impl

import (
	"fmt"
	"github.com/vaporz/turbo-example/yourservice/gen/thrift/gen-go/gen"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func TProcessor() thrift.TProcessor {
	return gen.NewYourServiceProcessor(YourService{})
}

type YourService struct {
}

func (s YourService) SayHello(yourName string, values *gen.CommonValues, helloValues *gen.HelloValues) (r *gen.SayHelloResponse, err error) {
	fmt.Println(values.TransactionId)
	fmt.Println(helloValues.Message)
	return &gen.SayHelloResponse{Message: "[thrift server]Hello, " + yourName}, nil
}

func (s YourService) EatApple(num int32, stringValue string, boolValue bool) (r *gen.EatAppleResponse, err error) {
	msg := fmt.Sprintf("[thrift server]Good taste! Apple num=%d, string=%s, bool=%t", num, stringValue, boolValue)
	return &gen.EatAppleResponse{Message: msg}, nil
}

func (s YourService) TestProto() (r *gen.TestProtoResponse, err error) {
	resp := &gen.TestProtoResponse{}
	resp.Int64Slice = []int64{11, 22, 33}
	resp.Values = &gen.CommonValues{TransactionId: 111}
	resp.StringValue = "string value!"
	resp.C1 = &gen.Child{ChildValue: 222}
	resp.C2 = &gen.Child{}
	resp.Int32Value = 333
	resp.Int64Value = 555
	resp.Childs = []*gen.Child{}
	return resp, nil
}
