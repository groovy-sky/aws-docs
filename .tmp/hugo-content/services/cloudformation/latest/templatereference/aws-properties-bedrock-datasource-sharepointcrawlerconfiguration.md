This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SharePointCrawlerConfiguration

The configuration of the SharePoint content. For example, configuring
specific types of SharePoint content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterConfiguration" : CrawlFilterConfiguration
}

```

### YAML

```yaml

  FilterConfiguration:
    CrawlFilterConfiguration

```

## Properties

`FilterConfiguration`

The configuration of filtering the SharePoint content. For example,
configuring regular expression patterns to include or exclude certain content.

_Required_: No

_Type_: [CrawlFilterConfiguration](aws-properties-bedrock-datasource-crawlfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionConfiguration

SharePointDataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
