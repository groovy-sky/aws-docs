This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule RepublishAction

Describes an action to republish to another topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Headers" : RepublishActionHeaders,
  "Qos" : Integer,
  "RoleArn" : String,
  "Topic" : String
}

```

### YAML

```yaml

  Headers:
    RepublishActionHeaders
  Qos: Integer
  RoleArn: String
  Topic: String

```

## Properties

`Headers`

MQTT Version 5.0 headers information. For more information, see [MQTT](../../../iot/latest/developerguide/mqtt.md) in the IoT
Core Developer Guide.

_Required_: No

_Type_: [RepublishActionHeaders](aws-properties-iot-topicrule-republishactionheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Qos`

The Quality of Service (QoS) level to use when republishing messages. The default value
is 0.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topic`

The name of the MQTT topic.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutItemInput

RepublishActionHeaders

All content copied from https://docs.aws.amazon.com/.
