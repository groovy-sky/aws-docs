This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerSiteMapsConfiguration

Provides the configuration information of the sitemap URLs to crawl.

_When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://aws.amazon.com/aup) and all other_
_Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index_
_your own webpages, or webpages that you have authorization to index._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SiteMaps" : [ String, ... ]
}

```

### YAML

```yaml

  SiteMaps:
    - String

```

## Properties

`SiteMaps`

The list of sitemap URLs of the websites you want to crawl.

The list can include a maximum of three sitemap URLs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerSeedUrlConfiguration

WebCrawlerUrls

All content copied from https://docs.aws.amazon.com/.
