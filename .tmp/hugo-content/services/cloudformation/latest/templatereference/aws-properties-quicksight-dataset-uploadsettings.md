This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet UploadSettings

Information about the format for a source file or files.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainsHeader" : Boolean,
  "Delimiter" : String,
  "Format" : String,
  "StartFromRow" : Number,
  "TextQualifier" : String
}

```

### YAML

```yaml

  ContainsHeader: Boolean
  Delimiter: String
  Format: String
  StartFromRow: Number
  TextQualifier: String

```

## Properties

`ContainsHeader`

Whether the file has a header row, or the files each have a header row.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Delimiter`

The delimiter between values in the file.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

File format.

_Required_: No

_Type_: String

_Allowed values_: `CSV | TSV | CLF | ELF | XLSX | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartFromRow`

A row number to start reading data from.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextQualifier`

Text qualifier.

_Required_: No

_Type_: String

_Allowed values_: `DOUBLE_QUOTE | SINGLE_QUOTE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagColumnOperation

ValueColumnConfiguration

All content copied from https://docs.aws.amazon.com/.
