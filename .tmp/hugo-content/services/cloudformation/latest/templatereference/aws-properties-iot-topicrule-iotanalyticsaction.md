This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule IotAnalyticsAction

Sends message data to an AWS IoT Analytics channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchMode" : Boolean,
  "ChannelName" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  BatchMode: Boolean
  ChannelName: String
  RoleArn: String

```

## Properties

`BatchMode`

Whether to process the action as a batch. The default value is
`false`.

When `batchMode` is `true` and the rule SQL statement evaluates
to an Array, each Array element is delivered as a separate message when passed by [`BatchPutMessage`](../../../../reference/iotanalytics/latest/apireference/api-batchputmessage.md) The resulting array can't have more than 100 messages.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelName`

The name of the IoT Analytics channel to which message data will be sent.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role which has a policy that grants IoT Analytics permission to send
message data via IoT Analytics (iotanalytics:BatchPutMessage).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [IotAnalyticsAction](../../../../reference/iot/latest/apireference/api-iotanalyticsaction.md) in the _AWS IoT API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpAuthorization

IotEventsAction

All content copied from https://docs.aws.amazon.com/.
