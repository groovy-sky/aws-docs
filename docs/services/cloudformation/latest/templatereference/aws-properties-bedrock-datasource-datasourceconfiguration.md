This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource DataSourceConfiguration

The connection configuration for the data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfluenceConfiguration" : ConfluenceDataSourceConfiguration,
  "S3Configuration" : S3DataSourceConfiguration,
  "SalesforceConfiguration" : SalesforceDataSourceConfiguration,
  "SharePointConfiguration" : SharePointDataSourceConfiguration,
  "Type" : String,
  "WebConfiguration" : WebDataSourceConfiguration
}

```

### YAML

```yaml

  ConfluenceConfiguration:
    ConfluenceDataSourceConfiguration
  S3Configuration:
    S3DataSourceConfiguration
  SalesforceConfiguration:
    SalesforceDataSourceConfiguration
  SharePointConfiguration:
    SharePointDataSourceConfiguration
  Type: String
  WebConfiguration:
    WebDataSourceConfiguration

```

## Properties

`ConfluenceConfiguration`

The configuration information to connect to Confluence as your data source.

###### Note

Confluence data source connector is in preview release and is subject to change.

_Required_: No

_Type_: [ConfluenceDataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-confluencedatasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

The configuration information to connect to Amazon S3 as your data source.

_Required_: No

_Type_: [S3DataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-s3datasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SalesforceConfiguration`

The configuration information to connect to Salesforce as your data source.

###### Note

Salesforce data source connector is in preview release and is subject to change.

_Required_: No

_Type_: [SalesforceDataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePointConfiguration`

The configuration information to connect to SharePoint as your data source.

###### Note

SharePoint data source connector is in preview release and is subject to change.

_Required_: No

_Type_: [SharePointDataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of data source.

_Required_: Yes

_Type_: String

_Allowed values_: `S3 | CONFLUENCE | SALESFORCE | SHAREPOINT | WEB | CUSTOM | REDSHIFT_METADATA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WebConfiguration`

The configuration of web URLs to crawl for your data source.
You should be authorized to crawl the URLs.

###### Note

Crawling web URLs as your data source is in preview release
and is subject to change.

_Required_: No

_Type_: [WebDataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-webdatasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomTransformationConfiguration

EnrichmentStrategyConfiguration
