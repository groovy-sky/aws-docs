This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource DynamoDBConfig

The `DynamoDBConfig` property type specifies the `AwsRegion` and
`TableName` for an Amazon DynamoDB table in your account for an
AWS AppSync data source.

`DynamoDBConfig` is a property of the [AWS::AppSync::DataSource](../userguide/aws-resource-appsync-datasource.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsRegion" : String,
  "DeltaSyncConfig" : DeltaSyncConfig,
  "TableName" : String,
  "UseCallerCredentials" : Boolean,
  "Versioned" : Boolean
}

```

### YAML

```yaml

  AwsRegion: String
  DeltaSyncConfig:
    DeltaSyncConfig
  TableName: String
  UseCallerCredentials: Boolean
  Versioned: Boolean

```

## Properties

`AwsRegion`

The AWS Region.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeltaSyncConfig`

The `DeltaSyncConfig` for a versioned datasource.

_Required_: No

_Type_: [DeltaSyncConfig](aws-properties-appsync-datasource-deltasyncconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The table name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseCallerCredentials`

Set to `TRUE` to use AWS Identity and Access Management with this data source.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Versioned`

Set to TRUE to use Conflict Detection and Resolution with this data source.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeltaSyncConfig

EventBridgeConfig

All content copied from https://docs.aws.amazon.com/.
