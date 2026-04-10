This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FilterScopeConfiguration

The scope configuration for a `FilterGroup`.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllSheets" : Json,
  "SelectedSheets" : SelectedSheetsFilterScopeConfiguration
}

```

### YAML

```yaml

  AllSheets: Json
  SelectedSheets:
    SelectedSheetsFilterScopeConfiguration

```

## Properties

`AllSheets`

The configuration that applies a filter to all sheets. When you choose `AllSheets` as the value for a `FilterScopeConfiguration`, this filter is applied to all visuals of all sheets in an Analysis, Dashboard, or Template. The `AllSheetsFilterScopeConfiguration` is chosen.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedSheets`

The configuration for applying a filter to specific sheets.

_Required_: No

_Type_: [SelectedSheetsFilterScopeConfiguration](aws-properties-quicksight-analysis-selectedsheetsfilterscopeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterRelativeDateTimeControl

FilterSelectableValues

All content copied from https://docs.aws.amazon.com/.
