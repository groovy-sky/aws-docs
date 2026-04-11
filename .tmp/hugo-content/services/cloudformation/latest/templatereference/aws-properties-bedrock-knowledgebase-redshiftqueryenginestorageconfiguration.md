This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftQueryEngineStorageConfiguration

Contains configurations for Amazon Redshift data storage. Specify the data storage service to use in the `type` field and include the corresponding field. For more information, see [Build a knowledge base by connecting to a structured data source](../../../bedrock/latest/userguide/knowledge-base-build-structured.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsDataCatalogConfiguration" : RedshiftQueryEngineAwsDataCatalogStorageConfiguration,
  "RedshiftConfiguration" : RedshiftQueryEngineRedshiftStorageConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  AwsDataCatalogConfiguration:
    RedshiftQueryEngineAwsDataCatalogStorageConfiguration
  RedshiftConfiguration:
    RedshiftQueryEngineRedshiftStorageConfiguration
  Type: String

```

## Properties

`AwsDataCatalogConfiguration`

Specifies configurations for storage in AWS Glue Data Catalog.

_Required_: No

_Type_: [RedshiftQueryEngineAwsDataCatalogStorageConfiguration](aws-properties-bedrock-knowledgebase-redshiftqueryengineawsdatacatalogstorageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RedshiftConfiguration`

Specifies configurations for storage in Amazon Redshift.

_Required_: No

_Type_: [RedshiftQueryEngineRedshiftStorageConfiguration](aws-properties-bedrock-knowledgebase-redshiftqueryengineredshiftstorageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The data storage service to use.

_Required_: Yes

_Type_: String

_Allowed values_: `REDSHIFT | AWS_DATA_CATALOG`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftQueryEngineRedshiftStorageConfiguration

RedshiftServerlessAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
