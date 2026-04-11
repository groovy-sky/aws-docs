This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GridLayoutElement

An element within a grid layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundStyle" : GridLayoutElementBackgroundStyle,
  "BorderRadius" : String,
  "BorderStyle" : GridLayoutElementBorderStyle,
  "ColumnIndex" : Number,
  "ColumnSpan" : Number,
  "ElementId" : String,
  "ElementType" : String,
  "LoadingAnimation" : LoadingAnimation,
  "Padding" : String,
  "RowIndex" : Number,
  "RowSpan" : Number,
  "SelectedBorderStyle" : GridLayoutElementBorderStyle
}

```

### YAML

```yaml

  BackgroundStyle:
    GridLayoutElementBackgroundStyle
  BorderRadius: String
  BorderStyle:
    GridLayoutElementBorderStyle
  ColumnIndex: Number
  ColumnSpan: Number
  ElementId: String
  ElementType: String
  LoadingAnimation:
    LoadingAnimation
  Padding: String
  RowIndex: Number
  RowSpan: Number
  SelectedBorderStyle:
    GridLayoutElementBorderStyle

```

## Properties

`BackgroundStyle`

The background style configuration of a grid layout element.

_Required_: No

_Type_: [GridLayoutElementBackgroundStyle](aws-properties-quicksight-dashboard-gridlayoutelementbackgroundstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderRadius`

The border radius of a grid layout element.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderStyle`

The border style configuration of a grid layout element.

_Required_: No

_Type_: [GridLayoutElementBorderStyle](aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnIndex`

The column index for the upper left corner of an element.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `35`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnSpan`

The width of a grid element expressed as a number of grid columns.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementId`

A unique identifier for an element within a grid layout.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementType`

The type of element.

_Required_: Yes

_Type_: String

_Allowed values_: `VISUAL | FILTER_CONTROL | PARAMETER_CONTROL | TEXT_BOX | IMAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadingAnimation`

The configuration of loading animation in free-form layout.

_Required_: No

_Type_: [LoadingAnimation](aws-properties-quicksight-dashboard-loadinganimation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Padding`

The padding of a grid layout element.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowIndex`

The row index for the upper left corner of an element.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `9009`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowSpan`

The height of a grid element expressed as a number of grid rows.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `21`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedBorderStyle`

The border style configuration of a grid layout element. This border style is used when the element is selected.

_Required_: No

_Type_: [GridLayoutElementBorderStyle](aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GridLayoutConfiguration

GridLayoutElementBackgroundStyle

All content copied from https://docs.aws.amazon.com/.
