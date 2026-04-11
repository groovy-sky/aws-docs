This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewaySpec

An object that represents the specification of a service mesh resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackendDefaults" : VirtualGatewayBackendDefaults,
  "Listeners" : [ VirtualGatewayListener, ... ],
  "Logging" : VirtualGatewayLogging
}

```

### YAML

```yaml

  BackendDefaults:
    VirtualGatewayBackendDefaults
  Listeners:
    - VirtualGatewayListener
  Logging:
    VirtualGatewayLogging

```

## Properties

`BackendDefaults`

A reference to an object that represents the defaults for backends.

_Required_: No

_Type_: [VirtualGatewayBackendDefaults](aws-properties-appmesh-virtualgateway-virtualgatewaybackenddefaults.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Listeners`

The listeners that the mesh endpoint is expected to receive inbound traffic from. You
can specify one listener.

_Required_: Yes

_Type_: Array of [VirtualGatewayListener](aws-properties-appmesh-virtualgateway-virtualgatewaylistener.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

An object that represents logging information.

_Required_: No

_Type_: [VirtualGatewayLogging](aws-properties-appmesh-virtualgateway-virtualgatewaylogging.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayPortMapping

VirtualGatewayTlsValidationContext

All content copied from https://docs.aws.amazon.com/.
