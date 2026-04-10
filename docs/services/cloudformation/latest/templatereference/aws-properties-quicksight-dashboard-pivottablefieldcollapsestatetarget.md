This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PivotTableFieldCollapseStateTarget

The target of a pivot table field collapse state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldDataPathValues" : [ DataPathValue, ... ],
  "FieldId" : String
}

```

### YAML

```yaml

  FieldDataPathValues:
    - DataPathValue
  FieldId: String

```

## Properties

`FieldDataPathValues`

The data path of the pivot table's header. Used to set the collapse state.

_Required_: No

_Type_: Array of [DataPathValue](aws-properties-quicksight-dashboard-datapathvalue.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field ID of the pivot table that the collapse state needs to be set to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableFieldCollapseStateOption

PivotTableFieldOption

All content copied from https://docs.aws.amazon.com/.
