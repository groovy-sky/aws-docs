This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SharePointDataSourceConfiguration

The configuration information to connect to SharePoint as your data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerConfiguration" : SharePointCrawlerConfiguration,
  "SourceConfiguration" : SharePointSourceConfiguration
}

```

### YAML

```yaml

  CrawlerConfiguration:
    SharePointCrawlerConfiguration
  SourceConfiguration:
    SharePointSourceConfiguration

```

## Properties

`CrawlerConfiguration`

The configuration of the SharePoint content. For example, configuring
specific types of SharePoint content.

_Required_: No

_Type_: [SharePointCrawlerConfiguration](aws-properties-bedrock-datasource-sharepointcrawlerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConfiguration`

The endpoint information to connect to your SharePoint data source.

_Required_: Yes

_Type_: [SharePointSourceConfiguration](aws-properties-bedrock-datasource-sharepointsourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SharePointCrawlerConfiguration

SharePointSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
