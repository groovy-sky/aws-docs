This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel IotTopicPublish

Information required to publish the MQTT message through the AWS IoT message broker.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MqttTopic" : String,
  "Payload" : Payload
}

```

### YAML

```yaml

  MqttTopic: String
  Payload:
    Payload

```

## Properties

`MqttTopic`

The MQTT topic of the message. You can use a string expression that includes variables
( `$variable.<variable-name>`) and input values
( `$input.<input-name>.<path-to-datum>`) as the topic string.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

You can configure the action payload when you publish a message to an AWS IoT Core
topic.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-alarmmodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotSiteWise

Lambda

All content copied from https://docs.aws.amazon.com/.
