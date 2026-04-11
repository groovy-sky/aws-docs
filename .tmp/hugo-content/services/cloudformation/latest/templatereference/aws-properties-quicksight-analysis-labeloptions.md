This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LabelOptions

The share label options for the labels.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "FontConfiguration" : FontConfiguration,
  "Visibility" : String
}

```

### YAML

```yaml

  CustomLabel: String
  FontConfiguration:
    FontConfiguration
  Visibility: String

```

## Properties

`CustomLabel`

The text for the label.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontConfiguration`

The font configuration of the label.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-analysis-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Determines whether or not the label is visible.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPIVisualStandardLayout

LayerCustomAction

All content copied from https://docs.aws.amazon.com/.
