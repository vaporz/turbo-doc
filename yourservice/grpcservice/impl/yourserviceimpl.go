package impl

import (
	"golang.org/x/net/context"
	"strconv"
	"github.com/vaporz/turbo-example/yourservice/gen/proto"
	"google.golang.org/grpc"
)

func RegisterServer(s *grpc.Server) {
	proto.RegisterYourServiceServer(s, &YourService{})
}

type YourService struct {
}

func (s *YourService) SayHello(ctx context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	someId := strconv.FormatInt(req.Values.SomeId, 10)
	return &proto.SayHelloResponse{Message: "[grpc server]Hello, " + req.YourName + ", someId=" + someId}, nil
}

func (s *YourService) EatApple(ctx context.Context, req *proto.EatAppleRequest) (*proto.EatAppleResponse, error) {
	return &proto.EatAppleResponse{Message: "Good taste! Apple num=" + strconv.FormatInt(int64(req.Num), 10)}, nil
}

func (s *YourService) TestProto(ctx context.Context, req *proto.TestProtoRequest) (*proto.TestProtoResponse, error) {
	resp := &proto.TestProtoResponse{}
	resp.Int64Slice = []int64{11, 22, 33}
	resp.Values = &proto.CommonValues{SomeId: 111}
	resp.StringValue = "string value!"
	resp.C1 = &proto.TestProtoResponseChild{ChildValue: 222}
	resp.C2 = &proto.TestProtoResponseChild{}
	resp.Int32Value = 333
	resp.Int64Value = 555
	resp.Childs = []*proto.TestProtoResponseChild{}
	return resp, nil
}
