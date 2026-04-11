This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode VirtualNodeConnectionPool

An object that represents the type of virtual node connection pool.

Only one protocol is used at a time and should be the same protocol as the one chosen
under port mapping.

If not present the default value for `maxPendingRequests` is
`2147483647`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GRPC" : VirtualNodeGrpcConnectionPool,
  "HTTP" : VirtualNodeHttpConnectionPool,
  "HTTP2" : VirtualNodeHttp2ConnectionPool,
  "TCP" : VirtualNodeTcpConnectionPool
}

```

### YAML

```yaml

  GRPC:
    VirtualNodeGrpcConnectionPool
  HTTP:
    VirtualNodeHttpConnectionPool
  HTTP2:
    VirtualNodeHttp2ConnectionPool
  TCP:
    VirtualNodeTcpConnectionPool

```

## Properties

`GRPC`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualNodeGrpcConnectionPool](aws-properties-appmesh-virtualnode-virtualnodegrpcconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualNodeHttpConnectionPool](aws-properties-appmesh-virtualnode-virtualnodehttpconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP2`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualNodeHttp2ConnectionPool](aws-properties-appmesh-virtualnode-virtualnodehttp2connectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TCP`

An object that represents a type of connection pool.

_Required_: No

_Type_: [VirtualNodeTcpConnectionPool](aws-properties-appmesh-virtualnode-virtualnodetcpconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TlsValidationContextTrust

VirtualNodeGrpcConnectionPool

All content copied from https://docs.aws.amazon.com/.
