This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DataSourceConfiguration

Provides the configuration information for an Amazon Kendra data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfluenceConfiguration" : ConfluenceConfiguration,
  "DatabaseConfiguration" : DatabaseConfiguration,
  "GoogleDriveConfiguration" : GoogleDriveConfiguration,
  "OneDriveConfiguration" : OneDriveConfiguration,
  "S3Configuration" : S3DataSourceConfiguration,
  "SalesforceConfiguration" : SalesforceConfiguration,
  "ServiceNowConfiguration" : ServiceNowConfiguration,
  "SharePointConfiguration" : SharePointConfiguration,
  "TemplateConfiguration" : TemplateConfiguration,
  "WebCrawlerConfiguration" : WebCrawlerConfiguration,
  "WorkDocsConfiguration" : WorkDocsConfiguration
}

```

### YAML

```yaml

  ConfluenceConfiguration:
    ConfluenceConfiguration
  DatabaseConfiguration:
    DatabaseConfiguration
  GoogleDriveConfiguration:
    GoogleDriveConfiguration
  OneDriveConfiguration:
    OneDriveConfiguration
  S3Configuration:
    S3DataSourceConfiguration
  SalesforceConfiguration:
    SalesforceConfiguration
  ServiceNowConfiguration:
    ServiceNowConfiguration
  SharePointConfiguration:
    SharePointConfiguration
  TemplateConfiguration:
    TemplateConfiguration
  WebCrawlerConfiguration:
    WebCrawlerConfiguration
  WorkDocsConfiguration:
    WorkDocsConfiguration

```

## Properties

`ConfluenceConfiguration`

Provides the configuration information to connect to Confluence as your data
source.

_Required_: No

_Type_: [ConfluenceConfiguration](aws-properties-kendra-datasource-confluenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseConfiguration`

Provides the configuration information to connect to a database as your data
source.

_Required_: No

_Type_: [DatabaseConfiguration](aws-properties-kendra-datasource-databaseconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleDriveConfiguration`

Provides the configuration information to connect to Google Drive as your data
source.

_Required_: No

_Type_: [GoogleDriveConfiguration](aws-properties-kendra-datasource-googledriveconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneDriveConfiguration`

Provides the configuration information to connect to Microsoft OneDrive as your data
source.

_Required_: No

_Type_: [OneDriveConfiguration](aws-properties-kendra-datasource-onedriveconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Provides the configuration information to connect to an Amazon S3 bucket as your
data source.

###### Note

Amazon Kendra now supports an upgraded Amazon S3 connector.

You must now use the [TemplateConfiguration](../../../../reference/kendra/latest/apireference/api-templateconfiguration.md) object instead of the
`S3DataSourceConfiguration` object to configure your connector.

Connectors configured using the older console and API architecture will continue to
function as configured. However, you won't be able to edit or update them. If you want
to edit or update your connector configuration, you must create a new connector.

We recommended migrating your connector workflow to the upgraded version. Support for
connectors configured using the older architecture is scheduled to end by June 2024.

_Required_: No

_Type_: [S3DataSourceConfiguration](aws-properties-kendra-datasource-s3datasourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SalesforceConfiguration`

Provides the configuration information to connect to Salesforce as your data
source.

_Required_: No

_Type_: [SalesforceConfiguration](aws-properties-kendra-datasource-salesforceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNowConfiguration`

Provides the configuration information to connect to ServiceNow as your data
source.

_Required_: No

_Type_: [ServiceNowConfiguration](aws-properties-kendra-datasource-servicenowconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePointConfiguration`

Provides the configuration information to connect to Microsoft SharePoint as your data
source.

_Required_: No

_Type_: [SharePointConfiguration](aws-properties-kendra-datasource-sharepointconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

Provides a template for the configuration information to connect to your data
source.

_Required_: No

_Type_: [TemplateConfiguration](aws-properties-kendra-datasource-templateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebCrawlerConfiguration`

Provides the configuration information required for Amazon Kendra Web Crawler.

_Required_: No

_Type_: [WebCrawlerConfiguration](aws-properties-kendra-datasource-webcrawlerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkDocsConfiguration`

Provides the configuration information to connect to WorkDocs as your data
source.

_Required_: No

_Type_: [WorkDocsConfiguration](aws-properties-kendra-datasource-workdocsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseConfiguration

DataSourceToIndexFieldMapping

All content copied from https://docs.aws.amazon.com/.
