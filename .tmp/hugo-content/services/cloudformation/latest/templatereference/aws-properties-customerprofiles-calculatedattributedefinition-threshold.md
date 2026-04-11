This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition Threshold

The threshold for the calculated attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Value" : String
}

```

### YAML

```yaml

  Operator: String
  Value: String

```

## Properties

`Operator`

The operator of the threshold.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUAL_TO | GREATER_THAN | LESS_THAN | NOT_EQUAL_TO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the threshold.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ValueRange

All content copied from https://docs.aws.amazon.com/.
