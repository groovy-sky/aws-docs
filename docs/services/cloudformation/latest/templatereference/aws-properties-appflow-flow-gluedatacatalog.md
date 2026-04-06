This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow GlueDataCatalog

The `GlueDataCatalog` property type specifies Property description not available. for an [AWS::AppFlow::Flow](aws-resource-appflow-flow.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "RoleArn" : String,
  "TablePrefix" : String
}

```

### YAML

```yaml

  DatabaseName: String
  RoleArn: String
  TablePrefix: String

```

## Properties

`DatabaseName`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:iam:.*:[0-9]+:.*`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TablePrefix`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridgeDestinationProperties

GoogleAnalyticsSourceProperties
