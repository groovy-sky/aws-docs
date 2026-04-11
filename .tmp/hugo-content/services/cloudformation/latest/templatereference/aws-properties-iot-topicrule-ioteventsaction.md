This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule IotEventsAction

Sends an input to an AWS IoT Events detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchMode" : Boolean,
  "InputName" : String,
  "MessageId" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  BatchMode: Boolean
  InputName: String
  MessageId: String
  RoleArn: String

```

## Properties

`BatchMode`

Whether to process the event actions as a batch. The default value is
`false`.

When `batchMode` is `true`, you can't specify a
`messageId`.

When `batchMode` is `true` and the rule SQL statement evaluates
to an Array, each Array element is treated as a separate message when Events by calling
[`BatchPutMessage`](../../../../reference/iotevents/latest/apireference/api-iotevents-data-batchputmessage.md). The resulting array can't have more than 10 messages.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputName`

The name of the AWS IoT Events input.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageId`

The ID of the message. The default `messageId` is a new UUID value.

When `batchMode` is `true`, you can't specify a
`messageId`--a new UUID value will be assigned.

Assign a value to this property to ensure that only one input (message) with a given
`messageId` will be processed by an AWS IoT Events detector.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT permission to send an input to an AWS IoT
Events detector. ("Action":"iotevents:BatchPutMessage").

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotAnalyticsAction

IotSiteWiseAction

All content copied from https://docs.aws.amazon.com/.
