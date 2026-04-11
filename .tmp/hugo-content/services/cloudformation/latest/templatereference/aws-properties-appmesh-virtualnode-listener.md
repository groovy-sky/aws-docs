This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode Listener

An object that represents a listener for a virtual node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionPool" : VirtualNodeConnectionPool,
  "HealthCheck" : HealthCheck,
  "OutlierDetection" : OutlierDetection,
  "PortMapping" : PortMapping,
  "Timeout" : ListenerTimeout,
  "TLS" : ListenerTls
}

```

### YAML

```yaml

  ConnectionPool:
    VirtualNodeConnectionPool
  HealthCheck:
    HealthCheck
  OutlierDetection:
    OutlierDetection
  PortMapping:
    PortMapping
  Timeout:
    ListenerTimeout
  TLS:
    ListenerTls

```

## Properties

`ConnectionPool`

The connection pool information for the listener.

_Required_: No

_Type_: [VirtualNodeConnectionPool](aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheck`

The health check information for the listener.

_Required_: No

_Type_: [HealthCheck](aws-properties-appmesh-virtualnode-healthcheck.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutlierDetection`

The outlier detection information for the listener.

_Required_: No

_Type_: [OutlierDetection](aws-properties-appmesh-virtualnode-outlierdetection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortMapping`

The port mapping information for the listener.

_Required_: Yes

_Type_: [PortMapping](aws-properties-appmesh-virtualnode-portmapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

An object that represents timeouts for different protocols.

_Required_: No

_Type_: [ListenerTimeout](aws-properties-appmesh-virtualnode-listenertimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLS`

A reference to an object that represents the Transport Layer Security (TLS) properties for a listener.

_Required_: No

_Type_: [ListenerTls](aws-properties-appmesh-virtualnode-listenertls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JsonFormatRef

ListenerTimeout

All content copied from https://docs.aws.amazon.com/.
