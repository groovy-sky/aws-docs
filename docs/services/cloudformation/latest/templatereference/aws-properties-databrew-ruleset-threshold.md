This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Ruleset Threshold

The threshold used with a non-aggregate check expression. The non-aggregate check
expression will be applied to each row in a specific column. Then the threshold will be
used to determine whether the validation succeeds.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Unit" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Type: String
  Unit: String
  Value: Number

```

## Properties

`Type`

The type of a threshold. Used for comparison of an actual count of rows that satisfy
the rule to the threshold value.

_Required_: No

_Type_: String

_Allowed values_: `GREATER_THAN_OR_EQUAL | LESS_THAN_OR_EQUAL | GREATER_THAN | LESS_THAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

Unit of threshold value. Can be either a COUNT or PERCENTAGE of the full sample size
used for validation.

_Required_: No

_Type_: String

_Allowed values_: `COUNT | PERCENTAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a threshold.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::DataBrew::Schedule

All content copied from https://docs.aws.amazon.com/.
