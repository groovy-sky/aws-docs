This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template WordCloudAggregatedFieldWells

The aggregated field wells of a word cloud.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupBy" : [ DimensionField, ... ],
  "Size" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  GroupBy:
    - DimensionField
  Size:
    - MeasureField

```

## Properties

`GroupBy`

The group by field well of a word cloud. Values are grouped by group by fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-template-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The size field well of a word cloud. Values are aggregated based on group by fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WhatIfRangeScenario

WordCloudChartConfiguration

All content copied from https://docs.aws.amazon.com/.
