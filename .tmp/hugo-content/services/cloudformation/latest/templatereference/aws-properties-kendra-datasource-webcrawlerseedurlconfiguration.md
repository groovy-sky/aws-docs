This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerSeedUrlConfiguration

Provides the configuration information of the seed or starting point URLs to
crawl.

_When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://aws.amazon.com/aup) and all other_
_Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index_
_your own webpages, or webpages that you have authorization to index._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SeedUrls" : [ String, ... ],
  "WebCrawlerMode" : String
}

```

### YAML

```yaml

  SeedUrls:
    - String
  WebCrawlerMode: String

```

## Properties

`SeedUrls`

The list of seed or starting point URLs of the websites you want to crawl.

The list can include a maximum of 100 seed URLs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebCrawlerMode`

You can choose one of the following modes:

- `HOST_ONLY`—crawl only the website host names. For
example, if the seed URL is "abc.example.com", then only URLs with host name
"abc.example.com" are crawled.

- `SUBDOMAINS`—crawl the website host names with subdomains.
For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and
"b.abc.example.com" are also crawled.

- `EVERYTHING`—crawl the website host names with subdomains
and other domains that the web pages link to.

The default mode is set to `HOST_ONLY`.

_Required_: No

_Type_: String

_Allowed values_: `HOST_ONLY | SUBDOMAINS | EVERYTHING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerConfiguration

WebCrawlerSiteMapsConfiguration

All content copied from https://docs.aws.amazon.com/.
