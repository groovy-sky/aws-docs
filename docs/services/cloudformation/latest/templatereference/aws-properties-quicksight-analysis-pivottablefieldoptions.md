This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotTableFieldOptions

The field options for a pivot table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollapseStateOptions" : [ PivotTableFieldCollapseStateOption, ... ],
  "DataPathOptions" : [ PivotTableDataPathOption, ... ],
  "SelectedFieldOptions" : [ PivotTableFieldOption, ... ]
}

```

### YAML

```yaml

  CollapseStateOptions:
    - PivotTableFieldCollapseStateOption
  DataPathOptions:
    - PivotTableDataPathOption
  SelectedFieldOptions:
    - PivotTableFieldOption

```

## Properties

`CollapseStateOptions`

The collapse state options for the pivot table field options.

_Required_: No

_Type_: Array of [PivotTableFieldCollapseStateOption](aws-properties-quicksight-analysis-pivottablefieldcollapsestateoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPathOptions`

The data path options for the pivot table field options.

_Required_: No

_Type_: Array of [PivotTableDataPathOption](aws-properties-quicksight-analysis-pivottabledatapathoption.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedFieldOptions`

The selected field options for the pivot table field options.

_Required_: No

_Type_: Array of [PivotTableFieldOption](aws-properties-quicksight-analysis-pivottablefieldoption.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableFieldOption

PivotTableFieldSubtotalOptions

All content copied from https://docs.aws.amazon.com/.
