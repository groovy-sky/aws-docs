This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerUrls

Specifies the seed or starting point URLs of the websites or the sitemap URLs of the
websites you want to crawl.

You can include website subdomains. You can list up to 100 seed URLs and up to three
sitemap URLs.

You can only crawl websites that use the secure communication protocol, Hypertext
Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it
could be that the website is blocked from crawling.

_When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://aws.amazon.com/aup) and all other_
_Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index_
_your own webpages, or webpages that you have authorization to index._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SeedUrlConfiguration" : WebCrawlerSeedUrlConfiguration,
  "SiteMapsConfiguration" : WebCrawlerSiteMapsConfiguration
}

```

### YAML

```yaml

  SeedUrlConfiguration:
    WebCrawlerSeedUrlConfiguration
  SiteMapsConfiguration:
    WebCrawlerSiteMapsConfiguration

```

## Properties

`SeedUrlConfiguration`

Configuration of the seed or starting point URLs of the websites you want to
crawl.

You can choose to crawl only the website host names, or the website host names with
subdomains, or the website host names with subdomains and other domains that the
web pages link to.

You can list up to 100 seed URLs.

_Required_: No

_Type_: [WebCrawlerSeedUrlConfiguration](aws-properties-kendra-datasource-webcrawlerseedurlconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SiteMapsConfiguration`

Configuration of the sitemap URLs of the websites you want to crawl.

Only URLs belonging to the same website host names are crawled. You can list up to
three sitemap URLs.

_Required_: No

_Type_: [WebCrawlerSiteMapsConfiguration](aws-properties-kendra-datasource-webcrawlersitemapsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerSiteMapsConfiguration

WorkDocsConfiguration

All content copied from https://docs.aws.amazon.com/.
