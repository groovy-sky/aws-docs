This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace HandlerConfigs

The `HandlerConfigs` property type specifies the configuration for the `OnPublish` and `OnSubscribe`
handlers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnPublish" : HandlerConfig,
  "OnSubscribe" : HandlerConfig
}

```

### YAML

```yaml

  OnPublish:
    HandlerConfig
  OnSubscribe:
    HandlerConfig

```

## Properties

`OnPublish`

The configuration for the `OnPublish` handler.

_Required_: No

_Type_: [HandlerConfig](aws-properties-appsync-channelnamespace-handlerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnSubscribe`

The configuration for the `OnSubscribe` handler.

_Required_: No

_Type_: [HandlerConfig](aws-properties-appsync-channelnamespace-handlerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HandlerConfig

Integration

All content copied from https://docs.aws.amazon.com/.
