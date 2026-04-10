This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase WebCrawlerConfiguration

The configuration details for the web data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerLimits" : CrawlerLimits,
  "ExclusionFilters" : [ String, ... ],
  "InclusionFilters" : [ String, ... ],
  "Scope" : String,
  "UrlConfiguration" : UrlConfiguration
}

```

### YAML

```yaml

  CrawlerLimits:
    CrawlerLimits
  ExclusionFilters:
    - String
  InclusionFilters:
    - String
  Scope: String
  UrlConfiguration:
    UrlConfiguration

```

## Properties

`CrawlerLimits`

The configuration of crawl limits for the web URLs.

_Required_: No

_Type_: [CrawlerLimits](aws-properties-wisdom-knowledgebase-crawlerlimits.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExclusionFilters`

A list of one or more exclusion regular expression patterns to exclude certain URLs. If you specify an inclusion
and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the
URL isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InclusionFilters`

A list of one or more inclusion regular expression patterns to include certain URLs. If you specify an inclusion
and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the
URL isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scope`

The scope of what is crawled for your URLs. You can choose to crawl only web pages that belong to the same host
or primary domain. For example, only web pages that contain the seed URL
`https://docs.aws.amazon.com/bedrock/latest/userguide/` and no other domains. You can choose to include
sub domains in addition to the host or primary domain. For example, web pages that contain
`aws.amazon.com` can also include sub domain `docs.aws.amazon.com`.

_Required_: No

_Type_: String

_Allowed values_: `HOST_ONLY | SUBDOMAINS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UrlConfiguration`

The configuration of the URL/URLs for the web content that you want to crawl. You should be authorized to crawl
the URLs.

_Required_: Yes

_Type_: [UrlConfiguration](aws-properties-wisdom-knowledgebase-urlconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorIngestionConfiguration

AWS::Wisdom::MessageTemplate

All content copied from https://docs.aws.amazon.com/.
