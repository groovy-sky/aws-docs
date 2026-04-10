This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource RelationalDatabaseConfig

Use the `RelationalDatabaseConfig` property type to specify
`RelationalDatabaseConfig` for an AWS AppSync data source.

`RelationalDatabaseConfig` is a property of the [AWS::AppSync::DataSource](../userguide/aws-resource-appsync-datasource.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RdsHttpEndpointConfig" : RdsHttpEndpointConfig,
  "RelationalDatabaseSourceType" : String
}

```

### YAML

```yaml

  RdsHttpEndpointConfig:
    RdsHttpEndpointConfig
  RelationalDatabaseSourceType: String

```

## Properties

`RdsHttpEndpointConfig`

Information about the Amazon RDS resource.

_Required_: No

_Type_: [RdsHttpEndpointConfig](aws-properties-appsync-datasource-rdshttpendpointconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalDatabaseSourceType`

The type of relational data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsHttpEndpointConfig

AWS::AppSync::DomainName

All content copied from https://docs.aws.amazon.com/.
