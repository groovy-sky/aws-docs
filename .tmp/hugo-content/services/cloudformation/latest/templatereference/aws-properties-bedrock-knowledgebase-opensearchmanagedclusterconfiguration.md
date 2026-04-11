This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase OpenSearchManagedClusterConfiguration

Contains details about the Managed Cluster configuration of the knowledge base in Amazon OpenSearch Service. For more information,
see [Create a vector index in OpenSearch Managed Cluster](../../../bedrock/latest/userguide/knowledge-base-setup-osm.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainArn" : String,
  "DomainEndpoint" : String,
  "FieldMapping" : OpenSearchManagedClusterFieldMapping,
  "VectorIndexName" : String
}

```

### YAML

```yaml

  DomainArn: String
  DomainEndpoint: String
  FieldMapping:
    OpenSearchManagedClusterFieldMapping
  VectorIndexName: String

```

## Properties

`DomainArn`

The Amazon Resource Name (ARN) of the OpenSearch domain.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-eusc|-iso(-[b-f])?)?:es:([a-z]{2,}-){2,}\d:\d{12}:domain/[a-z][a-z0-9-]{3,28}$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainEndpoint`

The endpoint URL the OpenSearch domain.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [OpenSearchManagedClusterFieldMapping](aws-properties-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorIndexName`

The name of the vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^(?![\-_+.])[a-z0-9][a-z0-9\-_\.]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NeptuneAnalyticsFieldMapping

OpenSearchManagedClusterFieldMapping

All content copied from https://docs.aws.amazon.com/.
