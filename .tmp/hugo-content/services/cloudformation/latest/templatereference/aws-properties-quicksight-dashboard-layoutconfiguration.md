This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard LayoutConfiguration

The configuration that determines what the type of layout will be used on a sheet.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FreeFormLayout" : FreeFormLayoutConfiguration,
  "GridLayout" : GridLayoutConfiguration,
  "SectionBasedLayout" : SectionBasedLayoutConfiguration
}

```

### YAML

```yaml

  FreeFormLayout:
    FreeFormLayoutConfiguration
  GridLayout:
    GridLayoutConfiguration
  SectionBasedLayout:
    SectionBasedLayoutConfiguration

```

## Properties

`FreeFormLayout`

A free-form is optimized for a fixed width and has more control over the exact placement of layout elements.

_Required_: No

_Type_: [FreeFormLayoutConfiguration](aws-properties-quicksight-dashboard-freeformlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GridLayout`

A type of layout that can be used on a sheet. In a grid layout, visuals snap to a grid with standard spacing and alignment. Dashboards are displayed as designed, with options to fit to screen or view at actual size. A grid layout can be configured to behave in one of two ways when the viewport is resized: `FIXED` or `RESPONSIVE`.

_Required_: No

_Type_: [GridLayoutConfiguration](aws-properties-quicksight-dashboard-gridlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SectionBasedLayout`

A section based layout organizes visuals into multiple sections and has customized header, footer and page break.

_Required_: No

_Type_: [SectionBasedLayoutConfiguration](aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Layout

LegendOptions

All content copied from https://docs.aws.amazon.com/.
