This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource RelationalDatabaseConfig

Use the `RelationalDatabaseConfig` property type to specify
`RelationalDatabaseConfig` for an AWS AppSync data source.

`RelationalDatabaseConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.

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

_Type_: [RdsHttpEndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-datasource-rdshttpendpointconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalDatabaseSourceType`

The type of relational data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RdsHttpEndpointConfig

AWS::AppSync::DomainName
