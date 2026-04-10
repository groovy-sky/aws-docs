This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource GoogleDriveConfiguration

Provides the configuration information to connect to Google Drive as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeMimeTypes" : [ String, ... ],
  "ExcludeSharedDrives" : [ String, ... ],
  "ExcludeUserAccounts" : [ String, ... ],
  "ExclusionPatterns" : [ String, ... ],
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "InclusionPatterns" : [ String, ... ],
  "SecretArn" : String
}

```

### YAML

```yaml

  ExcludeMimeTypes:
    - String
  ExcludeSharedDrives:
    - String
  ExcludeUserAccounts:
    - String
  ExclusionPatterns:
    - String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  InclusionPatterns:
    - String
  SecretArn: String

```

## Properties

`ExcludeMimeTypes`

A list of MIME types to exclude from the index. All documents matching the specified
MIME type are excluded.

For a list of MIME types, see [Using a\
Google Workspace Drive data source](../../../kendra/latest/dg/data-source-google-drive.md).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeSharedDrives`

A list of identifiers or shared drives to exclude from the index. All files and
folders stored on the shared drive are excluded.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeUserAccounts`

A list of email addresses of the users. Documents owned by these users are excluded
from the index. Documents shared with excluded users are indexed unless they are
excluded in another way.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of regular expression patterns to exclude certain items in your Google Drive,
including shared drives and users' My Drives. Items that match the patterns are excluded
from the index. Items that don't match the patterns are included in the index. If an
item matches both an inclusion and exclusion pattern, the exclusion pattern takes
precedence and the item isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

Maps Google Drive data source attributes or field names to Amazon Kendra index
field names. To create custom fields, use the `UpdateIndex` API before you
map to Google Drive fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Google Drive data source field names must exist in your Google Drive custom
metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of regular expression patterns to include certain items in your Google Drive,
including shared drives and users' My Drives. Items that match the patterns are included
in the index. Items that don't match the patterns are excluded from the index. If an
item matches both an inclusion and exclusion pattern, the exclusion pattern takes
precedence and the item isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of a AWS Secrets Managersecret that contains the
credentials required to connect to Google Drive. For more information, see [Using a\
Google Workspace Drive data source](../../../kendra/latest/dg/data-source-google-drive.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentsMetadataConfiguration

HookConfiguration

All content copied from https://docs.aws.amazon.com/.
