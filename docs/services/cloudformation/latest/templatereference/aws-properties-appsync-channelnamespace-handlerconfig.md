This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace HandlerConfig

The `HandlerConfig` property type specifies the configuration for the handler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Behavior" : String,
  "Integration" : Integration
}

```

### YAML

```yaml

  Behavior: String
  Integration:
    Integration

```

## Properties

`Behavior`

The behavior for the handler.

_Required_: Yes

_Type_: String

_Allowed values_: `CODE | DIRECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Integration`

The integration data source configuration for the handler.

_Required_: Yes

_Type_: [Integration](aws-properties-appsync-channelnamespace-integration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthMode

HandlerConfigs

All content copied from https://docs.aws.amazon.com/.
