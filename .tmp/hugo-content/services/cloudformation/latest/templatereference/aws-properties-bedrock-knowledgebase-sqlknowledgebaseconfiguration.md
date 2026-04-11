This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase SqlKnowledgeBaseConfiguration

Contains configurations for a knowledge base connected to an SQL database. Specify the SQL database type in the `type` field and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](../../../bedrock/latest/userguide/knowledge-base-build-structured.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RedshiftConfiguration" : RedshiftConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  RedshiftConfiguration:
    RedshiftConfiguration
  Type: String

```

## Properties

`RedshiftConfiguration`

Specifies configurations for a knowledge base connected to an Amazon Redshift database.

_Required_: No

_Type_: [RedshiftConfiguration](aws-properties-bedrock-knowledgebase-redshiftconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of SQL database to connect to the knowledge base.

_Required_: Yes

_Type_: String

_Allowed values_: `REDSHIFT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3VectorsConfiguration

StorageConfiguration

All content copied from https://docs.aws.amazon.com/.
