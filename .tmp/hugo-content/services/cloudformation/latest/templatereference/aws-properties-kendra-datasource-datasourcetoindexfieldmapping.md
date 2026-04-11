This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DataSourceToIndexFieldMapping

Maps a column or attribute in the data source to an index field. You must first create
the fields in the index using the [UpdateIndex](../../../kendra/latest/dg/api-updateindex.md) operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceFieldName" : String,
  "DateFieldFormat" : String,
  "IndexFieldName" : String
}

```

### YAML

```yaml

  DataSourceFieldName: String
  DateFieldFormat: String
  IndexFieldName: String

```

## Properties

`DataSourceFieldName`

The name of the field in the data source. You must first create the index field
using the `UpdateIndex` API.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateFieldFormat`

The format for date fields in the data source. If the field specified in
`DataSourceFieldName` is a date field, you must specify the date
format. If the field is not a date field, an exception is thrown.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexFieldName`

The name of the index field to map to the data source field. The index field type
must match the data source field type.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceConfiguration

DataSourceVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
