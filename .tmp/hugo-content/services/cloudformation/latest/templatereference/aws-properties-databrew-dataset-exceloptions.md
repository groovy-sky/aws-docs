This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset ExcelOptions

Represents a set of options that define how DataBrew will interpret a Microsoft Excel file when
creating a dataset from that file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderRow" : Boolean,
  "SheetIndexes" : [ Integer, ... ],
  "SheetNames" : [ String, ... ]
}

```

### YAML

```yaml

  HeaderRow: Boolean
  SheetIndexes:
    - Integer
  SheetNames:
    - String

```

## Properties

`HeaderRow`

A variable that specifies whether the first row in the file is parsed as the
header. If this value is false, column names are auto-generated.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetIndexes`

One or more sheet numbers in the Excel file that will be included in the
dataset.

_Required_: No

_Type_: Array of Integer

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetNames`

One or more named sheets in the Excel file that will be included in the dataset.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatetimeOptions

FilesLimit

All content copied from https://docs.aws.amazon.com/.
