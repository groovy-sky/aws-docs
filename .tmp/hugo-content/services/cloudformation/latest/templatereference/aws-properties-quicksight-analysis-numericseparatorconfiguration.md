This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis NumericSeparatorConfiguration

The options that determine the numeric separator configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DecimalSeparator" : String,
  "ThousandsSeparator" : ThousandSeparatorOptions
}

```

### YAML

```yaml

  DecimalSeparator: String
  ThousandsSeparator:
    ThousandSeparatorOptions

```

## Properties

`DecimalSeparator`

Determines the decimal separator.

_Required_: No

_Type_: String

_Allowed values_: `COMMA | DOT | SPACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThousandsSeparator`

The options that determine the thousands separator configuration.

_Required_: No

_Type_: [ThousandSeparatorOptions](aws-properties-quicksight-analysis-thousandseparatoroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericRangeFilterValue

PaginationConfiguration

All content copied from https://docs.aws.amazon.com/.
