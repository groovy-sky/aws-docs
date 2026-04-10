This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis KPISortConfiguration

The sort configuration of a KPI visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TrendGroupSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  TrendGroupSort:
    - FieldSortOptions

```

## Properties

`TrendGroupSort`

The sort configuration of the trend group fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-analysis-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPIProgressBarConditionalFormatting

KPISparklineOptions

All content copied from https://docs.aws.amazon.com/.
