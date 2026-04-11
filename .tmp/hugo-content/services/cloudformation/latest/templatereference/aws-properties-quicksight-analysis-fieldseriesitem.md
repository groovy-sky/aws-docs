This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FieldSeriesItem

The field series item configuration of a line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisBinding" : String,
  "FieldId" : String,
  "Settings" : LineChartSeriesSettings
}

```

### YAML

```yaml

  AxisBinding: String
  FieldId: String
  Settings:
    LineChartSeriesSettings

```

## Properties

`AxisBinding`

The axis that you are binding the field to.

_Required_: Yes

_Type_: String

_Allowed values_: `PRIMARY_YAXIS | SECONDARY_YAXIS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field ID of the field for which you are setting the axis binding.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

The options that determine the presentation of line series associated to the field.

_Required_: No

_Type_: [LineChartSeriesSettings](aws-properties-quicksight-analysis-linechartseriessettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldLabelType

FieldSort

All content copied from https://docs.aws.amazon.com/.
