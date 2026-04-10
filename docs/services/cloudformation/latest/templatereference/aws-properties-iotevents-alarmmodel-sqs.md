This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel Sqs

Sends information about the detector model instance and the event that triggered the
action to an Amazon SQS queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Payload" : Payload,
  "QueueUrl" : String,
  "UseBase64" : Boolean
}

```

### YAML

```yaml

  Payload:
    Payload
  QueueUrl: String
  UseBase64: Boolean

```

## Properties

`Payload`

You can configure the action payload when you send a message to an Amazon SQS
queue.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-alarmmodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueUrl`

The URL of the SQS queue where the data is written.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBase64`

Set this to TRUE if you want the data to be base-64 encoded before it is written to the
queue. Otherwise, set this to FALSE.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sns

Tag

All content copied from https://docs.aws.amazon.com/.
