This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayConnectionPool

An object that represents the type of virtual gateway connection pool.

Only one protocol is used at a time and should be the same protocol as the one chosen
under port mapping.

If not present the default value for `maxPendingRequests` is
`2147483647`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GRPC" : VirtualGatewayGrpcConnectionPool,
  "HTTP" : VirtualGatewayHttpConnectionPool,
  "HTTP2" : VirtualGatewayHttp2ConnectionPool
}

```

### YAML

```yaml

  GRPC:
    VirtualGatewayGrpcConnectionPool
  HTTP:
    VirtualGatewayHttpConnectionPool
  HTTP2:
    VirtualGatewayHttp2ConnectionPool

```

## Properties

`GRPC`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualGatewayGrpcConnectionPool](aws-properties-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualGatewayHttpConnectionPool](aws-properties-appmesh-virtualgateway-virtualgatewayhttpconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP2`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualGatewayHttp2ConnectionPool](aws-properties-appmesh-virtualgateway-virtualgatewayhttp2connectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayClientTlsCertificate

VirtualGatewayFileAccessLog

All content copied from https://docs.aws.amazon.com/.
