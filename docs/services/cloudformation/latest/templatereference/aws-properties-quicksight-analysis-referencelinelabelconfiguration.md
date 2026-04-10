This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ReferenceLineLabelConfiguration

The label configuration of a reference line.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabelConfiguration" : ReferenceLineCustomLabelConfiguration,
  "FontColor" : String,
  "FontConfiguration" : FontConfiguration,
  "HorizontalPosition" : String,
  "ValueLabelConfiguration" : ReferenceLineValueLabelConfiguration,
  "VerticalPosition" : String
}

```

### YAML

```yaml

  CustomLabelConfiguration:
    ReferenceLineCustomLabelConfiguration
  FontColor: String
  FontConfiguration:
    FontConfiguration
  HorizontalPosition: String
  ValueLabelConfiguration:
    ReferenceLineValueLabelConfiguration
  VerticalPosition: String

```

## Properties

`CustomLabelConfiguration`

The custom label configuration of the label in a reference line.

_Required_: No

_Type_: [ReferenceLineCustomLabelConfiguration](aws-properties-quicksight-analysis-referencelinecustomlabelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontColor`

The font color configuration of the label in a reference line.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontConfiguration`

The font configuration of the label in a reference line.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-analysis-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HorizontalPosition`

The horizontal position configuration of the label in a reference line. Choose one of
the following options:

- `LEFT`

- `CENTER`

- `RIGHT`

_Required_: No

_Type_: String

_Allowed values_: `LEFT | CENTER | RIGHT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueLabelConfiguration`

The value label configuration of the label in a reference line.

_Required_: No

_Type_: [ReferenceLineValueLabelConfiguration](aws-properties-quicksight-analysis-referencelinevaluelabelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerticalPosition`

The vertical position configuration of the label in a reference line. Choose one of the following options:

- `ABOVE`

- `BELOW`

_Required_: No

_Type_: String

_Allowed values_: `ABOVE | BELOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineDynamicDataConfiguration

ReferenceLineStaticDataConfiguration

All content copied from https://docs.aws.amazon.com/.
