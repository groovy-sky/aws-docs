This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase PineconeConfiguration

Contains details about the storage configuration of the knowledge base in Pinecone. For more information, see [Create a vector index in Pinecone](../../../bedrock/latest/userguide/knowledge-base-setup-pinecone.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionString" : String,
  "CredentialsSecretArn" : String,
  "FieldMapping" : PineconeFieldMapping,
  "Namespace" : String
}

```

### YAML

```yaml

  ConnectionString:
    String
  CredentialsSecretArn: String
  FieldMapping:
    PineconeFieldMapping
  Namespace: String

```

## Properties

`ConnectionString`

The endpoint URL for your index management page.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CredentialsSecretArn`

The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [PineconeFieldMapping](aws-properties-bedrock-knowledgebase-pineconefieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The namespace to be used to write new data to your database.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchServerlessFieldMapping

PineconeFieldMapping

All content copied from https://docs.aws.amazon.com/.
