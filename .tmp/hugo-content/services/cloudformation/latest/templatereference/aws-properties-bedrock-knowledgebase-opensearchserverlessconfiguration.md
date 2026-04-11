This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase OpenSearchServerlessConfiguration

Contains details about the storage configuration of the knowledge base in Amazon OpenSearch Service. For more information, see [Create a vector index in Amazon OpenSearch Service](../../../bedrock/latest/userguide/knowledge-base-setup-oss.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectionArn" : String,
  "FieldMapping" : OpenSearchServerlessFieldMapping,
  "VectorIndexName" : String
}

```

### YAML

```yaml

  CollectionArn: String
  FieldMapping:
    OpenSearchServerlessFieldMapping
  VectorIndexName: String

```

## Properties

`CollectionArn`

The Amazon Resource Name (ARN) of the OpenSearch Service vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov|-iso):aoss:[a-z]{2}(-gov)?-[a-z]+-\d{1}:\d{12}:collection/[a-z0-9-]{3,32}$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [OpenSearchServerlessFieldMapping](aws-properties-bedrock-knowledgebase-opensearchserverlessfieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorIndexName`

The name of the vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchManagedClusterFieldMapping

OpenSearchServerlessFieldMapping

All content copied from https://docs.aws.amazon.com/.
