This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis NumericFormatConfiguration

The options that determine the numeric format configuration.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CurrencyDisplayFormatConfiguration" : CurrencyDisplayFormatConfiguration,
  "NumberDisplayFormatConfiguration" : NumberDisplayFormatConfiguration,
  "PercentageDisplayFormatConfiguration" : PercentageDisplayFormatConfiguration
}

```

### YAML

```yaml

  CurrencyDisplayFormatConfiguration:
    CurrencyDisplayFormatConfiguration
  NumberDisplayFormatConfiguration:
    NumberDisplayFormatConfiguration
  PercentageDisplayFormatConfiguration:
    PercentageDisplayFormatConfiguration

```

## Properties

`CurrencyDisplayFormatConfiguration`

The options that determine the currency display format configuration.

_Required_: No

_Type_: [CurrencyDisplayFormatConfiguration](aws-properties-quicksight-analysis-currencydisplayformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberDisplayFormatConfiguration`

The options that determine the number display format configuration.

_Required_: No

_Type_: [NumberDisplayFormatConfiguration](aws-properties-quicksight-analysis-numberdisplayformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentageDisplayFormatConfiguration`

The options that determine the percentage display format configuration.

_Required_: No

_Type_: [PercentageDisplayFormatConfiguration](aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericEqualityFilter

NumericRangeFilter

All content copied from https://docs.aws.amazon.com/.
