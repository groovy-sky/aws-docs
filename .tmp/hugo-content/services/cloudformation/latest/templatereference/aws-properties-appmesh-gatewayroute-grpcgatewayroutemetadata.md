This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GrpcGatewayRouteMetadata

An object representing the metadata of the gateway route.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Invert" : Boolean,
  "Match" : GatewayRouteMetadataMatch,
  "Name" : String
}

```

### YAML

```yaml

  Invert: Boolean
  Match:
    GatewayRouteMetadataMatch
  Name: String

```

## Properties

`Invert`

Specify `True` to match anything except the match criteria. The default value
is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

The criteria for determining a metadata match.

_Required_: No

_Type_: [GatewayRouteMetadataMatch](aws-properties-appmesh-gatewayroute-gatewayroutemetadatamatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the gateway route metadata.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcGatewayRouteMatch

GrpcGatewayRouteRewrite

All content copied from https://docs.aws.amazon.com/.
