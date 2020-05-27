.. _create:

快速创建一个新的服务
====================

1, 编译和安装命令行工具
-------------------------

在 `release页面 <https://github.com/vaporz/turbo/releases>`_ 下载并安装可执行程序“turbo”和“protoc-gen-buildfields”。

如果你想自己编译，请继续阅读，否则跳到下一步。

先安装 `go <https://golang.org>`_ 和 `protoc <https://developers.google.com/protocol-buffers/>`_ ，
然后下载 `Turbo代码 <https://github.com/vaporz/turbo>`_ ：

.. code-block:: shell

    git clone https://github.com/vaporz/turbo.git

turbo通过“go modules”管理依赖::

 $ cd turbo
 $ go mod tidy
 $ make install
 ...
 $ turbo -v
 turbo version v0.4.0

2, 创建你的服务
----------------

::

 $ turbo create package/path/to/yourservice YourService -r grpc -p ~/

其中，“package/path/to/yourservice”是package名称，“YourService”是gRPC服务的名称，“-r grpc”表示生成一个后端为gRPC服务的项目，
使用“-r thrift”可以生成以thrift服务为后端的项目，“-p ~/”是生成文件存放的位置。

执行成功后，文件夹 "~/package/path/to/yourservice" 会被创建，里面有一些生成的代码。

3, 运行
-------

首先通过“go mod”安装依赖包::

 cd ~/package/path/to/yourservice
 go mod init

好了！运行一个看看！

同时启动 gRPC server 和 HTTP server::

 cd ~/package/path/to/yourservice
 go run main.go

发送请求::

    $ curl -w "\n" "http://localhost:8081/hello?your_name=Alice"
    message:"Hello, Alice"

或者，你也可以分别启动 gRPC server 和 HTTP server::

    $ cd ~/package/path/to/yourservice
    # start grpc service
    $ go run grpcservice/yourservice.go
    # start http server
    $ go run grpcapi/yourserviceapi.go

