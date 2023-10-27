# Name resolving(命名解析)

This examples shows how `ClientConn` can pick different name resolvers.
该示例展示了`ClientConn`如何选择不同的命名解析器(name resolvers)。

## What is a name resolver(命名解析器)

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.
一个命名解析器可以被视为一个(以service name作为键，以实际后端服务ip地址列表作为值的)map映射`map[service-name][]backend-ip`。
它可以通过一个服务名，来获取该服务名对应的实际后端服务的IP地址列表。DNS就是一个通用的命名解析器。

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.
在该示例中，创建了一个用来将`resolver.example.grpc.io`解析为`localhost:50051`的命名解析器。

## Try it

```
go run server/main.go
```

```
go run client/main.go
```

## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.