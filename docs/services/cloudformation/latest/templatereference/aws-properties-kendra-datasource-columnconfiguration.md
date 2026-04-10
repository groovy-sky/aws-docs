This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ColumnConfiguration

Provides information about how Amazon Kendra should use the columns of a database
in an index.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChangeDetectingColumns" : [ String, ... ],
  "DocumentDataColumnName" : String,
  "DocumentIdColumnName" : String,
  "DocumentTitleColumnName" : String,
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  ChangeDetectingColumns:
    - String
  DocumentDataColumnName: String
  DocumentIdColumnName: String
  DocumentTitleColumnName: String
  FieldMappings:
    - DataSourceToIndexFieldMapping

```

## Properties

`ChangeDetectingColumns`

One to five columns that indicate when a document in the database has changed.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentDataColumnName`

The column that contains the contents of the document.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentIdColumnName`

The column that provides the document's identifier.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleColumnName`

The column that contains the title of the document.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

An array of objects that map database column names to the corresponding fields in an
index. You must first create the fields in the index using the [UpdateIndex](../../../kendra/latest/dg/api-updateindex.md)
operation.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AclConfiguration

ConfluenceAttachmentConfiguration

All content copied from https://docs.aws.amazon.com/.
