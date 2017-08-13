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
#### 使用指针类型
那么判断其中
            使用oneof
            使用标记map
            使用指针类型
            使用wrapper
            ??  github.com/golang/protobuf/ptypes/wrappers/wrappers.proto
            
      2.2 如何表达第三种状态：“key不存在”
	    在wrapper的基础上做扩展
                  优点和缺点
            通过额外的数据结构传递缺失的信息
                  优点和缺点
  3.  如何自定义error信息
      3.1 不要扩展grpc预定义的status code
      3.2 方案一：在response中添加err属性
      3.3 方案二：通过metadata传递
      3.4 方案三：通过Status中的details属性传递
      3.5 总结
