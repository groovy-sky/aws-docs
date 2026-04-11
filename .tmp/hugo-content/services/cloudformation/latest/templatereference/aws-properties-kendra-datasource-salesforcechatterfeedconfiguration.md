This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceChatterFeedConfiguration

The configuration information for syncing a Salesforce chatter feed. The contents of
the object comes from the Salesforce FeedItem table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentDataFieldName" : String,
  "DocumentTitleFieldName" : String,
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "IncludeFilterTypes" : [ String, ... ]
}

```

### YAML

```yaml

  DocumentDataFieldName: String
  DocumentTitleFieldName: String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  IncludeFilterTypes:
    - String

```

## Properties

`DocumentDataFieldName`

The name of the column in the Salesforce FeedItem table that contains the content to
index. Typically this is the `Body` column.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleFieldName`

The name of the column in the Salesforce FeedItem table that contains the title of the
document. This is typically the `Title` column.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

Maps fields from a Salesforce chatter feed into Amazon Kendra index
fields.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeFilterTypes`

Filters the documents in the feed based on status of the user. When you specify
`ACTIVE_USERS` only documents from users who have an active account are
indexed. When you specify `STANDARD_USER` only documents for Salesforce
standard users are documented. You can specify both.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Path

SalesforceConfiguration

All content copied from https://docs.aws.amazon.com/.
