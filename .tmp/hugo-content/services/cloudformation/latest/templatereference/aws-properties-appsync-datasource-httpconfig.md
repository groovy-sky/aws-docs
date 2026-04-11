This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource HttpConfig

Use the `HttpConfig` property type to specify `HttpConfig` for
an AWS AppSync data source.

`HttpConfig` is a property of the [AWS::AppSync::DataSource](../userguide/aws-resource-appsync-datasource.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : AuthorizationConfig,
  "Endpoint" : String
}

```

### YAML

```yaml

  AuthorizationConfig:
    AuthorizationConfig
  Endpoint: String

```

## Properties

`AuthorizationConfig`

The authorization configuration.

_Required_: No

_Type_: [AuthorizationConfig](aws-properties-appsync-datasource-authorizationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventBridgeConfig

LambdaConfig

All content copied from https://docs.aws.amazon.com/.
