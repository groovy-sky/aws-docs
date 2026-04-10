This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Channel

The AWS::IoTAnalytics::Channel resource collects data from an MQTT topic and archives the
raw, unprocessed messages before publishing the data to a pipeline. For more information, see
[How to Use AWS IoT Analytics](../../../iotanalytics/latest/userguide/welcome.md#aws-iot-analytics-how) in the _AWS IoT Analytics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTAnalytics::Channel",
  "Properties" : {
      "ChannelName" : String,
      "ChannelStorage" : ChannelStorage,
      "RetentionPeriod" : RetentionPeriod,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTAnalytics::Channel
Properties:
  ChannelName: String
  ChannelStorage:
    ChannelStorage
  RetentionPeriod:
    RetentionPeriod
  Tags:
    - Tag

```

## Properties

`ChannelName`

The name of the channel.

_Required_: No

_Type_: String

_Pattern_: `(^(?!_{2}))(^[a-zA-Z0-9_]+$)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelStorage`

Where channel data is stored.

_Required_: No

_Type_: [ChannelStorage](aws-properties-iotanalytics-channel-channelstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriod`

How long, in days, message data is kept for the channel.

_Required_: No

_Type_: [RetentionPeriod](aws-properties-iotanalytics-channel-retentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the channel.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-iotanalytics-channel-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Simple Channel](#aws-resource-iotanalytics-channel--examples--Simple_Channel)

- [Complex Channel](#aws-resource-iotanalytics-channel--examples--Complex_Channel)

### Simple Channel

The following example creates a simple channel that uses service-managed channel storage.

#### JSON

```json

{
    "Description": "Create a simple Channel",
    "Resources": {
        "Channel": {
            "Type": "AWS::IoTAnalytics::Channel",
            "Properties": {
                "ChannelName": "SimpleChannel"
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a simple Channel"
Resources:
  Channel:
    Type: "AWS::IoTAnalytics::Channel"
    Properties:
      ChannelName: "SimpleChannel"
```

### Complex Channel

The following example creates a complex channel.

#### JSON

```json

{
    "Description": "Create a complex channel",
    "Resources": {
        "Channel": {
            "Type": "AWS::IoTAnalytics::Channel",
            "Properties": {
                "ChannelName": "ComplexChannel",
                "RetentionPeriod": {
                    "Unlimited": false,
                    "NumberOfDays": 10
                },
                "Tags": [
                    {
                        "Key": "keyname1",
                        "Value": "value1"
                    },
                    {
                        "Key": "keyname2",
                        "Value": "value2"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a complex channel"
Resources:
  Channel:
    Type: "AWS::IoTAnalytics::Channel"
    Properties:
      ChannelName: "ComplexChannel"
      RetentionPeriod:
        Unlimited: false
        NumberOfDays: 10
      Tags:
        -
          Key: "keyname1"
          Value: "value1"
        -
          Key: "keyname2"
          Value: "value2"
```

## See also

- [How to Use AWS IoT Analytics](../../../iotanalytics/latest/userguide/welcome.md#aws-iot-analytics-how) in the _AWS IoT Analytics User Guide_

- [CreateChannel](../../../../reference/iotanalytics/latest/apireference/api-createchannel.md) in the
_AWS IoT Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT Analytics

ChannelStorage

All content copied from https://docs.aws.amazon.com/.
