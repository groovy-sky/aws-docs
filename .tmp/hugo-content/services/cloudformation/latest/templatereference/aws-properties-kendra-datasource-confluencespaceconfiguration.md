This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConfluenceSpaceConfiguration

Configuration information for indexing Confluence spaces.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlArchivedSpaces" : Boolean,
  "CrawlPersonalSpaces" : Boolean,
  "ExcludeSpaces" : [ String, ... ],
  "IncludeSpaces" : [ String, ... ],
  "SpaceFieldMappings" : [ ConfluenceSpaceToIndexFieldMapping, ... ]
}

```

### YAML

```yaml

  CrawlArchivedSpaces: Boolean
  CrawlPersonalSpaces: Boolean
  ExcludeSpaces:
    - String
  IncludeSpaces:
    - String
  SpaceFieldMappings:
    - ConfluenceSpaceToIndexFieldMapping

```

## Properties

`CrawlArchivedSpaces`

`TRUE` to index archived spaces.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlPersonalSpaces`

`TRUE` to index personal spaces. You can add restrictions to items in
personal spaces. If personal spaces are indexed, queries without user context
information may return restricted items from a personal space in their results. For more
information, see [Filtering on user\
context](../../../kendra/latest/dg/user-context-filter.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeSpaces`

A list of space keys of Confluence spaces. If you include a key, the blogs, documents,
and attachments in the space are not indexed. If a space is in both the
`ExcludeSpaces` and the `IncludeSpaces` list, the space is
excluded.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSpaces`

A list of space keys for Confluence spaces. If you include a key, the blogs,
documents, and attachments in the space are indexed. Spaces that aren't in the list
aren't indexed. A space in the list must exist. Otherwise, Amazon Kendra logs an
error when the data source is synchronized. If a space is in both the
`IncludeSpaces` and the `ExcludeSpaces` list, the space is
excluded.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceFieldMappings`

Maps attributes or field names of Confluence spaces to Amazon Kendra index field
names. To create custom fields, use the `UpdateIndex` API before you map to
Confluence fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Confluence data source field names must exist in your Confluence custom metadata.

If you specify the `SpaceFieldMappings` parameter, you must specify at
least one field mapping.

_Required_: No

_Type_: Array of [ConfluenceSpaceToIndexFieldMapping](aws-properties-kendra-datasource-confluencespacetoindexfieldmapping.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluencePageToIndexFieldMapping

ConfluenceSpaceToIndexFieldMapping

All content copied from https://docs.aws.amazon.com/.
