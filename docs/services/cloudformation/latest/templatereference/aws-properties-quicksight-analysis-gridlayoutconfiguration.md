This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GridLayoutConfiguration

The configuration for a grid layout. Also called a tiled layout.

Visuals snap to a grid with standard spacing and alignment. Dashboards are displayed as designed, with options to fit to screen or view at actual size.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanvasSizeOptions" : GridLayoutCanvasSizeOptions,
  "Elements" : [ GridLayoutElement, ... ]
}

```

### YAML

```yaml

  CanvasSizeOptions:
    GridLayoutCanvasSizeOptions
  Elements:
    - GridLayoutElement

```

## Properties

`CanvasSizeOptions`

Property description not available.

_Required_: No

_Type_: [GridLayoutCanvasSizeOptions](aws-properties-quicksight-analysis-gridlayoutcanvassizeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Elements`

The elements that are included in a grid layout.

_Required_: Yes

_Type_: Array of [GridLayoutElement](aws-properties-quicksight-analysis-gridlayoutelement.md)

_Minimum_: `0`

_Maximum_: `430`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GridLayoutCanvasSizeOptions

GridLayoutElement

All content copied from https://docs.aws.amazon.com/.
