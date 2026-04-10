This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DateTimeFormatConfiguration

Formatting configuration for `DateTime` fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeFormat" : String,
  "NullValueFormatConfiguration" : NullValueFormatConfiguration,
  "NumericFormatConfiguration" : NumericFormatConfiguration
}

```

### YAML

```yaml

  DateTimeFormat: String
  NullValueFormatConfiguration:
    NullValueFormatConfiguration
  NumericFormatConfiguration:
    NumericFormatConfiguration

```

## Properties

`DateTimeFormat`

Determines the `DateTime` format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullValueFormatConfiguration`

The options that determine the null value format configuration.

_Required_: No

_Type_: [NullValueFormatConfiguration](aws-properties-quicksight-analysis-nullvalueformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericFormatConfiguration`

The formatting configuration for numeric `DateTime` fields.

_Required_: No

_Type_: [NumericFormatConfiguration](aws-properties-quicksight-analysis-numericformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateTimeDefaultValues

DateTimeHierarchy

All content copied from https://docs.aws.amazon.com/.
