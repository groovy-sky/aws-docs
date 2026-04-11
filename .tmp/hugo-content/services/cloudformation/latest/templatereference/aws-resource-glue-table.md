This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table

The `AWS::Glue::Table` resource specifies tabular data in the AWS Glue data
catalog. For more information, see [Defining Tables in the AWS Glue Data\
Catalog](../../../glue/latest/dg/tables-described.md) and [Table Structure](../../../glue/latest/dg/aws-glue-api-catalog-tables.md#aws-glue-api-catalog-tables-Table) in the _AWS Glue Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Table",
  "Properties" : {
      "CatalogId" : String,
      "DatabaseName" : String,
      "OpenTableFormatInput" : OpenTableFormatInput,
      "TableInput" : TableInput
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Table
Properties:
  CatalogId: String
  DatabaseName: String
  OpenTableFormatInput:
    OpenTableFormatInput
  TableInput:
    TableInput

```

## Properties

`CatalogId`

The ID of the Data Catalog in which to create the `Table`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the database where the table metadata resides.
For Hive compatibility, this must be all lowercase.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OpenTableFormatInput`

Specifies an `OpenTableFormatInput` structure when creating an open format table.

_Required_: No

_Type_: [OpenTableFormatInput](aws-properties-glue-table-opentableformatinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableInput`

A structure used to define a table.

_Required_: Yes

_Type_: [TableInput](aws-properties-glue-table-tableinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the table name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Encryption

Column

All content copied from https://docs.aws.amazon.com/.
