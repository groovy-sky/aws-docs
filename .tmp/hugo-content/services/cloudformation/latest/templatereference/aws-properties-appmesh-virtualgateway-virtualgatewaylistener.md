This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayListener

An object that represents a listener for a virtual gateway.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionPool" : VirtualGatewayConnectionPool,
  "HealthCheck" : VirtualGatewayHealthCheckPolicy,
  "PortMapping" : VirtualGatewayPortMapping,
  "TLS" : VirtualGatewayListenerTls
}

```

### YAML

```yaml

  ConnectionPool:
    VirtualGatewayConnectionPool
  HealthCheck:
    VirtualGatewayHealthCheckPolicy
  PortMapping:
    VirtualGatewayPortMapping
  TLS:
    VirtualGatewayListenerTls

```

## Properties

`ConnectionPool`

The connection pool information for the listener.

_Required_: No

_Type_: [VirtualGatewayConnectionPool](aws-properties-appmesh-virtualgateway-virtualgatewayconnectionpool.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheck`

The health check information for the listener.

_Required_: No

_Type_: [VirtualGatewayHealthCheckPolicy](aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortMapping`

The port mapping information for the listener.

_Required_: Yes

_Type_: [VirtualGatewayPortMapping](aws-properties-appmesh-virtualgateway-virtualgatewayportmapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLS`

A reference to an object that represents the Transport Layer Security (TLS) properties for the listener.

_Required_: No

_Type_: [VirtualGatewayListenerTls](aws-properties-appmesh-virtualgateway-virtualgatewaylistenertls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayHttpConnectionPool

VirtualGatewayListenerTls

All content copied from https://docs.aws.amazon.com/.
