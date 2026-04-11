This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase SemanticChunkingConfiguration

Settings for semantic document chunking for a data source. Semantic chunking splits a document into smaller
documents based on groups of similar content derived from the text with natural language processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BreakpointPercentileThreshold" : Number,
  "BufferSize" : Number,
  "MaxTokens" : Number
}

```

### YAML

```yaml

  BreakpointPercentileThreshold: Number
  BufferSize: Number
  MaxTokens: Number

```

## Properties

`BreakpointPercentileThreshold`

The dissimilarity threshold for splitting chunks.

_Required_: Yes

_Type_: Number

_Minimum_: `50`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferSize`

The buffer size.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTokens`

The maximum number of tokens that a chunk can contain.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SeedUrl

ServerSideEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
