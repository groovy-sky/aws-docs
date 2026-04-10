This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PercentageDisplayFormatConfiguration

The options that determine the percentage display format configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DecimalPlacesConfiguration" : DecimalPlacesConfiguration,
  "NegativeValueConfiguration" : NegativeValueConfiguration,
  "NullValueFormatConfiguration" : NullValueFormatConfiguration,
  "Prefix" : String,
  "SeparatorConfiguration" : NumericSeparatorConfiguration,
  "Suffix" : String
}

```

### YAML

```yaml

  DecimalPlacesConfiguration:
    DecimalPlacesConfiguration
  NegativeValueConfiguration:
    NegativeValueConfiguration
  NullValueFormatConfiguration:
    NullValueFormatConfiguration
  Prefix: String
  SeparatorConfiguration:
    NumericSeparatorConfiguration
  Suffix: String

```

## Properties

`DecimalPlacesConfiguration`

The option that determines the decimal places configuration.

_Required_: No

_Type_: [DecimalPlacesConfiguration](aws-properties-quicksight-dashboard-decimalplacesconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NegativeValueConfiguration`

The options that determine the negative value configuration.

_Required_: No

_Type_: [NegativeValueConfiguration](aws-properties-quicksight-dashboard-negativevalueconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullValueFormatConfiguration`

The options that determine the null value format configuration.

_Required_: No

_Type_: [NullValueFormatConfiguration](aws-properties-quicksight-dashboard-nullvalueformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Determines the prefix value of the percentage format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeparatorConfiguration`

The options that determine the numeric separator configuration.

_Required_: No

_Type_: [NumericSeparatorConfiguration](aws-properties-quicksight-dashboard-numericseparatorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Suffix`

Determines the suffix value of the percentage format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterTextFieldControl

PercentileAggregation

All content copied from https://docs.aws.amazon.com/.
