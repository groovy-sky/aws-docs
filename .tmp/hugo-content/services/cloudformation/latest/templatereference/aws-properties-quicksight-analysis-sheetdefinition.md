This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SheetDefinition

A sheet is an object that contains a set of visuals that
are viewed together on one page in a paginated report. Every analysis and dashboard must contain at least one sheet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : String,
  "Description" : String,
  "FilterControls" : [ FilterControl, ... ],
  "Images" : [ SheetImage, ... ],
  "Layouts" : [ Layout, ... ],
  "Name" : String,
  "ParameterControls" : [ ParameterControl, ... ],
  "SheetControlLayouts" : [ SheetControlLayout, ... ],
  "SheetId" : String,
  "TextBoxes" : [ SheetTextBox, ... ],
  "Title" : String,
  "Visuals" : [ Visual, ... ]
}

```

### YAML

```yaml

  ContentType: String
  Description: String
  FilterControls:
    - FilterControl
  Images:
    - SheetImage
  Layouts:
    - Layout
  Name: String
  ParameterControls:
    - ParameterControl
  SheetControlLayouts:
    - SheetControlLayout
  SheetId: String
  TextBoxes:
    - SheetTextBox
  Title: String
  Visuals:
    - Visual

```

## Properties

`ContentType`

The layout content type of the sheet. Choose one of the following options:

- `PAGINATED`: Creates a sheet for a paginated report.

- `INTERACTIVE`: Creates a sheet for an interactive dashboard.

_Required_: No

_Type_: String

_Allowed values_: `PAGINATED | INTERACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the sheet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControls`

The list of filter controls that are on a sheet.

For more information, see [Adding filter controls to analysis sheets](../../../quicksight/latest/user/filter-controls.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [FilterControl](aws-properties-quicksight-analysis-filtercontrol.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Images`

A list of images on a sheet.

_Required_: No

_Type_: Array of [SheetImage](aws-properties-quicksight-analysis-sheetimage.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Layouts`

Layouts define how the components of a sheet are arranged.

For more information, see [Types of layout](../../../quicksight/latest/user/types-of-layout.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [Layout](aws-properties-quicksight-analysis-layout.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the sheet. This name is displayed on the sheet's tab in the Quick
console.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterControls`

The list of parameter controls that are on a sheet.

For more information, see [Using a Control with a Parameter in Amazon Quick Sight](../../../quicksight/latest/user/parameters-controls.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [ParameterControl](aws-properties-quicksight-analysis-parametercontrol.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetControlLayouts`

The control layouts of the sheet.

_Required_: No

_Type_: Array of [SheetControlLayout](aws-properties-quicksight-analysis-sheetcontrollayout.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetId`

The unique identifier of a sheet.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextBoxes`

The text boxes that are on a sheet.

_Required_: No

_Type_: Array of [SheetTextBox](aws-properties-quicksight-analysis-sheettextbox.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the sheet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visuals`

A list of the visuals that are on a sheet. Visual placement is determined by the layout of the sheet.

_Required_: No

_Type_: Array of [Visual](aws-properties-quicksight-analysis-visual.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetControlLayoutConfiguration

SheetElementConfigurationOverrides

All content copied from https://docs.aws.amazon.com/.
