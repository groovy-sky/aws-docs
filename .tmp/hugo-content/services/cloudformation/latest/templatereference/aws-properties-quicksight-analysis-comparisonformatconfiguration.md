This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ComparisonFormatConfiguration

The format of the comparison.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberDisplayFormatConfiguration" : NumberDisplayFormatConfiguration,
  "PercentageDisplayFormatConfiguration" : PercentageDisplayFormatConfiguration
}

```

### YAML

```yaml

  NumberDisplayFormatConfiguration:
    NumberDisplayFormatConfiguration
  PercentageDisplayFormatConfiguration:
    PercentageDisplayFormatConfiguration

```

## Properties

`NumberDisplayFormatConfiguration`

The number display format.

_Required_: No

_Type_: [NumberDisplayFormatConfiguration](aws-properties-quicksight-analysis-numberdisplayformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentageDisplayFormatConfiguration`

The percentage display format.

_Required_: No

_Type_: [PercentageDisplayFormatConfiguration](aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComparisonConfiguration

Computation

All content copied from https://docs.aws.amazon.com/.
