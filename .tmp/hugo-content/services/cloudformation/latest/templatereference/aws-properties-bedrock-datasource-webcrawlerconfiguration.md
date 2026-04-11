This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource WebCrawlerConfiguration

The configuration of web URLs that you want to crawl.
You should be authorized to crawl the URLs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerLimits" : WebCrawlerLimits,
  "ExclusionFilters" : [ String, ... ],
  "InclusionFilters" : [ String, ... ],
  "Scope" : String,
  "UserAgent" : String,
  "UserAgentHeader" : String
}

```

### YAML

```yaml

  CrawlerLimits:
    WebCrawlerLimits
  ExclusionFilters:
    - String
  InclusionFilters:
    - String
  Scope: String
  UserAgent: String
  UserAgentHeader: String

```

## Properties

`CrawlerLimits`

The configuration of crawl limits for the web URLs.

_Required_: No

_Type_: [WebCrawlerLimits](aws-properties-bedrock-datasource-webcrawlerlimits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionFilters`

A list of one or more exclusion regular expression patterns to exclude
certain URLs. If you specify an inclusion and exclusion filter/pattern
and both match a URL, the exclusion filter takes precedence and the web
content of the URL isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000 | 25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionFilters`

A list of one or more inclusion regular expression patterns to include
certain URLs. If you specify an inclusion and exclusion filter/pattern
and both match a URL, the exclusion filter takes precedence and the web
content of the URL isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000 | 25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The scope of what is crawled for your URLs.

You can choose to crawl only web pages that belong to the same host or primary
domain. For example, only web pages that contain the seed URL
"https://docs.aws.amazon.com/bedrock/latest/userguide/" and no other domains.
You can choose to include sub domains in addition to the host or primary domain.
For example, web pages that contain "aws.amazon.com" can also include sub domain
"docs.aws.amazon.com".

_Required_: No

_Type_: String

_Allowed values_: `HOST_ONLY | SUBDOMAINS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAgent`

Returns the user agent suffix for your web crawler.

_Required_: No

_Type_: String

_Minimum_: `15`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAgentHeader`

A string used for identifying the crawler or bot when it accesses a web server. The user agent header value
consists of the `bedrockbot`, UUID, and a user agent suffix for your crawler (if one is provided).
By default, it is set to `bedrockbot_UUID`. You can optionally append a custom
suffix to `bedrockbot_UUID` to allowlist a specific user agent permitted to access your source URLs.

_Required_: No

_Type_: String

_Minimum_: `61`

_Maximum_: `86`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorIngestionConfiguration

WebCrawlerLimits

All content copied from https://docs.aws.amazon.com/.
