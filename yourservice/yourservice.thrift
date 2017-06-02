namespace go gen
include "shared.thrift"

struct SayHelloResponse {
  1: string message,
}

struct EatAppleResponse {
  1: string message,
}

struct Child {
    1: i64 childValue,
}

struct TestProtoResponse {
    1:shared.CommonValues values,
    2:string stringValue,
    3:i32 int32Value,
    4:i64 int64Value,
    5:Child c1,
    6:Child c2,
    7:list<i64> int64Slice,
    8:list<Child> childs,
}

service YourService {
    SayHelloResponse sayHello (1:string yourName, 2:shared.CommonValues values, 3:shared.HelloValues helloValues)
    EatAppleResponse eatApple (1:i32 num, 2:string stringValue, 3:bool boolValue)
    TestProtoResponse testProto ()
}
