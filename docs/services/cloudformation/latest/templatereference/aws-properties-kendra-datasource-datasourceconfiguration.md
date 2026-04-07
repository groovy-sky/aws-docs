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

_Type_: [ConfluenceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-confluenceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseConfiguration`

Provides the configuration information to connect to a database as your data
source.

_Required_: No

_Type_: [DatabaseConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-databaseconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleDriveConfiguration`

Provides the configuration information to connect to Google Drive as your data
source.

_Required_: No

_Type_: [GoogleDriveConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-googledriveconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneDriveConfiguration`

Provides the configuration information to connect to Microsoft OneDrive as your data
source.

_Required_: No

_Type_: [OneDriveConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-onedriveconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Provides the configuration information to connect to an Amazon S3 bucket as your
data source.

###### Note

Amazon Kendra now supports an upgraded Amazon S3 connector.

You must now use the [TemplateConfiguration](https://docs.aws.amazon.com/kendra/latest/APIReference/API_TemplateConfiguration.html) object instead of the
`S3DataSourceConfiguration` object to configure your connector.

Connectors configured using the older console and API architecture will continue to
function as configured. However, you won't be able to edit or update them. If you want
to edit or update your connector configuration, you must create a new connector.

We recommended migrating your connector workflow to the upgraded version. Support for
connectors configured using the older architecture is scheduled to end by June 2024.

_Required_: No

_Type_: [S3DataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-s3datasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SalesforceConfiguration`

Provides the configuration information to connect to Salesforce as your data
source.

_Required_: No

_Type_: [SalesforceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-salesforceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNowConfiguration`

Provides the configuration information to connect to ServiceNow as your data
source.

_Required_: No

_Type_: [ServiceNowConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-servicenowconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePointConfiguration`

Provides the configuration information to connect to Microsoft SharePoint as your data
source.

_Required_: No

_Type_: [SharePointConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-sharepointconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

Provides a template for the configuration information to connect to your data
source.

_Required_: No

_Type_: [TemplateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-templateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebCrawlerConfiguration`

Provides the configuration information required for Amazon Kendra Web Crawler.

_Required_: No

_Type_: [WebCrawlerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-webcrawlerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkDocsConfiguration`

Provides the configuration information to connect to WorkDocs as your data
source.

_Required_: No

_Type_: [WorkDocsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-datasource-workdocsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatabaseConfiguration

DataSourceToIndexFieldMapping
