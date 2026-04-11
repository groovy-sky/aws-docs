This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase HierarchicalChunkingLevelConfiguration

Token settings for each layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTokens" : Number
}

```

### YAML

```yaml

  MaxTokens: Number

```

## Properties

`MaxTokens`

The maximum number of tokens that a chunk can contain in this layer.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HierarchicalChunkingConfiguration

ManagedSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
