This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute HttpGatewayRoutePrefixRewrite

An object representing the beginning characters of the route to rewrite.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultPrefix" : String,
  "Value" : String
}

```

### YAML

```yaml

  DefaultPrefix: String
  Value: String

```

## Properties

`DefaultPrefix`

The default prefix used to replace the incoming route prefix when rewritten.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value used to replace the incoming route prefix when rewritten.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpGatewayRoutePathRewrite

HttpGatewayRouteRewrite

All content copied from https://docs.aws.amazon.com/.
