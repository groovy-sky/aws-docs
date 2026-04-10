This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel Firehose

Sends information about the detector model instance and the event that triggered the
action to an Amazon Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStreamName" : String,
  "Payload" : Payload,
  "Separator" : String
}

```

### YAML

```yaml

  DeliveryStreamName: String
  Payload:
    Payload
  Separator: String

```

## Properties

`DeliveryStreamName`

The name of the Kinesis Data Firehose delivery stream where the data is written.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

You can configure the action payload when you send a message to an Amazon Data Firehose delivery
stream.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-detectormodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Separator`

A character separator that is used to separate records written to the Kinesis Data
Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows
newline), ',' (comma).

_Required_: No

_Type_: String

_Pattern_: `([\n\t])|(\r\n)|(,)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Event

IotEvents

All content copied from https://docs.aws.amazon.com/.
