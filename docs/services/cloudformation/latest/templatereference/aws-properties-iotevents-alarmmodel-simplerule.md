This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel SimpleRule

A rule that compares an input property value to a threshold value with a comparison operator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "InputProperty" : String,
  "Threshold" : String
}

```

### YAML

```yaml

  ComparisonOperator: String
  InputProperty: String
  Threshold: String

```

## Properties

`ComparisonOperator`

The comparison operator.

_Required_: Yes

_Type_: String

_Allowed values_: `GREATER | GREATER_OR_EQUAL | LESS | LESS_OR_EQUAL | EQUAL | NOT_EQUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputProperty`

The value on the left side of the comparison operator. You can specify an AWS IoT Events input
attribute as an input property.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The value on the right side of the comparison operator. You can enter a number or specify
an AWS IoT Events input attribute.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Payload

Sns

All content copied from https://docs.aws.amazon.com/.
