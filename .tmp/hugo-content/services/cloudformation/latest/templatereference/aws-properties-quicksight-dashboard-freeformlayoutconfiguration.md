This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FreeFormLayoutConfiguration

The configuration of a free-form layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanvasSizeOptions" : FreeFormLayoutCanvasSizeOptions,
  "Elements" : [ FreeFormLayoutElement, ... ]
}

```

### YAML

```yaml

  CanvasSizeOptions:
    FreeFormLayoutCanvasSizeOptions
  Elements:
    - FreeFormLayoutElement

```

## Properties

`CanvasSizeOptions`

Property description not available.

_Required_: No

_Type_: [FreeFormLayoutCanvasSizeOptions](aws-properties-quicksight-dashboard-freeformlayoutcanvassizeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Elements`

The elements that are included in a free-form layout.

_Required_: Yes

_Type_: Array of [FreeFormLayoutElement](aws-properties-quicksight-dashboard-freeformlayoutelement.md)

_Minimum_: `0`

_Maximum_: `430`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FreeFormLayoutCanvasSizeOptions

FreeFormLayoutElement

All content copied from https://docs.aws.amazon.com/.
