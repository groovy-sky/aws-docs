This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template HistogramAggregatedFieldWells

The field well configuration of a histogram.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Values:
    - MeasureField

```

## Properties

`Values`

The value field wells of a histogram. Values are aggregated by `COUNT` or `DISTINCT_COUNT`.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeatMapVisual

HistogramBinOptions

All content copied from https://docs.aws.amazon.com/.
