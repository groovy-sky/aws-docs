This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceStandardObjectAttachmentConfiguration

Provides the configuration information for processing attachments to Salesforce
standard objects.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentTitleFieldName" : String,
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  DocumentTitleFieldName: String
  FieldMappings:
    - DataSourceToIndexFieldMapping

```

## Properties

`DocumentTitleFieldName`

The name of the field used for the document title.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

One or more objects that map fields in attachments to Amazon Kendra index
fields.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceStandardKnowledgeArticleTypeConfiguration

SalesforceStandardObjectConfiguration

All content copied from https://docs.aws.amazon.com/.
