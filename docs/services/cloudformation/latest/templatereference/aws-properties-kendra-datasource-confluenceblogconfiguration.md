This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluenceBlogConfiguration

Configuration of blog settings for the Confluence data source. Blogs are always
indexed unless filtered from the index by the `ExclusionPatterns` or
`InclusionPatterns` fields in the `ConfluenceConfiguration`
object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlogFieldMappings" : [ ConfluenceBlogToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  BlogFieldMappings:
    - ConfluenceBlogToIndexFieldMapping

```

## Properties

`BlogFieldMappings`

Maps attributes or field names of Confluence blogs to Amazon Kendra index field
names. To create custom fields, use the `UpdateIndex` API before you map to
Confluence fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Confluence data source field names must exist in your Confluence custom metadata.

If you specify the `BlogFieldMappings` parameter, you must specify at least
one field mapping.

_Required_: No

_Type_: Array of [ConfluenceBlogToIndexFieldMapping](aws-properties-kendra-datasource-confluenceblogtoindexfieldmapping.md)

_Minimum_: `1`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceAttachmentToIndexFieldMapping

ConfluenceBlogToIndexFieldMapping

All content copied from https://docs.aws.amazon.com/.
