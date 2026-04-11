This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SharePointConfiguration

Provides the configuration information to connect to Microsoft SharePoint as your data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlAttachments" : Boolean,
  "DisableLocalGroups" : Boolean,
  "DocumentTitleFieldName" : String,
  "ExclusionPatterns" : [ String, ... ],
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "InclusionPatterns" : [ String, ... ],
  "SecretArn" : String,
  "SharePointVersion" : String,
  "SslCertificateS3Path" : S3Path,
  "Urls" : [ String, ... ],
  "UseChangeLog" : Boolean,
  "VpcConfiguration" : DataSourceVpcConfiguration
}

```

### YAML

```yaml

  CrawlAttachments: Boolean
  DisableLocalGroups: Boolean
  DocumentTitleFieldName: String
  ExclusionPatterns:
    - String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  InclusionPatterns:
    - String
  SecretArn: String
  SharePointVersion: String
  SslCertificateS3Path:
    S3Path
  Urls:
    - String
  UseChangeLog: Boolean
  VpcConfiguration:
    DataSourceVpcConfiguration

```

## Properties

`CrawlAttachments`

`TRUE` to index document attachments.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableLocalGroups`

`TRUE` to disable local groups information.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleFieldName`

The Microsoft SharePoint attribute field that contains the title of the
document.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionPatterns`

A list of regular expression patterns. Documents that match the patterns are excluded
from the index. Documents that don't match the patterns are included in the index. If a
document matches both an exclusion pattern and an inclusion pattern, the document is not
included in the index.

The regex is applied to the display URL of the SharePoint document.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

A list of `DataSourceToIndexFieldMapping` objects that map Microsoft
SharePoint attributes or fields to Amazon Kendra index fields. You must first create the
index fields using the [UpdateIndex](../../../kendra/latest/dg/api-updateindex.md) operation before you
map SharePoint attributes. For more information, see [Mapping Data Source Fields](../../../kendra/latest/dg/field-mapping.md).

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPatterns`

A list of regular expression patterns to include certain documents in your SharePoint.
Documents that match the patterns are included in the index. Documents that don't match
the patterns are excluded from the index. If a document matches both an inclusion and
exclusion pattern, the exclusion pattern takes precedence and the document isn't
included in the index.

The regex applies to the display URL of the SharePoint document.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the
user name and password required to connect to the SharePoint instance. For more
information, see [Microsoft\
SharePoint](../../../kendra/latest/dg/data-source-sharepoint.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePointVersion`

The version of Microsoft SharePoint that you use.

_Required_: Yes

_Type_: String

_Allowed values_: `SHAREPOINT_ONLINE | SHAREPOINT_2013 | SHAREPOINT_2016`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslCertificateS3Path`

Information required to find a specific file in an Amazon S3 bucket.

_Required_: No

_Type_: [S3Path](aws-properties-kendra-datasource-s3path.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Urls`

The Microsoft SharePoint site URLs for the documents you want to index.

_Required_: Yes

_Type_: Array of String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseChangeLog`

`TRUE` to use the SharePoint change log to determine which documents
require updating in the index. Depending on the change log's size, it may take longer
for Amazon Kendra to use the change log than to scan all of your documents in
SharePoint.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

Provides information for connecting to an Amazon VPC.

_Required_: No

_Type_: [DataSourceVpcConfiguration](aws-properties-kendra-datasource-datasourcevpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceNowServiceCatalogConfiguration

SqlConfiguration

All content copied from https://docs.aws.amazon.com/.
