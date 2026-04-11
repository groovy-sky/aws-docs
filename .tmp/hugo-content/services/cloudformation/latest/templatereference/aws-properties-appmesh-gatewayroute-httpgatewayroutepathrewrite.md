This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute HttpGatewayRoutePathRewrite

An object that represents the path to rewrite.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : String
}

```

### YAML

```yaml

  Exact: String

```

## Properties

`Exact`

The exact path to rewrite.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpGatewayRouteMatch

HttpGatewayRoutePrefixRewrite

All content copied from https://docs.aws.amazon.com/.
