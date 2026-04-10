This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ComparisonConfiguration

The comparison display configuration of a KPI or gauge chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonFormat" : ComparisonFormatConfiguration,
  "ComparisonMethod" : String
}

```

### YAML

```yaml

  ComparisonFormat:
    ComparisonFormatConfiguration
  ComparisonMethod: String

```

## Properties

`ComparisonFormat`

The format of the comparison.

_Required_: No

_Type_: [ComparisonFormatConfiguration](aws-properties-quicksight-dashboard-comparisonformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparisonMethod`

The method of the comparison. Choose from the following options:

- `DIFFERENCE`

- `PERCENT_DIFFERENCE`

- `PERCENT`

_Required_: No

_Type_: String

_Allowed values_: `DIFFERENCE | PERCENT_DIFFERENCE | PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComboChartVisual

ComparisonFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
