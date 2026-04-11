This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase FixedSizeChunkingConfiguration

Configurations for when you choose fixed-size chunking. If you set the `chunkingStrategy` as
`NONE`, exclude this field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTokens" : Number,
  "OverlapPercentage" : Number
}

```

### YAML

```yaml

  MaxTokens: Number
  OverlapPercentage: Number

```

## Properties

`MaxTokens`

The maximum number of tokens to include in a chunk.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverlapPercentage`

The percentage of overlap between adjacent chunks of a data source.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrawlerLimits

HierarchicalChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
