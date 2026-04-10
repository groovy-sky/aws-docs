This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource HierarchicalChunkingLevelConfiguration

Token settings for a layer in a hierarchical chunking configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTokens" : Integer
}

```

### YAML

```yaml

  MaxTokens: Integer

```

## Properties

`MaxTokens`

The maximum number of tokens that a chunk can contain in this layer.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HierarchicalChunkingConfiguration

IntermediateStorage

All content copied from https://docs.aws.amazon.com/.
