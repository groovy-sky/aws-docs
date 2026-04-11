This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode HealthCheck

An object that represents the health check policy for a virtual node's listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthyThreshold" : Integer,
  "IntervalMillis" : Integer,
  "Path" : String,
  "Port" : Integer,
  "Protocol" : String,
  "TimeoutMillis" : Integer,
  "UnhealthyThreshold" : Integer
}

```

### YAML

```yaml

  HealthyThreshold: Integer
  IntervalMillis: Integer
  Path: String
  Port: Integer
  Protocol: String
  TimeoutMillis: Integer
  UnhealthyThreshold: Integer

```

## Properties

`HealthyThreshold`

The number of consecutive successful health checks that must occur before declaring
listener healthy.

_Required_: Yes

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalMillis`

The time period in milliseconds between each health check execution.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The destination path for the health check request. This value is only used if the
specified protocol is HTTP or HTTP/2. For any other protocol, this value is ignored.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The destination port for the health check request. This port must match the port defined
in the [PortMapping](../userguide/aws-properties-appmesh-virtualnode-listener.md#cfn-appmesh-virtualnode-listener-portmapping) for the listener.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol for the health check request. If you specify `grpc`, then your
service must conform to the [GRPC Health\
Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

_Required_: Yes

_Type_: String

_Allowed values_: `http | tcp | http2 | grpc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMillis`

The amount of time to wait when receiving a response from the health check, in
milliseconds.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyThreshold`

The number of consecutive failed health checks that must occur before declaring a
virtual node unhealthy.

_Required_: Yes

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcTimeout

HttpTimeout

All content copied from https://docs.aws.amazon.com/.
