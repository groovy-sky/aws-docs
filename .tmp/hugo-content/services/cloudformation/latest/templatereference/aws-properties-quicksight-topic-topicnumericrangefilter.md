This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicNumericRangeFilter

A filter that filters topics based on the value of a numeric field. The filter includes only topics whose numeric field value falls within the specified range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : String,
  "Constant" : TopicRangeFilterConstant,
  "Inclusive" : Boolean
}

```

### YAML

```yaml

  Aggregation: String
  Constant:
    TopicRangeFilterConstant
  Inclusive: Boolean

```

## Properties

`Aggregation`

An aggregation function that specifies how to calculate the value of a numeric field for
a topic, Valid values for this structure are `NO_AGGREGATION`, `SUM`,
`AVERAGE`, `COUNT`, `DISTINCT_COUNT`, `MAX`,
`MEDIAN`, `MIN`, `STDEV`, `STDEVP`,
`VAR`,
and `VARP`.

_Required_: No

_Type_: String

_Allowed values_: `NO_AGGREGATION | SUM | AVERAGE | COUNT | DISTINCT_COUNT | MAX | MEDIAN | MIN | STDEV | STDEVP | VAR | VARP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Constant`

The constant used in a
numeric range filter.

_Required_: No

_Type_: [TopicRangeFilterConstant](aws-properties-quicksight-topic-topicrangefilterconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inclusive`

A Boolean value that indicates whether the endpoints of the numeric range are included in the filter.
If set to true, topics whose numeric field value is equal to the endpoint values will be
included in the filter. If set to false, topics whose numeric field value is equal to the
endpoint values will be excluded from the filter.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicNumericEqualityFilter

TopicRangeFilterConstant

All content copied from https://docs.aws.amazon.com/.
