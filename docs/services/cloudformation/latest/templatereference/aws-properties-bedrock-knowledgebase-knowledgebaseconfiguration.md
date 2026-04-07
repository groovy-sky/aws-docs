This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase KnowledgeBaseConfiguration

Configurations to apply to a knowledge base attached to the agent during query. For more information, see [Knowledge base retrieval configurations](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-session-state.html#session-state-kb).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KendraKnowledgeBaseConfiguration" : KendraKnowledgeBaseConfiguration,
  "SqlKnowledgeBaseConfiguration" : SqlKnowledgeBaseConfiguration,
  "Type" : String,
  "VectorKnowledgeBaseConfiguration" : VectorKnowledgeBaseConfiguration
}

```

### YAML

```yaml

  KendraKnowledgeBaseConfiguration:
    KendraKnowledgeBaseConfiguration
  SqlKnowledgeBaseConfiguration:
    SqlKnowledgeBaseConfiguration
  Type: String
  VectorKnowledgeBaseConfiguration:
    VectorKnowledgeBaseConfiguration

```

## Properties

`KendraKnowledgeBaseConfiguration`

Settings for an Amazon Kendra knowledge base.

_Required_: No

_Type_: [KendraKnowledgeBaseConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-knowledgebase-kendraknowledgebaseconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SqlKnowledgeBaseConfiguration`

Specifies configurations for a knowledge base connected to an SQL database.

_Required_: No

_Type_: [SqlKnowledgeBaseConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-knowledgebase-sqlknowledgebaseconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of data that the data source is converted into for the knowledge base.

_Required_: Yes

_Type_: String

_Allowed values_: `VECTOR | KENDRA | SQL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorKnowledgeBaseConfiguration`

Contains details about the model that's used to convert the data source into vector embeddings.

_Required_: No

_Type_: [VectorKnowledgeBaseConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KendraKnowledgeBaseConfiguration

MongoDbAtlasConfiguration
