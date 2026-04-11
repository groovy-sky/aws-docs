This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SelectedSheetsFilterScopeConfiguration

The configuration for applying a filter to specific sheets or visuals. You can apply this filter to multiple visuals that are on one sheet or to all visuals on a sheet.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SheetVisualScopingConfigurations" : [ SheetVisualScopingConfiguration, ... ]
}

```

### YAML

```yaml

  SheetVisualScopingConfigurations:
    - SheetVisualScopingConfiguration

```

## Properties

`SheetVisualScopingConfigurations`

The sheet ID and visual IDs of the sheet and visuals that the filter is applied to.

_Required_: No

_Type_: Array of [SheetVisualScopingConfiguration](aws-properties-quicksight-analysis-sheetvisualscopingconfiguration.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SectionStyle

SeriesItem

All content copied from https://docs.aws.amazon.com/.
