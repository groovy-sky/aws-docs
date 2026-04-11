This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WorkDocsConfiguration

Provides the configuration information to connect to WorkDocs
as your data source.

WorkDocs connector is available in Oregon, North Virginia, Sydney, Singapore and Ireland
regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlComments" : Boolean,
  "ExclusionPatterns" : [ String, ... ],
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "InclusionPatterns" : [ String, ... ],
  "OrganizationId" : String,
  "UseChangeLog" : Boolean
}

```

### YAML

```yaml

  CrawlComments: Boolean
  ExclusionPatterns:
    - String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  InclusionPatterns:
    - String
  OrganizationId: String
  UseChangeLog: Boolean

```

## Properties

`CrawlComments`

`TRUE` to include comments on documents
in your index. Including comments in your index means each comment
is a document that can be searched on.

The default is set to `FALSE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of regular expression patterns to exclude certain files
in your WorkDocs site repository. Files that match the patterns
are excluded from the index. Files that don’t match the patterns
are included in the index. If a file matches both an inclusion and exclusion
pattern, the exclusion pattern takes precedence and the file isn't included
in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

A list of `DataSourceToIndexFieldMapping` objects that
map WorkDocs data source attributes or field names to Amazon Kendra
index field names. To create custom fields, use the
`UpdateIndex` API before you map to WorkDocs fields.
For more information, see [Mapping \
data source fields](../../../kendra/latest/dg/field-mapping.md). The WorkDocs data source field names
must exist in your WorkDocs custom metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of regular expression patterns to include certain files
in your WorkDocs site repository. Files that match the patterns
are included in the index. Files that don't match the patterns are
excluded from the index. If a file matches both an inclusion and exclusion
pattern, the exclusion pattern takes precedence and the file isn't included
in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationId`

The identifier of the directory corresponding to your
WorkDocs site repository.

You can find the organization ID in the
[Directory Service](https://console.aws.amazon.com/directoryservicev2) by going to
**Active Directory**, then
**Directories**. Your WorkDocs site directory has an
ID, which is the organization ID. You can also set up a new WorkDocs
directory in the Directory Service console and enable a WorkDocs site
for the directory in the WorkDocs console.

_Required_: Yes

_Type_: String

_Pattern_: `d-[0-9a-fA-F]{10}`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseChangeLog`

`TRUE` to use the WorkDocs change log to determine
which documents require updating in the index. Depending on the change log's
size, it may take longer for Amazon Kendra to use the change log than to
scan all of your documents in WorkDocs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerUrls

AWS::Kendra::Faq

All content copied from https://docs.aws.amazon.com/.
