This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource WebCrawlerLimits

The rate limits for the URLs that you want to crawl.
You should be authorized to crawl the URLs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxPages" : Integer,
  "RateLimit" : Integer
}

```

### YAML

```yaml

  MaxPages: Integer
  RateLimit: Integer

```

## Properties

`MaxPages`

The max number of web pages crawled from your source URLs, up to 25,000 pages. If
the web pages exceed this limit, the data source sync will fail and no web pages will be ingested.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RateLimit`

The max rate at which pages are crawled, up to 300 per minute per host.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerConfiguration

WebDataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
