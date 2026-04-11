This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSetEventDestination EventBridgeDestination

An object that defines an Amazon EventBridge destination for email events. You can use Amazon EventBridge to
send notifications when certain email events occur.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventBusArn" : String
}

```

### YAML

```yaml

  EventBusArn: String

```

## Properties

`EventBusArn`

The Amazon Resource Name (ARN) of the Amazon EventBridge bus to publish email events to. Only the default bus is supported.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z0-9-]*:events:[a-z0-9-]+:\d{12}:event-bus/[^:]+$`

_Minimum_: `36`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DimensionConfiguration

EventDestination

All content copied from https://docs.aws.amazon.com/.
