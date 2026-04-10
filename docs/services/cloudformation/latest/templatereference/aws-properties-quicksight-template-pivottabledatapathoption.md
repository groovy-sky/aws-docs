This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PivotTableDataPathOption

The data path options for the pivot table field options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPathList" : [ DataPathValue, ... ],
  "Width" : String
}

```

### YAML

```yaml

  DataPathList:
    - DataPathValue
  Width: String

```

## Properties

`DataPathList`

The list of data path values for the data path options.

_Required_: Yes

_Type_: Array of [DataPathValue](aws-properties-quicksight-template-datapathvalue.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

The width of the data path option.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableConfiguration

PivotTableFieldCollapseStateOption

All content copied from https://docs.aws.amazon.com/.
