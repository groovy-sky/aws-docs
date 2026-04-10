This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftQueryEngineConfiguration

Contains configurations for an Amazon Redshift query engine. Specify the type of query engine in `type` and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](../../../bedrock/latest/userguide/knowledge-base-build-structured.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProvisionedConfiguration" : RedshiftProvisionedConfiguration,
  "ServerlessConfiguration" : RedshiftServerlessConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  ProvisionedConfiguration:
    RedshiftProvisionedConfiguration
  ServerlessConfiguration:
    RedshiftServerlessConfiguration
  Type: String

```

## Properties

`ProvisionedConfiguration`

Specifies configurations for a provisioned Amazon Redshift query engine.

_Required_: No

_Type_: [RedshiftProvisionedConfiguration](aws-properties-bedrock-knowledgebase-redshiftprovisionedconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerlessConfiguration`

Specifies configurations for a serverless Amazon Redshift query engine.

_Required_: No

_Type_: [RedshiftServerlessConfiguration](aws-properties-bedrock-knowledgebase-redshiftserverlessconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of query engine.

_Required_: Yes

_Type_: String

_Allowed values_: `SERVERLESS | PROVISIONED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftQueryEngineAwsDataCatalogStorageConfiguration

RedshiftQueryEngineRedshiftStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
