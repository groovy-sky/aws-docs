This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel Lambda

Calls a Lambda function, passing in information about the detector model instance and the
event that triggered the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionArn" : String,
  "Payload" : Payload
}

```

### YAML

```yaml

  FunctionArn: String
  Payload:
    Payload

```

## Properties

`FunctionArn`

The ARN of the Lambda function that is executed.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

You can configure the action payload when you send a message to a Lambda function.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-detectormodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotTopicPublish

OnEnter

All content copied from https://docs.aws.amazon.com/.
