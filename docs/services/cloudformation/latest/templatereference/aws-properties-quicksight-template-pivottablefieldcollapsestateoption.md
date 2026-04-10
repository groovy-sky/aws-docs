This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PivotTableFieldCollapseStateOption

The collapse state options for the pivot table field options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String,
  "Target" : PivotTableFieldCollapseStateTarget
}

```

### YAML

```yaml

  State: String
  Target:
    PivotTableFieldCollapseStateTarget

```

## Properties

`State`

The state of the field target of a pivot table. Choose one of the following options:

- `COLLAPSED`

- `EXPANDED`

_Required_: No

_Type_: String

_Allowed values_: `COLLAPSED | EXPANDED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

A tagged-union object that sets the collapse state.

_Required_: Yes

_Type_: [PivotTableFieldCollapseStateTarget](aws-properties-quicksight-template-pivottablefieldcollapsestatetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableDataPathOption

PivotTableFieldCollapseStateTarget

All content copied from https://docs.aws.amazon.com/.
