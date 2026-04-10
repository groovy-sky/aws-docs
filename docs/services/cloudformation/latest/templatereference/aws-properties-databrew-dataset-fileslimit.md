This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset FilesLimit

Represents a limit imposed on number of Amazon S3 files that should be
selected for a dataset from a connected Amazon S3 path.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxFiles" : Integer,
  "Order" : String,
  "OrderedBy" : String
}

```

### YAML

```yaml

  MaxFiles: Integer
  Order: String
  OrderedBy: String

```

## Properties

`MaxFiles`

The number of Amazon S3 files to select.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Order`

A criteria to use for Amazon S3 files sorting before their selection. By
default uses DESCENDING order, i.e. most recent files are selected first.
Anotherpossible value is ASCENDING.

_Required_: No

_Type_: String

_Allowed values_: `ASCENDING | DESCENDING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrderedBy`

A criteria to use for Amazon S3 files sorting before their selection. By
default uses LAST\_MODIFIED\_DATE as a sorting criteria. Currently it's the only allowed
value.

_Required_: No

_Type_: String

_Allowed values_: `LAST_MODIFIED_DATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExcelOptions

FilterExpression

All content copied from https://docs.aws.amazon.com/.
