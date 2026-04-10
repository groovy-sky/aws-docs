This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftProvisionedConfiguration

Contains configurations for a provisioned Amazon Redshift query engine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthConfiguration" : RedshiftProvisionedAuthConfiguration,
  "ClusterIdentifier" : String
}

```

### YAML

```yaml

  AuthConfiguration:
    RedshiftProvisionedAuthConfiguration
  ClusterIdentifier: String

```

## Properties

`AuthConfiguration`

Specifies configurations for authentication to Amazon Redshift.

_Required_: Yes

_Type_: [RedshiftProvisionedAuthConfiguration](aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterIdentifier`

The ID of the Amazon Redshift cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftProvisionedAuthConfiguration

RedshiftQueryEngineAwsDataCatalogStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
