This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RdsFieldMapping

Contains the names of the fields to which to map information about the vector store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomMetadataField" : String,
  "MetadataField" : String,
  "PrimaryKeyField" : String,
  "TextField" : String,
  "VectorField" : String
}

```

### YAML

```yaml

  CustomMetadataField: String
  MetadataField: String
  PrimaryKeyField: String
  TextField: String
  VectorField: String

```

## Properties

`CustomMetadataField`

Provide a name for the universal metadata field where Amazon Bedrock will store any custom metadata from
your data source.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataField`

The name of the field in which Amazon Bedrock stores metadata about the vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrimaryKeyField`

The name of the field in which Amazon Bedrock stores the ID for each entry.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextField`

The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorField`

The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsConfiguration

RedshiftConfiguration

All content copied from https://docs.aws.amazon.com/.
