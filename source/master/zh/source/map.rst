.. _map:

HTTP接口与后端接口映射
======================

前后端接口的映射关系通过配置文件中的"urlmapping"节点保存。

在 `前面的例子 <add>`_ 中，有如下的配置:

.. code-block:: yaml

	urlmapping:
 	  - GET /hello SayHello
	  - GET /eat_apple/{num:[0-9]+} EatApple

在这段配置中，发送给HTTP接口"GET /hello"的请求将转发到后端接口"SayHello"，
而发送给HTTP接口"GET /eat_apple/{num:[0-9]+}"的请求将转发给后端接口"EatApple"，而且在URL中定义了一个叫做"num"的参数，
参数的格式为"[0-9]+"，很明显，这个正则表达式的意思是数字。

在Turbo内部，URL路由和正则功能是由好用又强大的"github.com/gorilla/mux"提供的，因此，关于URL路由功能，请参考 `该项目的文档 <https://github.com/gorilla/mux>`_ 。

HTTP接口与后端接口的映射关系可以是多对一的，也就是可以多个HTTP接口映射到同一个后端接口上，比如:

.. code-block:: yaml

	urlmapping:
 	  - GET /hello SayHello
 	  - POST /hello SayHello
 	  - GET /new/hello SayHello
	  - GET /eat_apple/{num:[0-9]+} EatApple

这样，两个新接口"POST /hello"和"GET /new/hello"都将被映射到"SayHello"上。

如果多个HTTP接口只是谓词不同，而URL是相同的，那也可以写在同一行，比如:

.. code-block:: yaml

	urlmapping:
 	  - GET,POST,UPDATE /hello SayHello
 	  - GET /new/hello SayHello
	  - GET /eat_apple/{num:[0-9]+} EatApple

"GET,POST,UPDATE /hello"表示了三个HTTP接口，他们的URL相同，但谓词不同，并且都被映射到后端接口"SayHello"。
