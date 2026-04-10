This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluenceAttachmentConfiguration

Configuration of attachment settings for the Confluence data source. Attachment
settings are optional, if you don't specify settings attachments, Amazon Kendra
won't index them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentFieldMappings" : [ ConfluenceAttachmentToIndexFieldMapping, ... ],
  "CrawlAttachments" : Boolean
}

```

### YAML

```yaml

  AttachmentFieldMappings:
    - ConfluenceAttachmentToIndexFieldMapping
  CrawlAttachments: Boolean

```

## Properties

`AttachmentFieldMappings`

Maps attributes or field names of Confluence attachments to Amazon Kendra index
field names. To create custom fields, use the `UpdateIndex` API before you
map to Confluence fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Confluence data source field names must exist in your Confluence custom metadata.

If you specify the `AttachentFieldMappings` parameter, you must specify at
least one field mapping.

_Required_: No

_Type_: Array of [ConfluenceAttachmentToIndexFieldMapping](aws-properties-kendra-datasource-confluenceattachmenttoindexfieldmapping.md)

_Minimum_: `1`

_Maximum_: `11`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlAttachments`

`TRUE` to index attachments of pages and blogs in Confluence.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnConfiguration

ConfluenceAttachmentToIndexFieldMapping

All content copied from https://docs.aws.amazon.com/.
