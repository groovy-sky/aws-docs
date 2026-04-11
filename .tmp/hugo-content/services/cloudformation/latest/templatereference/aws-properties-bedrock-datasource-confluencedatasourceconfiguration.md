This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource ConfluenceDataSourceConfiguration

The configuration information to connect to Confluence as your data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerConfiguration" : ConfluenceCrawlerConfiguration,
  "SourceConfiguration" : ConfluenceSourceConfiguration
}

```

### YAML

```yaml

  CrawlerConfiguration:
    ConfluenceCrawlerConfiguration
  SourceConfiguration:
    ConfluenceSourceConfiguration

```

## Properties

`CrawlerConfiguration`

The configuration of the Confluence content. For example, configuring
specific types of Confluence content.

_Required_: No

_Type_: [ConfluenceCrawlerConfiguration](aws-properties-bedrock-datasource-confluencecrawlerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConfiguration`

The endpoint information to connect to your Confluence data source.

_Required_: Yes

_Type_: [ConfluenceSourceConfiguration](aws-properties-bedrock-datasource-confluencesourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceCrawlerConfiguration

ConfluenceSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
