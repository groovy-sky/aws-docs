This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSetEventDestination

Specifies a configuration set event destination. _Events_ include
message sends, deliveries, opens, clicks, bounces, and complaints. _Event_
_destinations_ are places that you can send information about these events
to. For example, you can send event data to Amazon SNS to receive notifications when you
receive bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream
data to Amazon S3 for long-term storage.

A single configuration set can include more than one event destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ConfigurationSetEventDestination",
  "Properties" : {
      "ConfigurationSetName" : String,
      "EventDestination" : EventDestination
    }
}

```

### YAML

```yaml

Type: AWS::SES::ConfigurationSetEventDestination
Properties:
  ConfigurationSetName: String
  EventDestination:
    EventDestination

```

## Properties

`ConfigurationSetName`

The name of the configuration set that contains the event destination.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventDestination`

An object that defines the event destination.

_Required_: Yes

_Type_: [EventDestination](aws-properties-ses-configurationseteventdestination-eventdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

Specifies an event destination for a configuration set.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES ConfigurationSetEventDestination Sample Template",
    "Parameters": {
        "ConfigSetName": {
            "Type": "String"
        },
        "EventDestinationName": {
            "Type": "String"
        },
        "EventType1": {
            "Type": "String"
        },
        "EventType2": {
            "Type": "String"
        },
        "EventType3": {
            "Type": "String"
        },
        "DimensionName1": {
            "Type": "String"
        },
        "DimensionValueSource1": {
            "Type": "String"
        },
        "DefaultDimensionValue1": {
            "Type": "String"
        },
        "DimensionName2": {
            "Type": "String"
        },
        "DimensionValueSource2": {
            "Type": "String"
        },
        "DefaultDimensionValue2": {
            "Type": "String"
        }
    },
    "Resources": {
        "ConfigSet": {
            "Type": "AWS::SES::ConfigurationSet",
            "Properties": {
                "Name": {
                    "Ref": "ConfigSetName"
                }
            }
        },
        "CWEventDestination": {
            "Type": "AWS::SES::ConfigurationSetEventDestination",
            "Properties": {
                "ConfigurationSetName": {
                    "Ref": "ConfigSet"
                },
                "EventDestination": {
                    "Name": {
                        "Ref": "EventDestinationName"
                    },
                    "Enabled": true,
                    "MatchingEventTypes": [
                        {
                            "Ref": "EventType1"
                        },
                        {
                            "Ref": "EventType2"
                        },
                        {
                            "Ref": "EventType3"
                        }
                    ],
                    "CloudWatchDestination": {
                        "DimensionConfigurations": [
                            {
                                "DimensionName": {
                                    "Ref": "DimensionName1"
                                },
                                "DimensionValueSource": {
                                    "Ref": "DimensionValueSource1"
                                },
                                "DefaultDimensionValue": {
                                    "Ref": "DefaultDimensionValue1"
                                }
                            },
                            {
                                "DimensionName": {
                                    "Ref": "DimensionName2"
                                },
                                "DimensionValueSource": {
                                    "Ref": "DimensionValueSource2"
                                },
                                "DefaultDimensionValue": {
                                    "Ref": "DefaultDimensionValue2"
                                }
                            }
                        ]
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES ConfigurationSetEventDestination Sample Template
Parameters:
  ConfigSetName:
    Type: String
  EventDestinationName:
    Type: String
  EventType1:
    Type: String
  EventType2:
    Type: String
  EventType3:
    Type: String
  DimensionName1:
    Type: String
  DimensionValueSource1:
    Type: String
  DefaultDimensionValue1:
    Type: String
  DimensionName2:
    Type: String
  DimensionValueSource2:
    Type: String
  DefaultDimensionValue2:
    Type: String
Resources:
  ConfigSet:
    Type: 'AWS::SES::ConfigurationSet'
    Properties:
      Name: !Ref ConfigSetName
  CWEventDestination:
    Type: 'AWS::SES::ConfigurationSetEventDestination'
    Properties:
      ConfigurationSetName: !Ref ConfigSet
      EventDestination:
        Name: !Ref EventDestinationName
        Enabled: true
        MatchingEventTypes:
          - !Ref EventType1
          - !Ref EventType2
          - !Ref EventType3
        CloudWatchDestination:
          DimensionConfigurations:
            - DimensionName: !Ref DimensionName1
              DimensionValueSource: !Ref DimensionValueSource1
              DefaultDimensionValue: !Ref DefaultDimensionValue1
            - DimensionName: !Ref DimensionName2
              DimensionValueSource: !Ref DimensionValueSource2
              DefaultDimensionValue: !Ref DefaultDimensionValue2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VdmOptions

CloudWatchDestination

All content copied from https://docs.aws.amazon.com/.
