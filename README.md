# Grpc实践中的两个问题与思考
## Grpc与Protocol buffer
如果你熟悉Thrift，那这部分内容你可以直接跳过，以后面的内容来说，你可以认为Thrift与gRPC做的是相同的事情。  
要介绍gRPC，首先要知道Protocol Buffer（以下简称protobuf）是什么。  
protobuf是谷歌公布的一种平台无关、编程语言无关的数据序列化和反序列化机制，通过官方对多种语言提供的库和编译器，可以实现跨语言的数据传递，目前支持的编程语言有Java，C++，Python，Ruby，Php，Go等。目前的主要版本有protobuf2和protobuf3。下文讨论中涉及的是protobuf3。  
gRPC也是谷歌公司公布的技术，是一个高性能的，基于HTTP/2协议的RPC框架，默认以protobuf作为数据传输方案。gRPC也提供了对多种编程语言的支持，比如Java，C++，Python，Ruby，Go等。  
换句话说，protobuf首先解决“怎么在不同语言之间传递数据”的问题，而gRPC在这个基础之上，解决了“怎么跨语言进行RPC调用”的问题。  
通过gRPC，你可以很方便的实现一个Client端和Server端，而他们可能是两种不同语言实现的。

接下来，我将介绍两个我们在实际使用gRPC的过程中遇到的问题，以及我们是怎样考虑的。

## 三状态问题
我们提供了一些API给客户使用，并要求将数据以Json格式发给我们，由于业务上的需求，我们需要允许客户在Json中表达这样三种信息：“有key且有值”，“有key但值为null”和“key不存在”。  
例如，某API中需要客户提供“a”，“b”，“c”三个属性的值，而客户发出请求的Json数据可能是：
{
  "b":"xxx",
  "c":null
}
这里表达的含义是“a”不存在，“b”的值是“xxx”，“c”的值是null，在业务逻辑中，它们代表着不同的处理方式。
### 如何表达两种状态：“有key且有值”与“有key但值为null”
首先来看看如何表达两种状态，这可能是多数人使用gRPC或者protobuf3时会遇到的。  
常见的有下面四种办法。
#### 使用oneof
oneof是Protobuf的一个关键字，在官方的介绍中，oneof的用途是：“如果你的message中有很多可选的（optional）属性，并且这些属性在同一时刻最多只有一个有值，那你可以使用oneof功能做到，同时还能节省存储空间。“  
所以，如果被oneof限制的属性只有一个，那表达的含义就等于”这个属性可能有值，也可能没值（相当于null）“。  
比如有message定义如下：
```
message Request {
  oneof body_oneof {
    string body = 1;
  }
}
```
生成的代码中会有方法“GetBodyOneof()”，返回类型是接口，可以通过判断该接口的实际类型是否是“Request_Body”来判断值是否是null，比如：
```
if x, ok := GetBodyOneof().(*Request_Body); ok {
   // body不为null
}
```
可见，虽然功能上能够做到，但首先oneof从设计本意上来说，并不是为了“值可以为null”这种需求，同时，这些代码也显得比较啰嗦，所以总的来说不是一个很好的办法。
#### 使用标记map
在github的一个有很长评论列表的Issue（https://github.com/google/protobuf/issues/1606 ）中，有人提出了这个方案。  
简单来说，就是利用一个map来记录message中哪些属性被赋了值，map的key为属性自己的编号，如果被赋值，则在map中将对应编号key的值值为true。因此，想知道某个属性的值是否是null，就只需要检查这个map中对应的key的值是否为true。  
具体实现上，可以自己写一个protobuf的插件，该插件可以对每个message生成一个map，还可以顺便生成一个HasXXX()方法来方便编码做判断。  
似乎很美，是么？  
注意，protobuf生成的struct的属性都是“导出”的（也就是大写字母开头），相当于Java中的“public”成员变量，而SetXXX()方法更是压根儿没有。这可能是Go语言的风格，好坏在此暂且不论，至少对这个方案来说，运行程序时，需要在每次在给属性赋值后手动去修改这个map，取值前也要专门做判断，这是很容易出错的，所以也不是一个很好的方案。  
当然，你也可以进一步修改代码生成的结果，干脆将属性的名字改成“未导出”的，生成对应的Setter方法，并修改Getter方法的逻辑，然而这样做的代价太大了，也不符合Go的风格，而不符合大家一致遵守的风格的后果，就可能是你的代码与一些第三方库不兼容，这会是一个更让人头疼的问题。
#### 使用wrapper
可能是因为在protobuf3中这个问题太常见了，为了平息“众怒”，谷歌官方“丢”了这么个东西出来：  
https://github.com/google/protobuf/blob/master/src/google/protobuf/wrappers.proto
里面定义了各种基本类型的包装器，有点儿像Java中的int与“Integer”的关系，比如：
```
message Int64Value {
  int64 value = 1;
}
```
你可以用”Int64Value“这个“message”，来替换掉”int64“这个“基本类型”，而message在Go中的零值（zero value）是nil，问题也就解决了。
不管怎么说，与前面两种方案比起来，这个方案至少看起来更合理一些，而且毕竟是谷歌自己提出的方案，官方“钦定”的感觉总是能让人不好拒绝。

前面的三个方案中，多数人最终选择的是第三种，wrapper方案。
### 如何表达第三种状态：“key不存在”
好了，再往前进一步，看看再增加一种状态该怎么办。
#### 在wrapper方案的基础上做扩展
自然而然的，wrapper方案的基础上，如果要做扩展，就会是这样的：
```
message Int64Value {
  int64 value = 1;
  bool exist = 2;
}
```
添加一个布尔类型的属性，为true时表示key存在，否则表示key不存在。可以看到有一些第三方库也是这么做的。  
这么做的优点很明显，容易实现，并且容易理解。而缺点也很明显，这一次，没有一个像谷歌这样的组织来带头提出一份定义文件，所以，你需要自己定义这个“包装器”。而你自己定义的“包装器”是绝对不会有第三方项目支持的，这会给你未来造成一些兼容性方面的不便利，这一点需要你自己去衡量。  
结合我们自己的情况考虑，我们最终选择的是这个方案。
#### 通过额外的数据传递第三种状态
这种思路的目的是，在只使用基本类型的基础上解决“三状态”问题，这样，也就不会有前一种方案中的兼容性问题了。  
该方案最终没有采用，原因是在权衡了复杂度和风险，与比较紧迫的进度安排之后，决定放弃，但我认为仍然值得一提。  
gRPC支持通过Header和Trailer在client和server端传递Metadata，我们可以利用这个功能。  
我们知道，在将用户输入的Json转成protobuf对象的时候，会丢失一些信息，使得无法表示“三种状态”。“丢失的信息”包括两部分，分别是“哪些key的值是null”和“哪些key在Json中不存在”。通过对比输入的Json和protobuf message定义，可以拿到这些丢失的信息，之后就可以利用Metadat将protobuf对象与这些信息一起发给后端，那么，后端得到的实际就是完整的数据了。这时只要再提供一个封装好的方法，将protobuf对象与Metadata中的数据组合到一起，就能够方便的拿到与client端输入的相同的数据。  
从Server端到cient端，反之亦然。  
实现该方案时，相对复杂一点的地方是“对比用户输入Json与message定义，拿到丢失的信息”，需要递归得利用反射来得到，但只要多花些时间，相信并不难。  
与前一种方案相比，这个方案的优点是不定义新的基本类型，至于缺点，是给不同服务间通过gRPC调用增加了额外的复杂度，每一个服务都需要使用一个专门的方法来调用其他服务的接口，其中封装了前面提到的对比和组合过程。

## 如何扩展server端返回的Error
在gRPC的server端返回error时，虽然接口的返回值中声明的是标准库的error，但gRPC内部会判断error是否特定的类型，如果不是，则会统一返回Code为“Unknown”的error。因此，多数情况下我们需要用专门的方法构造一个error，比如：
```
import "google.golang.org/grpc/status"
...
return status.Error(codes.Aborted, "aborted")
```
实际上，这个error的类型是“google.golang.org/genproto/googleapis/rpc/status” 包里面的“Status”，定义如下：
```
type Status struct {
	Code int32
	Message string
	Details []*google_protobuf.Any
}
```
除了Details字段外，只有Code和Message，十分简洁，但如果我们有更复杂的业务需求，这些字段是不能满足的，因此我们需要扩展Error。
### 不要扩展grpc预定义的status code
前面例子中，“codes.Aborted”是gRPC预定义的一些Code之一，很自然的，我们会想到，是否可以扩展这个Code列表。  
我曾经尝试过扩展这个列表，在Go语言的gRPC客户端中是可以拿到自定义的Code的，但官方的开发人员在一些问答中表示不推荐这么做，而且一些其他语言的gRPC客户端中，一旦发现Code不在预定义的列表中，有可能直接替换成预定义的“Unknown”错误，甚至直接抛出异常，因此，不要这么做。
### 在response中添加err属性
这也是一种显而易见的方案，既然已有的error不能满足需求，那就在Response对象中加入一个“err”属性，而它的类型是自己定义的，大概是这个样子：
```
type MyResponse struct {
  err MyErr
  ...
}
```
可以满足需求，但很不优雅，同意么？这么做直接违反了gRPC的错误处理机制，甚至不符合Go语言的规范，所以也不推荐这么做。
### 通过metadata传递
首先要说，这才是一个靠谱的方案，也是官方曾经推荐的方案。  
先不说“曾经”是什么意思，我们来看看是这个方案是怎么玩的。  
gRPC的Client端和Server端之间，可以借助名为“Metadata”的数据结构来传递额外的信息，而我们自己扩展的error信息就属于这个“额外信息”。  
详细的用法可以参考在github上的文档：  
https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md  
这里将其中的一些代码示例贴在这里，以便让大家快速的有一个直观的印象。  
以从Server端向Client端传递Metadata为例，首先在Server端将数据准备好：
```
func (s *server) SomeRPC(ctx context.Context, in *pb.someRequest) (*pb.someResponse, error) {
	   // 创建并发送Header
	   header := metadata.Pairs("header-key", "val")
     grpc.SendHeader(ctx, header)
     // 创建并发送Trailer
     trailer := metadata.Pairs("trailer-key", "val")
     grpc.SetTrailer(ctx, trailer)
}
```
接着在Client端接收数据：
```
var header, trailer metadata.MD // 用来保存header和trailer的变量
r, err := client.SomeRPC(
    ctx,
    someRequest,
    grpc.Header(&header),    // 将会接收header
    grpc.Trailer(&trailer),  // 将会接收trailer
)

// 按需求对header和trailer做处理
```
### 通过Status中的Details属性传递
最后出场的是我们实际采用的方案。  
我在前面将扩展的error信息称作“额外信息”，其实准确的说应该是“额外的error信息”，因此，直觉上最自然的方式还是在error对象内部携带这个信息。  
你一定注意到了“Status“对象里的那个”Details“属性，它是一个“Any”对象的数组，字面上看似乎是“可以保存任何类型对象的数组”的意思，确实是这样。  
gRPC最近刚刚添加了两个工具方法，使得“Details”属性变得非常易用：
```
// WithDetails返回一个新创建的Status对象，其中附加了参数details传入的Message列表，
// 如果有error发生，将返回nil和第一个遇到的error。
func (s *Status) WithDetails(details ...proto.Message) (*Status, error)

// Details返回Status中Details携带的Message列表，
// 如果decode某个Message时发生错误，那这个错误会被添加到结果列表中返回。
func (s *Status) Details() []interface{} 
```
实际使用时很方便。  
附加Details：
```
s, _ := status.New(codes.Aborted, "message").WithDetails(d1, d2)
return nil, s.Err()
```
获取Details：
```
details, _ := s.Details()
for _, d := range details {
  m := d.(YourType)
  // ...
}
```

好了，以上就是我们在使用gRPC过程中遇到的两个问题和相应的思考，也许并不是最优的方案，但希望能给你带来一些提示。

Go语言作为一种快速发展变化中的语言，相应的技术生态还不是十分健全，包括gRPC和由此衍生的gRPC-gateway等项目，仍然有不少提升空间。中国作为目前Go语言应用人数最多，气氛最火热的国家，希望能看到越来越多的国内开发者参与到开源项目的发展中，也希望有越来越多的优秀项目出现。

