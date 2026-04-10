This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase CrawlerLimits

The limits of the crawler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RateLimit" : Number
}

```

### YAML

```yaml

  RateLimit: Number

```

## Properties

`RateLimit`

The limit rate at which the crawler is configured.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `3000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChunkingConfiguration

FixedSizeChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
