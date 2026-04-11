This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource WebDataSourceConfiguration

The configuration details for the web data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerConfiguration" : WebCrawlerConfiguration,
  "SourceConfiguration" : WebSourceConfiguration
}

```

### YAML

```yaml

  CrawlerConfiguration:
    WebCrawlerConfiguration
  SourceConfiguration:
    WebSourceConfiguration

```

## Properties

`CrawlerConfiguration`

The Web Crawler configuration details for the web data source.

_Required_: No

_Type_: [WebCrawlerConfiguration](aws-properties-bedrock-datasource-webcrawlerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConfiguration`

The source configuration details for the web data source.

_Required_: Yes

_Type_: [WebSourceConfiguration](aws-properties-bedrock-datasource-websourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerLimits

WebSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
