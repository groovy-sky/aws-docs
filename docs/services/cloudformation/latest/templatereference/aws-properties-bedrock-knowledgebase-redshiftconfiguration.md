This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftConfiguration

Contains configurations for an Amazon Redshift database. For more information, see [Build a knowledge base by connecting to a structured data source](../../../bedrock/latest/userguide/knowledge-base-build-structured.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueryEngineConfiguration" : RedshiftQueryEngineConfiguration,
  "QueryGenerationConfiguration" : QueryGenerationConfiguration,
  "StorageConfigurations" : [ RedshiftQueryEngineStorageConfiguration, ... ]
}

```

### YAML

```yaml

  QueryEngineConfiguration:
    RedshiftQueryEngineConfiguration
  QueryGenerationConfiguration:
    QueryGenerationConfiguration
  StorageConfigurations:
    - RedshiftQueryEngineStorageConfiguration

```

## Properties

`QueryEngineConfiguration`

Specifies configurations for an Amazon Redshift query engine.

_Required_: Yes

_Type_: [RedshiftQueryEngineConfiguration](aws-properties-bedrock-knowledgebase-redshiftqueryengineconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryGenerationConfiguration`

Specifies configurations for generating queries.

_Required_: No

_Type_: [QueryGenerationConfiguration](aws-properties-bedrock-knowledgebase-querygenerationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageConfigurations`

Specifies configurations for Amazon Redshift database storage.

_Required_: Yes

_Type_: Array of [RedshiftQueryEngineStorageConfiguration](aws-properties-bedrock-knowledgebase-redshiftqueryenginestorageconfiguration.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsFieldMapping

RedshiftProvisionedAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
