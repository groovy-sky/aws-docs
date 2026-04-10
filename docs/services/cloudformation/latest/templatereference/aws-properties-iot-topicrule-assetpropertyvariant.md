This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule AssetPropertyVariant

Contains an asset property value (of a single type).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanValue" : String,
  "DoubleValue" : String,
  "IntegerValue" : String,
  "StringValue" : String
}

```

### YAML

```yaml

  BooleanValue: String
  DoubleValue: String
  IntegerValue: String
  StringValue:
    String

```

## Properties

`BooleanValue`

Optional. A string that contains the boolean value ( `true` or
`false`) of the value entry. Accepts substitution templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

Optional. A string that contains the double value of the value entry. Accepts substitution
templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerValue`

Optional. A string that contains the integer value of the value entry. Accepts
substitution templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

Optional. The string value of the value entry. Accepts substitution templates.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetPropertyValue

BatchConfig

All content copied from https://docs.aws.amazon.com/.
