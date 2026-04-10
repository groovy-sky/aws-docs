This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLensGroup MatchObjectAge

This resource contains `DaysGreaterThan` and `DaysLessThan` to
define the object age range (minimum and maximum number of days).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DaysGreaterThan" : Integer,
  "DaysLessThan" : Integer
}

```

### YAML

```yaml

  DaysGreaterThan: Integer
  DaysLessThan: Integer

```

## Properties

`DaysGreaterThan`

This property indicates the minimum object age in days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DaysLessThan`

This property indicates the maximum object age in days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

MatchObjectSize

All content copied from https://docs.aws.amazon.com/.
