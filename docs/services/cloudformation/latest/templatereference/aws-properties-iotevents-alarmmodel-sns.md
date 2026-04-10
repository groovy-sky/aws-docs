This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel Sns

Information required to publish the Amazon SNS message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Payload" : Payload,
  "TargetArn" : String
}

```

### YAML

```yaml

  Payload:
    Payload
  TargetArn: String

```

## Properties

`Payload`

You can configure the action payload when you send a message as an Amazon SNS push
notification.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-alarmmodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN of the Amazon SNS target where the message is sent.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SimpleRule

Sqs

All content copied from https://docs.aws.amazon.com/.
