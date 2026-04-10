This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase UrlConfiguration

The configuration of the URL/URLs for the web content that you want to crawl. You should be authorized to crawl
the URLs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SeedUrls" : [ SeedUrl, ... ]
}

```

### YAML

```yaml

  SeedUrls:
    - SeedUrl

```

## Properties

`SeedUrls`

List of URLs for crawling.

_Required_: No

_Type_: Array of [SeedUrl](aws-properties-wisdom-knowledgebase-seedurl.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VectorIngestionConfiguration

All content copied from https://docs.aws.amazon.com/.
