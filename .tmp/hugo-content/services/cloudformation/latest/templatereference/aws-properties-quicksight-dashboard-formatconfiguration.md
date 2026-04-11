This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FormatConfiguration

The formatting configuration for all types of field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeFormatConfiguration" : DateTimeFormatConfiguration,
  "NumberFormatConfiguration" : NumberFormatConfiguration,
  "StringFormatConfiguration" : StringFormatConfiguration
}

```

### YAML

```yaml

  DateTimeFormatConfiguration:
    DateTimeFormatConfiguration
  NumberFormatConfiguration:
    NumberFormatConfiguration
  StringFormatConfiguration:
    StringFormatConfiguration

```

## Properties

`DateTimeFormatConfiguration`

Formatting configuration for `DateTime` fields.

_Required_: No

_Type_: [DateTimeFormatConfiguration](aws-properties-quicksight-dashboard-datetimeformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberFormatConfiguration`

Formatting configuration for number fields.

_Required_: No

_Type_: [NumberFormatConfiguration](aws-properties-quicksight-dashboard-numberformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringFormatConfiguration`

Formatting configuration for string fields.

_Required_: No

_Type_: [StringFormatConfiguration](aws-properties-quicksight-dashboard-stringformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForecastScenario

FreeFormLayoutCanvasSizeOptions

All content copied from https://docs.aws.amazon.com/.
