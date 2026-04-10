This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet FieldFolder

A FieldFolder element is a folder that contains fields and nested subfolders.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ String, ... ],
  "Description" : String
}

```

### YAML

```yaml

  Columns:
    - String
  Description: String

```

## Properties

`Columns`

A folder has a list of columns. A column can only be in one folder.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for a field folder.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationTableSource

FilterOperation

All content copied from https://docs.aws.amazon.com/.
