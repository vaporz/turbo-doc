.. _index: 

Turbo
=====

A lightweight microservice tool, turn your grpc|thrift APIs into HTTP APIs!

Features
--------

* Turbo generates a reverse-proxy server which translates a HTTP request into a grpc/Thrift request.
 **(In other words, now you have a grpc|thrift service? Turbo turns your grpc|thrift APIs into HTTP APIs!)**
* Support gRPC and `Thrift <thrift>`_.
* Support `RESTFUL JSON API <json>`_ ("application/json").
* `Interceptor <interceptor>`_.
* `PreProcessor <preprocessor>`_ and `PostProcessor <postprocessor>`_: customizable URL-RPC mapping process.
* `Hijacker <hijacker>`_: Take over requests, do anything you want!
* `Convertor <convertor>`_: Tell Turbo how to set a struct field.
* Modify and reload `configuration <config>`_ file at runtime! Without restarting service.

Requirements
------------
* Golang version: >= 1.8.x
* Thrift version: >= 0.10.0