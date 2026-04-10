This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic CellValueSynonym

A structure that represents the cell value synonym.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CellValue" : String,
  "Synonyms" : [ String, ... ]
}

```

### YAML

```yaml

  CellValue: String
  Synonyms:
    - String

```

## Properties

`CellValue`

The cell value.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Synonyms`

Other names or aliases for the cell value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QuickSight::Topic

CollectiveConstant

All content copied from https://docs.aws.amazon.com/.
