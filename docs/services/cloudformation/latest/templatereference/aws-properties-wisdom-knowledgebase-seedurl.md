This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase SeedUrl

A URL for crawling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Url" : String
}

```

### YAML

```yaml

  Url: String

```

## Properties

`Url`

URL for crawling

_Required_: No

_Type_: String

_Pattern_: `^https?://[A-Za-z0-9][^\s]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RenderingConfiguration

SemanticChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
