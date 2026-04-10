This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SameSheetTargetVisualConfiguration

The configuration of the same-sheet target visuals that you want to be filtered.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetVisualOptions" : String,
  "TargetVisuals" : [ String, ... ]
}

```

### YAML

```yaml

  TargetVisualOptions: String
  TargetVisuals:
    - String

```

## Properties

`TargetVisualOptions`

The options that choose the target visual in the same sheet.

Valid values are defined as follows:

- `ALL_VISUALS`: Applies the filter operation to all visuals in the same sheet.

_Required_: No

_Type_: String

_Allowed values_: `ALL_VISUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetVisuals`

A list of the target visual IDs that are located in the same sheet of the analysis.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `512 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RowAlternateColorOptions

SankeyDiagramAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
