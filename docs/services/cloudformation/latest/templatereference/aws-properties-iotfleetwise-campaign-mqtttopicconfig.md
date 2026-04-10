This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign MqttTopicConfig

The MQTT topic to which the AWS IoT FleetWise campaign routes data. For more information, see [Device communication \
protocols](../../../iot/latest/developerguide/protocols.md) in the _AWS IoT Core Developer Guide_.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionRoleArn" : String,
  "MqttTopicArn" : String
}

```

### YAML

```yaml

  ExecutionRoleArn: String
  MqttTopicArn: String

```

## Properties

`ExecutionRoleArn`

The ARN of the role that grants AWS IoT FleetWise permission to access and act on messages sent to
the MQTT topic.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):iam::(\d{12})?:(role((\u002F)|(\u002F[\u0021-\u007F]+\u002F))[\w+=,.@-]+)$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MqttTopicArn`

The ARN of the MQTT topic.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPartitionUploadOptions

S3Config

All content copied from https://docs.aws.amazon.com/.
