This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GatewayRouteMetadataMatch

An object representing the method header to be matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : String,
  "Prefix" : String,
  "Range" : GatewayRouteRangeMatch,
  "Regex" : String,
  "Suffix" : String
}

```

### YAML

```yaml

  Exact: String
  Prefix: String
  Range:
    GatewayRouteRangeMatch
  Regex: String
  Suffix: String

```

## Properties

`Exact`

The exact method header to be matched on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The specified beginning characters of the method header to be matched on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Range`

An object that represents the range of values to match on.

_Required_: No

_Type_: [GatewayRouteRangeMatch](aws-properties-appmesh-gatewayroute-gatewayrouterangematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regex`

The regex used to match the method header.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Suffix`

The specified ending characters of the method header to match on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayRouteHostnameRewrite

GatewayRouteRangeMatch

All content copied from https://docs.aws.amazon.com/.
