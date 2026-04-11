This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceStandardKnowledgeArticleTypeConfiguration

Provides the configuration information for standard Salesforce knowledge
articles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentDataFieldName" : String,
  "DocumentTitleFieldName" : String,
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  DocumentDataFieldName: String
  DocumentTitleFieldName: String
  FieldMappings:
    - DataSourceToIndexFieldMapping

```

## Properties

`DocumentDataFieldName`

The name of the field that contains the document data to index.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleFieldName`

The name of the field that contains the document title.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

Maps attributes or field names of the knowledge article to Amazon Kendra index
field names. To create custom fields, use the `UpdateIndex` API before you
map to Salesforce fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Salesforce data source field names must exist in your Salesforce custom metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceKnowledgeArticleConfiguration

SalesforceStandardObjectAttachmentConfiguration

All content copied from https://docs.aws.amazon.com/.
