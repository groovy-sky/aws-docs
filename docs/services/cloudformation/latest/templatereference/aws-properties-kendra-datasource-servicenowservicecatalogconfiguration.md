This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ServiceNowServiceCatalogConfiguration

Provides the configuration information for crawling service catalog items in the
ServiceNow site

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlAttachments" : Boolean,
  "DocumentDataFieldName" : String,
  "DocumentTitleFieldName" : String,
  "ExcludeAttachmentFilePatterns" : [ String, ... ],
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "IncludeAttachmentFilePatterns" : [ String, ... ]
}

```

### YAML

```yaml

  CrawlAttachments: Boolean
  DocumentDataFieldName: String
  DocumentTitleFieldName: String
  ExcludeAttachmentFilePatterns:
    - String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  IncludeAttachmentFilePatterns:
    - String

```

## Properties

`CrawlAttachments`

`TRUE` to index attachments to service catalog items.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentDataFieldName`

The name of the ServiceNow field that is mapped to the index document contents field
in the Amazon Kendra index.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleFieldName`

The name of the ServiceNow field that is mapped to the index document title
field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeAttachmentFilePatterns`

A list of regular expression patterns to exclude certain attachments of catalogs in
your ServiceNow. Item that match the patterns are excluded from the index. Items that
don't match the patterns are included in the index. If an item matches both an inclusion
and exclusion pattern, the exclusion pattern takes precedence and the item isn't
included in the index.

The regex is applied to the file name of the attachment.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

Maps attributes or field names of catalogs to Amazon Kendra index field names. To
create custom fields, use the `UpdateIndex` API before you map to ServiceNow
fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
ServiceNow data source field names must exist in your ServiceNow custom metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeAttachmentFilePatterns`

A list of regular expression patterns to include certain attachments of catalogs in
your ServiceNow. Item that match the patterns are included in the index. Items that
don't match the patterns are excluded from the index. If an item matches both an
inclusion and exclusion pattern, the exclusion pattern takes precedence and the item
isn't included in the index.

The regex is applied to the file name of the attachment.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceNowKnowledgeArticleConfiguration

SharePointConfiguration

All content copied from https://docs.aws.amazon.com/.
