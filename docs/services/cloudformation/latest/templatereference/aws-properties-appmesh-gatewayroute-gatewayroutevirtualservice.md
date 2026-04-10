This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GatewayRouteVirtualService

An object that represents the virtual service that traffic is routed to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VirtualServiceName" : String
}

```

### YAML

```yaml

  VirtualServiceName: String

```

## Properties

`VirtualServiceName`

The name of the virtual service that traffic is routed to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayRouteTarget

GrpcGatewayRoute

All content copied from https://docs.aws.amazon.com/.
