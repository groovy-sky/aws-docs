This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PercentVisibleRange

The percent range in the visible range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "From" : Number,
  "To" : Number
}

```

### YAML

```yaml

  From: Number
  To: Number

```

## Properties

`From`

The lower bound of the range.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`To`

The top bound of the range.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PercentileAggregation

PeriodOverPeriodComputation

All content copied from https://docs.aws.amazon.com/.
