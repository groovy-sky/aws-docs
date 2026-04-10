This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase StorageConfiguration

Contains the storage configuration of the knowledge base.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MongoDbAtlasConfiguration" : MongoDbAtlasConfiguration,
  "NeptuneAnalyticsConfiguration" : NeptuneAnalyticsConfiguration,
  "OpensearchManagedClusterConfiguration" : OpenSearchManagedClusterConfiguration,
  "OpensearchServerlessConfiguration" : OpenSearchServerlessConfiguration,
  "PineconeConfiguration" : PineconeConfiguration,
  "RdsConfiguration" : RdsConfiguration,
  "S3VectorsConfiguration" : S3VectorsConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  MongoDbAtlasConfiguration:
    MongoDbAtlasConfiguration
  NeptuneAnalyticsConfiguration:
    NeptuneAnalyticsConfiguration
  OpensearchManagedClusterConfiguration:
    OpenSearchManagedClusterConfiguration
  OpensearchServerlessConfiguration:
    OpenSearchServerlessConfiguration
  PineconeConfiguration:
    PineconeConfiguration
  RdsConfiguration:
    RdsConfiguration
  S3VectorsConfiguration:
    S3VectorsConfiguration
  Type: String

```

## Properties

`MongoDbAtlasConfiguration`

Contains the storage configuration of the knowledge base in MongoDB Atlas.

_Required_: No

_Type_: [MongoDbAtlasConfiguration](aws-properties-bedrock-knowledgebase-mongodbatlasconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NeptuneAnalyticsConfiguration`

Contains details about the Neptune Analytics configuration of the knowledge base in Amazon Neptune. For more information,
see [Create a\
vector index in Amazon Neptune Analytics.](../../../bedrock/latest/userguide/knowledge-base-setup-neptune.md).

_Required_: No

_Type_: [NeptuneAnalyticsConfiguration](aws-properties-bedrock-knowledgebase-neptuneanalyticsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OpensearchManagedClusterConfiguration`

Contains details about the storage configuration of the knowledge base in OpenSearch Managed
Cluster. For more information, see [Create \
a vector index in Amazon OpenSearch Service](../../../bedrock/latest/userguide/knowledge-base-setup-osm.md).

_Required_: No

_Type_: [OpenSearchManagedClusterConfiguration](aws-properties-bedrock-knowledgebase-opensearchmanagedclusterconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OpensearchServerlessConfiguration`

Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.

_Required_: No

_Type_: [OpenSearchServerlessConfiguration](aws-properties-bedrock-knowledgebase-opensearchserverlessconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PineconeConfiguration`

Contains the storage configuration of the knowledge base in Pinecone.

_Required_: No

_Type_: [PineconeConfiguration](aws-properties-bedrock-knowledgebase-pineconeconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RdsConfiguration`

Contains details about the storage configuration of the knowledge base in Amazon RDS. For more information, see [Create a vector index in Amazon RDS](../../../bedrock/latest/userguide/knowledge-base-setup-rds.md).

_Required_: No

_Type_: [RdsConfiguration](aws-properties-bedrock-knowledgebase-rdsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3VectorsConfiguration`

The configuration settings for storing knowledge base data using S3 vectors. This includes vector index information and S3 bucket details for vector storage.

_Required_: No

_Type_: [S3VectorsConfiguration](aws-properties-bedrock-knowledgebase-s3vectorsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The vector store service in which the knowledge base is stored.

_Required_: Yes

_Type_: String

_Allowed values_: `OPENSEARCH_SERVERLESS | PINECONE | RDS | MONGO_DB_ATLAS | NEPTUNE_ANALYTICS | S3_VECTORS | OPENSEARCH_MANAGED_CLUSTER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SqlKnowledgeBaseConfiguration

SupplementalDataStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
