.. _create:

Create a service on the fly
===========================

1, Compile and install Turbo Command tools
------------------------------------------

Download executable binary files "turbo" and "protoc-gen-buildfields" at  `release page <https://github.com/vaporz/turbo/releases>`_ .

Continue reading if you want to compile by yourself, or else jump to next step.

First, install `go <https://golang.org>`_ 和 `protoc <https://developers.google.com/protocol-buffers/>`_ ,
then download `Turbo repo <https://github.com/vaporz/turbo>`_ ：

.. code-block:: shell

    git clone https://github.com/vaporz/turbo.git

Turbo manages dependencies with "go modules"::

 $ cd turbo
 $ go mod tidy
 $ make install
 ...
 $ turbo -v
 turbo version v0.4.0

2, Create your service
----------------------

::

 $ turbo create package/path/to/yourservice YourService -r grpc -p ~/

"package/path/to/yourservice" is package name, "YourService" is backend service's name, "-r grpc" means to create
a gRPC service as backend service, or you can create a Thrift service as backend service using "-r thrift",
"-p ~/" shows the directory in which new project will be created.

Once succeed, directory "~/package/path/to/yourservice" is created, along with some generated codes.

3, Run
------

Install dependencies with "go mod"::

 cd ~/package/path/to/yourservice
 go mod init

That's it! Now let's Play!

Start both gRPC server and HTTP server::

 cd ~/package/path/to/yourservice
 go run main.go

Send a request::

    $ curl -w "\n" "http://localhost:8081/hello?your_name=Alice"
    message:"Hello, Alice"

Or you can start gRPC server and HTTP server separately::

    $ cd ~/package/path/to/yourservice
    # start grpc service
    $ go run grpcservice/yourservice.go
    # start http server
    $ go run grpcapi/yourserviceapi.go

