This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GatewayRouteHostnameMatch

An object representing the gateway route host name to match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : String,
  "Suffix" : String
}

```

### YAML

```yaml

  Exact: String
  Suffix: String

```

## Properties

`Exact`

The exact host name to match on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Suffix`

The specified ending characters of the host name to match on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppMesh::GatewayRoute

GatewayRouteHostnameRewrite

All content copied from https://docs.aws.amazon.com/.
