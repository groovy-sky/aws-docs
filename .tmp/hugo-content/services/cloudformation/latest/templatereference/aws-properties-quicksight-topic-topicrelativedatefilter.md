This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicRelativeDateFilter

A structure that represents a relative date filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Constant" : TopicSingularFilterConstant,
  "RelativeDateFilterFunction" : String,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  Constant:
    TopicSingularFilterConstant
  RelativeDateFilterFunction: String
  TimeGranularity: String

```

## Properties

`Constant`

The constant used in a
relative date filter.

_Required_: No

_Type_: [TopicSingularFilterConstant](aws-properties-quicksight-topic-topicsingularfilterconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDateFilterFunction`

The function to be used in a relative date filter to determine the range of dates to include in the results. Valid values for this structure are `BEFORE`, `AFTER`, and `BETWEEN`.

_Required_: No

_Type_: String

_Allowed values_: `PREVIOUS | THIS | LAST | NEXT | NOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `SECOND | MINUTE | HOUR | DAY | WEEK | MONTH | QUARTER | YEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicRangeFilterConstant

TopicSingularFilterConstant

All content copied from https://docs.aws.amazon.com/.
