This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PercentileAggregation

An aggregation based on the percentile of values in a dimension or measure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PercentileValue" : Number
}

```

### YAML

```yaml

  PercentileValue: Number

```

## Properties

`PercentileValue`

The percentile value. This value can be any numeric constant 0–100. A percentile value of 50 computes the median value of the measure.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PercentageDisplayFormatConfiguration

PercentVisibleRange

All content copied from https://docs.aws.amazon.com/.
