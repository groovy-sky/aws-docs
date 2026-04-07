This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile MetricsExportConfig

Specifies the MQTT topic and role ARN required for metric export.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MqttTopic" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  MqttTopic: String
  RoleArn: String

```

## Properties

`MqttTopic`

The MQTT topic that Device Defender Detect should publish messages to for metrics
export.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

This role ARN has permission to publish MQTT messages, after which Device Defender Detect
can assume the role and publish messages on your behalf.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricDimension

MetricToRetain
