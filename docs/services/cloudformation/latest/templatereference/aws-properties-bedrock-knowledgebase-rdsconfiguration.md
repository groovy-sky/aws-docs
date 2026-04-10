This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RdsConfiguration

Contains details about the storage configuration of the knowledge base in Amazon RDS. For more information, see [Create a vector index in Amazon RDS](../../../bedrock/latest/userguide/knowledge-base-setup-rds.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialsSecretArn" : String,
  "DatabaseName" : String,
  "FieldMapping" : RdsFieldMapping,
  "ResourceArn" : String,
  "TableName" : String
}

```

### YAML

```yaml

  CredentialsSecretArn: String
  DatabaseName: String
  FieldMapping:
    RdsFieldMapping
  ResourceArn: String
  TableName: String

```

## Properties

`CredentialsSecretArn`

The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of your Amazon RDS database.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldMapping`

Contains the names of the fields to which to map information about the vector store.

_Required_: Yes

_Type_: [RdsFieldMapping](aws-properties-bedrock-knowledgebase-rdsfieldmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArn`

The Amazon Resource Name (ARN) of the vector store.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):rds:[a-zA-Z0-9-]*:[0-9]{12}:cluster:[a-zA-Z0-9-]{1,63}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableName`

The name of the table in the database.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\.\-]+$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryGenerationTable

RdsFieldMapping

All content copied from https://docs.aws.amazon.com/.
