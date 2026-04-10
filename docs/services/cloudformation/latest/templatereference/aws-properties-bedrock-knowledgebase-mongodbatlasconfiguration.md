This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase MongoDbAtlasConfiguration

Contains details about the storage configuration of the knowledge base in MongoDB Atlas.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectionName" : String,
  "CredentialsSecretArn" : String,
  "DatabaseName" : String,
  "Endpoint" : String,
  "EndpointServiceName" : String,
  "FieldMapping" : MongoDbAtlasFieldMapping,
  "TextIndexName" : String,
  "VectorIndexName" : String
}

```

### YAML

```yaml

  CollectionName: String
  CredentialsSecretArn: String
  DatabaseName: String
  Endpoint: String
  EndpointServiceName: String
  FieldMapping:
    MongoDbAtlasFieldMapping
  TextIndexName: String
  VectorIndexName: String

```

## Properties

`CollectionName`

The collection name of the knowledge base in MongoDB Atlas.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CredentialsSecretArn`

The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that contains user credentials for your MongoDB Atlas cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The database name in your MongoDB Atlas cluster for your knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Endpoint`

The endpoint URL of your MongoDB Atlas cluster for your knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.mongodb\.net$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointServiceName`

The name of the VPC endpoint service in your account that is connected to your MongoDB Atlas cluster.

_Required_: No

_Type_: String

_Pattern_: `^(?:arn:aws(?:-us-gov|-cn|-iso|-iso-[a-z])*:.+:.*:\d+:.+/.+$|[a-zA-Z0-9*]+[a-zA-Z0-9._-]*)$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [MongoDbAtlasFieldMapping](aws-properties-bedrock-knowledgebase-mongodbatlasfieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextIndexName`

The name of the text search index in the MongoDB collection. This is required for using the hybrid search
feature.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorIndexName`

The name of the MongoDB Atlas vector search index.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KnowledgeBaseConfiguration

MongoDbAtlasFieldMapping

All content copied from https://docs.aws.amazon.com/.
