This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerConfiguration

Provides the configuration information required for Amazon Kendra Web Crawler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationConfiguration" : WebCrawlerAuthenticationConfiguration,
  "CrawlDepth" : Integer,
  "MaxContentSizePerPageInMegaBytes" : Number,
  "MaxLinksPerPage" : Integer,
  "MaxUrlsPerMinuteCrawlRate" : Integer,
  "ProxyConfiguration" : ProxyConfiguration,
  "UrlExclusionPatterns" : [ String, ... ],
  "UrlInclusionPatterns" : [ String, ... ],
  "Urls" : WebCrawlerUrls
}

```

### YAML

```yaml

  AuthenticationConfiguration:
    WebCrawlerAuthenticationConfiguration
  CrawlDepth: Integer
  MaxContentSizePerPageInMegaBytes: Number
  MaxLinksPerPage: Integer
  MaxUrlsPerMinuteCrawlRate: Integer
  ProxyConfiguration:
    ProxyConfiguration
  UrlExclusionPatterns:
    - String
  UrlInclusionPatterns:
    - String
  Urls:
    WebCrawlerUrls

```

## Properties

`AuthenticationConfiguration`

Configuration information required to connect to websites using authentication.

You can connect to websites using basic authentication of user name and password. You
use a secret in [AWS Secrets Manager](../../../secretsmanager/latest/userguide/intro.md) to
store your authentication credentials.

You must provide the website host name and port number. For example, the host name of
https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard
port for HTTPS.

_Required_: No

_Type_: [WebCrawlerAuthenticationConfiguration](aws-properties-kendra-datasource-webcrawlerauthenticationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlDepth`

The 'depth' or number of levels from the seed level to crawl. For example, the seed
URL page is depth 1 and any hyperlinks on this page that are also crawled are depth 2.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxContentSizePerPageInMegaBytes`

The maximum size (in MB) of a web page or attachment to crawl.

Files larger than this size (in MB) are skipped/not crawled.

The default maximum size of a web page or attachment is set to 50 MB.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLinksPerPage`

The maximum number of URLs on a web page to include when crawling a website. This
number is per web page.

As a website’s web pages are crawled, any URLs the web pages link to are also crawled.
URLs on a web page are crawled in order of appearance.

The default maximum links per page is 100.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUrlsPerMinuteCrawlRate`

The maximum number of URLs crawled per website host per minute.

A minimum of one URL is required.

The default maximum number of URLs crawled per website host per minute is 300.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProxyConfiguration`

Configuration information required to connect to your internal websites via a web
proxy.

You must provide the website host name and port number. For example, the host name of
https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard
port for HTTPS.

Web proxy credentials are optional and you can use them to connect to a web proxy
server that requires basic authentication. To store web proxy credentials, you use a
secret in [AWS Secrets Manager](../../../secretsmanager/latest/userguide/intro.md).

_Required_: No

_Type_: [ProxyConfiguration](aws-properties-kendra-datasource-proxyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlExclusionPatterns`

A list of regular expression patterns to exclude certain URLs to crawl. URLs that
match the patterns are excluded from the index. URLs that don't match the patterns are
included in the index. If a URL matches both an inclusion and exclusion pattern, the
exclusion pattern takes precedence and the URL file isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlInclusionPatterns`

A list of regular expression patterns to include certain URLs to crawl. URLs that
match the patterns are included in the index. URLs that don't match the patterns are
excluded from the index. If a URL matches both an inclusion and exclusion pattern, the
exclusion pattern takes precedence and the URL file isn't included in the index.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Urls`

Specifies the seed or starting point URLs of the websites or the sitemap URLs of the
websites you want to crawl.

You can include website subdomains. You can list up to 100 seed URLs and up to three
sitemap URLs.

You can only crawl websites that use the secure communication protocol, Hypertext
Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it
could be that the website is blocked from crawling.

_When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://aws.amazon.com/aup) and all other_
_Amazon terms. Remember that you must only use Amazon Kendra Web Crawler to index_
_your own webpages, or webpages that you have authorization to index._

_Required_: Yes

_Type_: [WebCrawlerUrls](aws-properties-kendra-datasource-webcrawlerurls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebCrawlerBasicAuthentication

WebCrawlerSeedUrlConfiguration

All content copied from https://docs.aws.amazon.com/.
