This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ListenerTimeout

An object that represents timeouts for different protocols.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GRPC" : GrpcTimeout,
  "HTTP" : HttpTimeout,
  "HTTP2" : HttpTimeout,
  "TCP" : TcpTimeout
}

```

### YAML

```yaml

  GRPC:
    GrpcTimeout
  HTTP:
    HttpTimeout
  HTTP2:
    HttpTimeout
  TCP:
    TcpTimeout

```

## Properties

`GRPC`

An object that represents types of timeouts.

_Required_: No

_Type_: [GrpcTimeout](aws-properties-appmesh-virtualnode-grpctimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP`

An object that represents types of timeouts.

_Required_: No

_Type_: [HttpTimeout](aws-properties-appmesh-virtualnode-httptimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTP2`

An object that represents types of timeouts.

_Required_: No

_Type_: [HttpTimeout](aws-properties-appmesh-virtualnode-httptimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TCP`

An object that represents types of timeouts.

_Required_: No

_Type_: [TcpTimeout](aws-properties-appmesh-virtualnode-tcptimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Listener

ListenerTls

All content copied from https://docs.aws.amazon.com/.
