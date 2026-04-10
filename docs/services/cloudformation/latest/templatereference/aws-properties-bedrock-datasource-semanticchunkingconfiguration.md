This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource SemanticChunkingConfiguration

Settings for semantic document chunking for a data source. Semantic chunking splits
a document into into smaller documents based on groups of similar content derived from the text
with natural language processing.

With semantic chunking, each sentence is compared to the next to determine how similar they are.
You specify a threshold in the form of a percentile, where adjacent sentences that are less similar than
that percentage of sentence pairs are divided into separate chunks. For example, if you set the threshold to
90, then the 10 percent of sentence pairs that are least similar are split. So if you have 101 sentences,
100 sentence pairs are compared, and the 10 with the least similarity are split, creating 11 chunks. These
chunks are further split if they exceed the max token size.

You must also specify a buffer size, which determines whether sentences are compared in isolation, or
within a moving context window that includes the previous and following sentence. For example, if you set
the buffer size to `1`, the embedding for sentence 10 is derived from sentences 9, 10, and 11
combined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BreakpointPercentileThreshold" : Integer,
  "BufferSize" : Integer,
  "MaxTokens" : Integer
}

```

### YAML

```yaml

  BreakpointPercentileThreshold: Integer
  BufferSize: Integer
  MaxTokens: Integer

```

## Properties

`BreakpointPercentileThreshold`

The dissimilarity threshold for splitting chunks.

_Required_: Yes

_Type_: Integer

_Minimum_: `50`

_Maximum_: `99`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BufferSize`

The buffer size.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxTokens`

The maximum number of tokens that a chunk can contain.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SeedUrl

ServerSideEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
