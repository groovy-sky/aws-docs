This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FreeFormLayoutElement

An element within a free-form layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundStyle" : FreeFormLayoutElementBackgroundStyle,
  "BorderRadius" : String,
  "BorderStyle" : FreeFormLayoutElementBorderStyle,
  "ElementId" : String,
  "ElementType" : String,
  "Height" : String,
  "LoadingAnimation" : LoadingAnimation,
  "Padding" : String,
  "RenderingRules" : [ SheetElementRenderingRule, ... ],
  "SelectedBorderStyle" : FreeFormLayoutElementBorderStyle,
  "Visibility" : String,
  "Width" : String,
  "XAxisLocation" : String,
  "YAxisLocation" : String
}

```

### YAML

```yaml

  BackgroundStyle:
    FreeFormLayoutElementBackgroundStyle
  BorderRadius: String
  BorderStyle:
    FreeFormLayoutElementBorderStyle
  ElementId: String
  ElementType: String
  Height: String
  LoadingAnimation:
    LoadingAnimation
  Padding: String
  RenderingRules:
    - SheetElementRenderingRule
  SelectedBorderStyle:
    FreeFormLayoutElementBorderStyle
  Visibility: String
  Width: String
  XAxisLocation: String
  YAxisLocation: String

```

## Properties

`BackgroundStyle`

The background style configuration of a free-form layout element.

_Required_: No

_Type_: [FreeFormLayoutElementBackgroundStyle](aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderRadius`

The border radius of a free-form layout element.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BorderStyle`

The border style configuration of a free-form layout element.

_Required_: No

_Type_: [FreeFormLayoutElementBorderStyle](aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementId`

A unique identifier for an element within a free-form layout.

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

`Height`

The height of an element within a free-form layout.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadingAnimation`

The loading animation configuration of a free-form layout element.

_Required_: No

_Type_: [LoadingAnimation](aws-properties-quicksight-dashboard-loadinganimation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Padding`

The padding of a free-form layout element.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenderingRules`

The rendering rules that determine when an element should be displayed within a free-form layout.

_Required_: No

_Type_: Array of [SheetElementRenderingRule](aws-properties-quicksight-dashboard-sheetelementrenderingrule.md)

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedBorderStyle`

The border style configuration of a free-form layout element. This border style is used when the element is selected.

_Required_: No

_Type_: [FreeFormLayoutElementBorderStyle](aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of an element within a free-form layout.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

The width of an element within a free-form layout.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisLocation`

The x-axis coordinate of the element.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxisLocation`

The y-axis coordinate of the element.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FreeFormLayoutConfiguration

FreeFormLayoutElementBackgroundStyle

All content copied from https://docs.aws.amazon.com/.
