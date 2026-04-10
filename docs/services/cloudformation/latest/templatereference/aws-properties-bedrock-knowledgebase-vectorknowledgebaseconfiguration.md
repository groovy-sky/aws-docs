This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase VectorKnowledgeBaseConfiguration

Contains details about the model used to create vector embeddings for the knowledge base.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmbeddingModelArn" : String,
  "EmbeddingModelConfiguration" : EmbeddingModelConfiguration,
  "SupplementalDataStorageConfiguration" : SupplementalDataStorageConfiguration
}

```

### YAML

```yaml

  EmbeddingModelArn: String
  EmbeddingModelConfiguration:
    EmbeddingModelConfiguration
  SupplementalDataStorageConfiguration:
    SupplementalDataStorageConfiguration

```

## Properties

`EmbeddingModelArn`

The Amazon Resource Name (ARN) of the model used to create vector embeddings for the knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-[^:]+)?:[a-z0-9-]+:[a-z0-9-]{1,20}:[0-9]{0,12}:[a-zA-Z0-9-:/._+]+)$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EmbeddingModelConfiguration`

The embeddings model configuration details for the vector model used in Knowledge Base.

_Required_: No

_Type_: [EmbeddingModelConfiguration](aws-properties-bedrock-knowledgebase-embeddingmodelconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupplementalDataStorageConfiguration`

If you include multimodal data from your data source, use this object to specify configurations for the storage location of the images extracted from your documents. These images can be retrieved and returned to the end user. They can also be used in generation when using [RetrieveAndGenerate](../../../../reference/bedrock/latest/apireference/api-agent-runtime-retrieveandgenerate.md).

_Required_: No

_Type_: [SupplementalDataStorageConfiguration](aws-properties-bedrock-knowledgebase-supplementaldatastorageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SupplementalDataStorageLocation

VideoConfiguration

All content copied from https://docs.aws.amazon.com/.
