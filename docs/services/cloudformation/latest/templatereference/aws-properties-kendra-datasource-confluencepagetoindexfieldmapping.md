This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluencePageToIndexFieldMapping

Maps attributes or field names of Confluence pages to Amazon Kendra index field
names. To create custom fields, use the `UpdateIndex` API before you map to
Confluence fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Confluence data source field names must exist in your Confluence custom metadata.

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

The name of the field in the data source.

_Required_: Yes

_Type_: String

_Allowed values_: `AUTHOR | CONTENT_STATUS | CREATED_DATE | DISPLAY_URL | ITEM_TYPE | LABELS | MODIFIED_DATE | PARENT_ID | SPACE_KEY | SPACE_NAME | URL | VERSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateFieldFormat`

The format for date fields in the data source. If the field specified in
`DataSourceFieldName` is a date field you must specify the date format.
If the field is not a date field, an exception is thrown.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexFieldName`

The name of the index field to map to the Confluence data source field. The index
field type must match the Confluence field type.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluencePageConfiguration

ConfluenceSpaceConfiguration

All content copied from https://docs.aws.amazon.com/.
