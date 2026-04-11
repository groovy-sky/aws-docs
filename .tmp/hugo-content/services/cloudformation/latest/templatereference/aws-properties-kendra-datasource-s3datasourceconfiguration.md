This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource S3DataSourceConfiguration

Provides the configuration information to connect to an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessControlListConfiguration" : AccessControlListConfiguration,
  "BucketName" : String,
  "DocumentsMetadataConfiguration" : DocumentsMetadataConfiguration,
  "ExclusionPatterns" : [ String, ... ],
  "InclusionPatterns" : [ String, ... ],
  "InclusionPrefixes" : [ String, ... ]
}

```

### YAML

```yaml

  AccessControlListConfiguration:
    AccessControlListConfiguration
  BucketName: String
  DocumentsMetadataConfiguration:
    DocumentsMetadataConfiguration
  ExclusionPatterns:
    - String
  InclusionPatterns:
    - String
  InclusionPrefixes:
    - String

```

## Properties

`AccessControlListConfiguration`

Provides the path to the S3 bucket that contains the user context filtering files for
the data source. For the format of the file, see [Access control for S3 data\
sources](../../../kendra/latest/dg/s3-acl.md).

_Required_: No

_Type_: [AccessControlListConfiguration](aws-properties-kendra-datasource-accesscontrollistconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

The name of the bucket that contains the documents.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentsMetadataConfiguration`

Specifies document metadata files that contain information such as the document access
control information, source URI, document author, and custom attributes. Each metadata
file contains metadata about a single document.

_Required_: No

_Type_: [DocumentsMetadataConfiguration](aws-properties-kendra-datasource-documentsmetadataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of glob patterns (patterns that can expand a wildcard pattern into a list of
path names that match the given pattern) for certain file names and file types to exclude
from your index. If a document matches both an inclusion and exclusion prefix or pattern,
the exclusion prefix takes precendence and the document is not indexed. Examples of glob
patterns include:

- _/myapp/config/\*_—All files inside config directory.

- _\*\*/\*.png_—All .png files in all directories.

- _\*\*/\*.{png, ico, md}_—All .png, .ico or .md files in all
directories.

- _/myapp/src/\*\*/\*.ts_—All .ts files inside src directory (and all
its subdirectories).

- _\*\*/!(\*.module).ts_—All .ts files but not .module.ts

- _\*.png , \*.jpg_—All PNG and JPEG image files
in a directory (files with the extensions .png and .jpg).

- _\*internal\*_—All files in a directory that
contain 'internal' in the file name, such as 'internal', 'internal\_only',
'company\_internal'.

- _\*\*/\*internal\*_—All internal-related files in
a directory and its subdirectories.

For more examples, see [Use of Exclude and\
Include Filters](../../../cli/latest/reference/s3.md#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of glob patterns (patterns that can expand a wildcard pattern into a list of
path names that match the given pattern) for certain file names and file types to include
in your index. If a document matches both an inclusion and exclusion prefix or pattern,
the exclusion prefix takes precendence and the document is not indexed. Examples of glob
patterns include:

- _/myapp/config/\*_—All files inside config directory.

- _\*\*/\*.png_—All .png files in all directories.

- _\*\*/\*.{png, ico, md}_—All .png, .ico or .md files in all
directories.

- _/myapp/src/\*\*/\*.ts_—All .ts files inside src directory (and all
its subdirectories).

- _\*\*/!(\*.module).ts_—All .ts files but not .module.ts

- _\*.png , \*.jpg_—All PNG and JPEG image files
in a directory (files with the extensions .png and .jpg).

- _\*internal\*_—All files in a directory that
contain 'internal' in the file name, such as 'internal', 'internal\_only',
'company\_internal'.

- _\*\*/\*internal\*_—All internal-related files in
a directory and its subdirectories.

For more examples, see [Use of Exclude and\
Include Filters](../../../cli/latest/reference/s3.md#use-of-exclude-and-include-filters) in the AWS CLI Command Reference.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPrefixes`

A list of S3 prefixes for the documents that should be included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProxyConfiguration

S3Path

All content copied from https://docs.aws.amazon.com/.
