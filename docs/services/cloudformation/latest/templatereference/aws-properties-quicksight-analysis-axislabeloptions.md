This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AxisLabelOptions

The label options for a chart axis. You must specify the field that the label is targeted to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplyTo" : AxisLabelReferenceOptions,
  "CustomLabel" : String,
  "FontConfiguration" : FontConfiguration
}

```

### YAML

```yaml

  ApplyTo:
    AxisLabelReferenceOptions
  CustomLabel: String
  FontConfiguration:
    FontConfiguration

```

## Properties

`ApplyTo`

The options that indicate which field the label belongs to.

_Required_: No

_Type_: [AxisLabelReferenceOptions](aws-properties-quicksight-analysis-axislabelreferenceoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomLabel`

The text for the axis label.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontConfiguration`

The font configuration of the axis label.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-analysis-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisDisplayRange

AxisLabelReferenceOptions

All content copied from https://docs.aws.amazon.com/.
