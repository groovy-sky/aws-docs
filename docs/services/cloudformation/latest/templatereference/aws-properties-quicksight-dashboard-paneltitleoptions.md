This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PanelTitleOptions

The options that determine the title styles for each small multiples
panel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FontConfiguration" : FontConfiguration,
  "HorizontalTextAlignment" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  FontConfiguration:
    FontConfiguration
  HorizontalTextAlignment: String
  Visibility: String

```

## Properties

`FontConfiguration`

Property description not available.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-dashboard-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HorizontalTextAlignment`

Sets the horizontal text alignment of the title within each panel.

_Required_: No

_Type_: String

_Allowed values_: `LEFT | CENTER | RIGHT | AUTO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Determines whether or not panel titles are displayed.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PanelConfiguration

ParameterControl

All content copied from https://docs.aws.amazon.com/.
