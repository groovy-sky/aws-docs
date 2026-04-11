This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase OpenSearchServerlessFieldMapping

Contains the names of the fields to which to map information about the vector store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetadataField" : String,
  "TextField" : String,
  "VectorField" : String
}

```

### YAML

```yaml

  MetadataField: String
  TextField: String
  VectorField: String

```

## Properties

`MetadataField`

The name of the field in which Amazon Bedrock stores metadata about the vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextField`

The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorField`

The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchServerlessConfiguration

PineconeConfiguration

All content copied from https://docs.aws.amazon.com/.
