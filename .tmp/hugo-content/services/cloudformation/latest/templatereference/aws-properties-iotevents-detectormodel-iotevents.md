This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel IotEvents

Sends an AWS IoT Events input, passing in information about the detector model instance and the
event that triggered the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputName" : String,
  "Payload" : Payload
}

```

### YAML

```yaml

  InputName: String
  Payload:
    Payload

```

## Properties

`InputName`

The name of the AWS IoT Events input where the data is sent.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

You can configure the action payload when you send a message to an AWS IoT Events input.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-detectormodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Firehose

IotSiteWise

All content copied from https://docs.aws.amazon.com/.
