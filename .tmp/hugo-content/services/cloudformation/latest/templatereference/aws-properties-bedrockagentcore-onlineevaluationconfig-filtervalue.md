This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig FilterValue

The value to compare against using the specified operator. Can be a string, double, or boolean value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanValue" : Boolean,
  "DoubleValue" : Number,
  "StringValue" : String
}

```

### YAML

```yaml

  BooleanValue:
    Boolean
  DoubleValue: Number
  StringValue:
    String

```

## Properties

`BooleanValue`

The boolean value for true/false filtering conditions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

The numeric value for numerical filtering and comparisons.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

The string value for text-based filtering.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

OutputConfig

All content copied from https://docs.aws.amazon.com/.
