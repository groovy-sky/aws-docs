This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign DataDestinationConfig

The destination where the AWS IoT FleetWise campaign sends data. You can send data to be stored in Amazon S3 or Amazon Timestream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MqttTopicConfig" : MqttTopicConfig,
  "S3Config" : S3Config,
  "TimestreamConfig" : TimestreamConfig
}

```

### YAML

```yaml

  MqttTopicConfig:
    MqttTopicConfig
  S3Config:
    S3Config
  TimestreamConfig:
    TimestreamConfig

```

## Properties

`MqttTopicConfig`

The MQTT topic to which the AWS IoT FleetWise campaign routes data.

_Required_: No

_Type_: [MqttTopicConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-mqtttopicconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Config`

The Amazon S3 bucket where the AWS IoT FleetWise campaign sends data.

_Required_: No

_Type_: [S3Config](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-s3config.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimestreamConfig`

The Amazon Timestream table where the campaign sends data.

_Required_: No

_Type_: [TimestreamConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-timestreamconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConditionBasedSignalFetchConfig

DataPartition
