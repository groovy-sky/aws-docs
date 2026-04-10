This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaEventSource

Contains information about an event source for an AWS Lambda function. The
event source defines the topics on which this Lambda function subscribes to
receive messages that run the function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Topic" : String,
  "Type" : String
}

```

### YAML

```yaml

  Topic: String
  Type: String

```

## Properties

`Topic`

The topic to which to subscribe to receive event messages.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of event source. Choose from the following options:

- `PUB_SUB` – Subscribe to local publish/subscribe messages. This event
source type doesn't support MQTT wildcards ( `+` and `#`) in the
event source topic.

- `IOT_CORE` – Subscribe to AWS IoT Core MQTT messages. This
event source type supports MQTT wildcards ( `+` and `#`) in the event
source topic.

_Required_: No

_Type_: String

_Allowed values_: `PUB_SUB | IOT_CORE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaDeviceMount

LambdaExecutionParameters

All content copied from https://docs.aws.amazon.com/.
