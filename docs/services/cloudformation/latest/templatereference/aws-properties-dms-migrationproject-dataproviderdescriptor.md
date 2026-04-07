This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::MigrationProject DataProviderDescriptor

Information about a data provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataProviderArn" : String,
  "DataProviderIdentifier" : String,
  "DataProviderName" : String,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String
}

```

### YAML

```yaml

  DataProviderArn: String
  DataProviderIdentifier: String
  DataProviderName: String
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String

```

## Properties

`DataProviderArn`

The Amazon Resource Name (ARN) of the data provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProviderIdentifier`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProviderName`

The user-friendly name of the data provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The ARN of the role used to access AWS Secrets Manager.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecretId`

The identifier of the AWS Secrets Manager Secret used to store access credentials for the data provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DMS::MigrationProject

SchemaConversionApplicationAttributes
