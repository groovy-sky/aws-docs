This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluencePageConfiguration

Configuration of the page settings for the Confluence data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PageFieldMappings" : [ ConfluencePageToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  PageFieldMappings:
    - ConfluencePageToIndexFieldMapping

```

## Properties

`PageFieldMappings`

Maps attributes or field names of Confluence pages to Amazon Kendra index field
names. To create custom fields, use the `UpdateIndex` API before you map to
Confluence fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Confluence data source field names must exist in your Confluence custom metadata.

If you specify the `PageFieldMappings` parameter, you must specify at least
one field mapping.

_Required_: No

_Type_: Array of [ConfluencePageToIndexFieldMapping](aws-properties-kendra-datasource-confluencepagetoindexfieldmapping.md)

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceConfiguration

ConfluencePageToIndexFieldMapping

All content copied from https://docs.aws.amazon.com/.
