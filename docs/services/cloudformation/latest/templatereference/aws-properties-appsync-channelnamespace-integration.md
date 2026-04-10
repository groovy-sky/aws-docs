This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace Integration

The `Integration` property type specifies the integration data source configuration for the handler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceName" : String,
  "LambdaConfig" : LambdaConfig
}

```

### YAML

```yaml

  DataSourceName: String
  LambdaConfig:
    LambdaConfig

```

## Properties

`DataSourceName`

The unique name of the data source that has been configured on the API.

_Required_: Yes

_Type_: String

_Pattern_: `([_A-Za-z][_0-9A-Za-z]{0,511})?`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaConfig`

The configuration for a Lambda data source.

_Required_: No

_Type_: [LambdaConfig](aws-properties-appsync-channelnamespace-lambdaconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HandlerConfigs

LambdaConfig

All content copied from https://docs.aws.amazon.com/.
