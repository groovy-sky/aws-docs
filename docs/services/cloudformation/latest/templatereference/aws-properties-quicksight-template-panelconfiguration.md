This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PanelConfiguration

A collection of options that configure how each panel displays in a small multiples chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "BackgroundVisibility" : String,
  "BorderColor" : String,
  "BorderStyle" : String,
  "BorderThickness" : String,
  "BorderVisibility" : String,
  "GutterSpacing" : String,
  "GutterVisibility" : String,
  "Title" : PanelTitleOptions
}

```

### YAML

```yaml

  BackgroundColor: String
  BackgroundVisibility: String
  BorderColor: String
  BorderStyle: String
  BorderThickness: String
  BorderVisibility: String
  GutterSpacing: String
  GutterVisibility: String
  Title:
    PanelTitleOptions

```

## Properties

`BackgroundColor`

Sets the background color for each panel.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackgroundVisibility`

Determines whether or not a background for each small multiples panel is rendered.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderColor`

Sets the line color of panel borders.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderStyle`

Sets the line style of panel borders.

_Required_: No

_Type_: String

_Allowed values_: `SOLID | DASHED | DOTTED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderThickness`

Sets the line thickness of panel borders.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderVisibility`

Determines whether or not each panel displays a border.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GutterSpacing`

Sets the total amount of negative space to display between sibling panels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GutterVisibility`

Determines whether or not negative space between sibling panels is rendered.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

Configures the title display within each small multiples panel.

_Required_: No

_Type_: [PanelTitleOptions](aws-properties-quicksight-template-paneltitleoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PaginationConfiguration

PanelTitleOptions

All content copied from https://docs.aws.amazon.com/.
