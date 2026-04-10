This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GatewayRouteHostnameRewrite

An object representing the gateway route host name to rewrite.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultTargetHostname" : String
}

```

### YAML

```yaml

  DefaultTargetHostname: String

```

## Properties

`DefaultTargetHostname`

The default target host name to write to.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayRouteHostnameMatch

GatewayRouteMetadataMatch

All content copied from https://docs.aws.amazon.com/.
