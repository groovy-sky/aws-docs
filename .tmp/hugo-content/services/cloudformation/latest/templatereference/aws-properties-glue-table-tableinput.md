This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table TableInput

A structure used to define a table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "Owner" : String,
  "Parameters" : Json,
  "PartitionKeys" : [ Column, ... ],
  "Retention" : Integer,
  "StorageDescriptor" : StorageDescriptor,
  "TableType" : String,
  "TargetTable" : TableIdentifier,
  "ViewExpandedText" : String,
  "ViewOriginalText" : String
}

```

### YAML

```yaml

  Description: String
  Name: String
  Owner: String
  Parameters: Json
  PartitionKeys:
    - Column
  Retention: Integer
  StorageDescriptor:
    StorageDescriptor
  TableType: String
  TargetTable:
    TableIdentifier
  ViewExpandedText: String
  ViewOriginalText: String

```

## Properties

`Description`

A description of the table.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The table name. For Hive compatibility, this is folded to
lowercase when it is stored.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

The table owner. Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

These key-value pairs define properties associated with the table.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionKeys`

A list of columns by which the table is partitioned. Only primitive
types are supported as partition keys.

When you create a table used by Amazon Athena, and you do not specify any
`partitionKeys`, you must at least set the value of `partitionKeys` to
an empty list. For example:

`"PartitionKeys": []`

_Required_: No

_Type_: Array of [Column](aws-properties-glue-table-column.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Retention`

The retention time for this table.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageDescriptor`

A storage descriptor containing information about the physical storage
of this table.

_Required_: No

_Type_: [StorageDescriptor](aws-properties-glue-table-storagedescriptor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableType`

The type of this table.
AWS Glue will create tables with the `EXTERNAL_TABLE` type.
Other services, such as Athena, may create tables with additional table types.

AWS Glue related table types:

EXTERNAL\_TABLE

Hive compatible attribute - indicates a non-Hive managed table.

GOVERNED

Used by AWS Lake Formation.
The AWS Glue Data Catalog understands `GOVERNED`.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTable`

A `TableIdentifier` structure that describes a target table for resource linking.

_Required_: No

_Type_: [TableIdentifier](aws-properties-glue-table-tableidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewExpandedText`

Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.

_Required_: No

_Type_: String

_Maximum_: `409600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewOriginalText`

Included for Apache Hive compatibility. Not used in the normal course of AWS Glue operations.
If the table is a `VIRTUAL_VIEW`, certain Athena configuration encoded in base64.

_Required_: No

_Type_: String

_Maximum_: `409600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableIdentifier

AWS::Glue::TableOptimizer

All content copied from https://docs.aws.amazon.com/.
