.. _map:

HTTP API Mapping
================

Mapping between HTTP API and backend API is stored under node "urlmapping" in config file.

From `previous example <add>`_ , we have a config like this:

.. code-block:: yaml

	urlmapping:
 	  - GET /hello YourService SayHello
	  - GET /eat_apple/{num:[0-9]+} YourService EatApple

As this config says, any request sent to HTTP API "GET /hello" is redirected to backend API "SayHello".
And any request sent to HTTP API "GET /eat_apple/{num:[0-9]+}" is redirected to backend API "EatApple",
along with a parameter "num" defined in URL, the format to "num" is "[0-9]+", which means, "num" is a number.

Inside Turbo, URL routing functionality and regex support is provided by easy-to-use and powerful package "github.com/gorilla/mux".
Please refer to `mux's documentation <https://github.com/gorilla/mux>`_ for more details.

Mapping between HTTP API and backend API can be N:1, for example:

.. code-block:: yaml

	urlmapping:
 	  - GET /hello YourService SayHello
 	  - POST /hello YourService SayHello
 	  - GET /new/hello YourService SayHello
	  - GET /eat_apple/{num:[0-9]+} YourService EatApple

So, new HTTP API "POST /hello" and "GET /new/hello" are both mapped to "SayHello".

If the difference between multiple HTTP APIs is the request method only (GET,POST, etc), the URL are same,
then you can also write like this:

.. code-block:: yaml

	urlmapping:
 	  - GET,POST,UPDATE /hello YourService SayHello
 	  - GET /new/hello YourService SayHello
	  - GET /eat_apple/{num:[0-9]+} YourService EatApple

"GET,POST,UPDATE /hello" presents 3 HTTP APIs, their URL are same, but with different request method, and they
are all mapped to backend API "SayHello".
