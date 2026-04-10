This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SalesforceDataSourceConfiguration

The configuration information to connect to Salesforce as your data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerConfiguration" : SalesforceCrawlerConfiguration,
  "SourceConfiguration" : SalesforceSourceConfiguration
}

```

### YAML

```yaml

  CrawlerConfiguration:
    SalesforceCrawlerConfiguration
  SourceConfiguration:
    SalesforceSourceConfiguration

```

## Properties

`CrawlerConfiguration`

The configuration of the Salesforce content. For example, configuring
specific types of Salesforce content.

_Required_: No

_Type_: [SalesforceCrawlerConfiguration](aws-properties-bedrock-datasource-salesforcecrawlerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConfiguration`

The endpoint information to connect to your Salesforce data source.

_Required_: Yes

_Type_: [SalesforceSourceConfiguration](aws-properties-bedrock-datasource-salesforcesourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceCrawlerConfiguration

SalesforceSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
