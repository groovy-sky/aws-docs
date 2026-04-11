This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FreeFormSectionLayoutConfiguration

The free-form layout configuration of a section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Elements" : [ FreeFormLayoutElement, ... ]
}

```

### YAML

```yaml

  Elements:
    - FreeFormLayoutElement

```

## Properties

`Elements`

The elements that are included in the free-form layout.

_Required_: Yes

_Type_: Array of [FreeFormLayoutElement](aws-properties-quicksight-template-freeformlayoutelement.md)

_Minimum_: `0`

_Maximum_: `430`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FreeFormLayoutScreenCanvasSizeOptions

FunnelChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
