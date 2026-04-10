This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline Math

An activity that computes an arithmetic expression using the message's attributes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : String,
  "Math" : String,
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  Attribute: String
  Math: String
  Name: String
  Next: String

```

## Properties

`Attribute`

The name of the attribute that contains the result of the math operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Math`

An expression that uses one or more existing attributes and must return an integer value.

_Required_: Yes

_Type_: [String](aws-properties-iotanalytics-pipeline-math.md)

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'math' activity.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Next`

The next activity in the pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Lambda

RemoveAttributes

All content copied from https://docs.aws.amazon.com/.
