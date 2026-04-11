This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline

The AWS::IoTAnalytics::Pipeline resource consumes messages from one or more channels and allows
you to process the messages before storing them in a data store. You must specify both a
`channel` and a `datastore` activity and, optionally, as many
as 23 additional activities in the `pipelineActivities` array. For more information, see
[How to Use AWS IoT Analytics](../../../iotanalytics/latest/userguide/welcome.md#aws-iot-analytics-how) in the _AWS IoT Analytics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTAnalytics::Pipeline",
  "Properties" : {
      "PipelineActivities" : [ Activity, ... ],
      "PipelineName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTAnalytics::Pipeline
Properties:
  PipelineActivities:
    - Activity
  PipelineName: String
  Tags:
    - Tag

```

## Properties

`PipelineActivities`

A list of "PipelineActivity" objects. Activities perform transformations on your messages,
such as removing, renaming or adding message attributes; filtering messages based on attribute
values; invoking your Lambda functions on messages for advanced processing; or performing
mathematical transformations to normalize device data.

The list can be 2-25 **PipelineActivity** objects and must
contain both a `channel` and a `datastore` activity. Each entry in the
list must contain only one activity, for example:

`pipelineActivities = [
{
    "channel": { ... }
},
{
    "lambda": { ... }
},
...
]`

_Required_: Yes

_Type_: Array of [Activity](aws-properties-iotanalytics-pipeline-activity.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineName`

The name of the pipeline.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata which can be used to manage the pipeline.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-iotanalytics-pipeline-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Simple Pipeline](#aws-resource-iotanalytics-pipeline--examples--Simple_Pipeline)

- [Complex Pipeline](#aws-resource-iotanalytics-pipeline--examples--Complex_Pipeline)

### Simple Pipeline

The following example creates a simple pipeline.

#### JSON

```json

{
    "Description": "Create a simple Pipeline",
    "Resources": {
        "Pipeline": {
            "Type": "AWS::IoTAnalytics::Pipeline",
            "Properties": {
                "PipelineName": "SimplePipeline",
                "PipelineActivities": [
                    {
                        "Channel": {
                            "Name": "ChannelActivity",
                            "ChannelName": "SimpleChannel",
                            "Next": "DatastoreActivity"
                        },
                        "Datastore": {
                            "Name": "DatastoreActivity",
                            "DatastoreName": "SimpleDatastore"
                        }
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
Description: "Create a simple Pipeline"
Resources:
  Pipeline:
    Type: "AWS::IoTAnalytics::Pipeline"
    Properties:
      PipelineName: "SimplePipeline"
      PipelineActivities:
        -
          Channel:
            Name: "ChannelActivity"
            ChannelName: "SimpleChannel"
            Next: "DatastoreActivity"
          Datastore:
            Name: "DatastoreActivity"
            DatastoreName: "SimpleDatastore"
```

### Complex Pipeline

The following example creates a complex pipeline.

#### JSON

```json

{
    "Description": "Create a complex Pipeline",
    "Resources": {
        "Pipeline": {
            "Type": "AWS::IoTAnalytics::Pipeline",
            "Properties": {
                "PipelineName": "ComplexPipeline",
                "PipelineActivities": [
                    {
                        "Channel": {
                            "Name": "ChannelActivity",
                            "ChannelName": "Channel",
                            "Next": "LambdaActivity"
                        },
                        "Lambda": {
                            "Name": "LambdaActivity",
                            "LambdaName": "Lambda",
                            "BatchSize": 1,
                            "Next": "AddAttributesActivity"
                        },
                        "AddAttributes": {
                            "Name": "AddAttributesActivity",
                            "Attributes": {
                                "key1": "attribute1",
                                "key2": "attribute2"
                            },
                            "Next": "RemoveAttributesActivity"
                        },
                        "RemoveAttributes": {
                            "Name": "RemoveAttributesActivity",
                            "Attributes": [
                                "attribute1",
                                "attribute2"
                            ],
                            "Next": "SelectAttributesActivity"
                        },
                        "SelectAttributes": {
                            "Name": "SelectAttributesActivity",
                            "Attributes": [
                                "attribute1",
                                "attribute2"
                            ],
                            "Next": "FilterActivity"
                        },
                        "Filter": {
                            "Name": "FilterActivity",
                            "Filter": "attribute1 > 40 AND attribute2 < 20",
                            "Next": "MathActivity"
                        },
                        "Math": {
                            "Name": "MathActivity",
                            "Attribute": "attribute",
                            "Math": "attribute - 10",
                            "Next": "DeviceRegistryEnrichActivity"
                        },
                        "DeviceRegistryEnrich": {
                            "Name": "DeviceRegistryEnrichActivity",
                            "Attribute": "attribute",
                            "ThingName": "thingName",
                            "RoleArn": "arn:aws:iam::<your_Account_Id>:role/Enrich",
                            "Next": "DeviceShadowEnrichActivity"
                        },
                        "DeviceShadowEnrich": {
                            "Name": "DeviceShadowEnrichActivity",
                            "Attribute": "attribute",
                            "ThingName": "thingName",
                            "RoleArn": "arn:aws:iam::<your_Account_Id>:role/Enrich",
                            "Next": "DatastoreActivity"
                        },
                        "Datastore": {
                            "Name": "DatastoreActivity",
                            "DatastoreName": "Datastore"
                        }
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
Description: "Create a complex Pipeline"
Resources:
  Pipeline:
    Type: "AWS::IoTAnalytics::Pipeline"
    Properties:
      PipelineName: "ComplexPipeline"
      PipelineActivities:
        -
          Channel:
            Name: "ChannelActivity"
            ChannelName: "Channel"
            Next: "LambdaActivity"
          Lambda:
            Name: "LambdaActivity"
            LambdaName: "Lambda"
            BatchSize: 1
            Next: "AddAttributesActivity"
          AddAttributes:
            Name: "AddAttributesActivity"
            Attributes:
              key1: "attribute1"
              key2: "attribute2"
            Next: "RemoveAttributesActivity"
          RemoveAttributes:
            Name: "RemoveAttributesActivity"
            Attributes:
              -
                "attribute1"
              -
                "attribute2"
            Next: "SelectAttributesActivity"
          SelectAttributes:
            Name: "SelectAttributesActivity"
            Attributes:
              -
                "attribute1"
              -
                "attribute2"
            Next: "FilterActivity"
          Filter:
            Name: "FilterActivity"
            Filter: "attribute1 > 40 AND attribute2 < 20"
            Next: "MathActivity"
          Math:
            Name: "MathActivity"
            Attribute: "attribute"
            Math: "attribute - 10"
            Next: "DeviceRegistryEnrichActivity"
          DeviceRegistryEnrich:
            Name: "DeviceRegistryEnrichActivity"
            Attribute: "attribute"
            ThingName: "thingName"
            RoleArn: "arn:aws:iam::<your_Account_Id>:role/Enrich"
            Next: "DeviceShadowEnrichActivity"
          DeviceShadowEnrich:
            Name: "DeviceShadowEnrichActivity"
            Attribute: "attribute"
            ThingName: "thingName"
            RoleArn: "arn:aws:iam::<your_Account_Id>:role/Enrich"
            Next: "DatastoreActivity"
          Datastore:
            Name: "DatastoreActivity"
            DatastoreName: "Datastore"
```

## See also

- [How to Use AWS IoT Analytics](../../../iotanalytics/latest/userguide/welcome.md#aws-iot-analytics-how)
in the _AWS IoT Analytics User Guide_

- [CreatePipeline](../../../../reference/iotanalytics/latest/apireference/api-createpipeline.md) in the
_AWS IoT Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimestampPartition

Activity

All content copied from https://docs.aws.amazon.com/.
