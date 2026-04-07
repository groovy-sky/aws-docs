This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table StorageDescriptor

Describes the physical storage of table data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketColumns" : [ String, ... ],
  "Columns" : [ Column, ... ],
  "Compressed" : Boolean,
  "InputFormat" : String,
  "Location" : String,
  "NumberOfBuckets" : Integer,
  "OutputFormat" : String,
  "Parameters" : Json,
  "SchemaReference" : SchemaReference,
  "SerdeInfo" : SerdeInfo,
  "SkewedInfo" : SkewedInfo,
  "SortColumns" : [ Order, ... ],
  "StoredAsSubDirectories" : Boolean
}

```

### YAML

```yaml

  BucketColumns:
    - String
  Columns:
    - Column
  Compressed: Boolean
  InputFormat: String
  Location: String
  NumberOfBuckets: Integer
  OutputFormat: String
  Parameters: Json
  SchemaReference:
    SchemaReference
  SerdeInfo:
    SerdeInfo
  SkewedInfo:
    SkewedInfo
  SortColumns:
    - Order
  StoredAsSubDirectories: Boolean

```

## Properties

`BucketColumns`

A list of reducer grouping columns, clustering columns, and
bucketing columns in the table.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Columns`

A list of the `Columns` in the table.

_Required_: No

_Type_: Array of [Column](aws-properties-glue-table-column.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compressed`

`True` if the data in the table is compressed, or `False` if
not.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputFormat`

The input format: `SequenceFileInputFormat` (binary),
or `TextInputFormat`, or a custom format.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

The physical location of the table. By default, this takes the form of the warehouse
location, followed by the database location in the warehouse, followed by the table
name.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Maximum_: `2056`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfBuckets`

Must be specified if the table contains any dimension columns.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFormat`

The output format: `SequenceFileOutputFormat` (binary),
or `IgnoreKeyTextOutputFormat`, or a custom format.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The user-supplied properties in key-value form.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaReference`

An object that references a schema stored in the AWS Glue Schema Registry.

_Required_: No

_Type_: [SchemaReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-table-schemareference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SerdeInfo`

The serialization/deserialization (SerDe) information.

_Required_: No

_Type_: [SerdeInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-table-serdeinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkewedInfo`

The information about values that appear frequently in a column (skewed values).

_Required_: No

_Type_: [SkewedInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-table-skewedinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortColumns`

A list specifying the sort order of each bucket in the table.

_Required_: No

_Type_: Array of [Order](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-table-order.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StoredAsSubDirectories`

`True` if the table data is stored in subdirectories, or `False` if
not.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [StorageDescriptor Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) in the _AWS Glue Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SkewedInfo

TableIdentifier
