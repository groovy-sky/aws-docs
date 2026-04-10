This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource RdsHttpEndpointConfig

Use the `RdsHttpEndpointConfig` property type to specify the
`RdsHttpEndpoint` for an AWS AppSync relational
database.

`RdsHttpEndpointConfig` is a property of the [AWS AppSync DataSource RelationalDatabaseConfig](../userguide/aws-properties-appsync-datasource-relationaldatabaseconfig.md)
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsRegion" : String,
  "AwsSecretStoreArn" : String,
  "DatabaseName" : String,
  "DbClusterIdentifier" : String,
  "Schema" : String
}

```

### YAML

```yaml

  AwsRegion: String
  AwsSecretStoreArn: String
  DatabaseName: String
  DbClusterIdentifier: String
  Schema: String

```

## Properties

`AwsRegion`

AWS Region for RDS HTTP endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsSecretStoreArn`

The ARN for database credentials stored in AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

Logical database name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbClusterIdentifier`

Amazon RDS cluster Amazon Resource Name (ARN).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

Logical schema name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchServiceConfig

RelationalDatabaseConfig

All content copied from https://docs.aws.amazon.com/.
