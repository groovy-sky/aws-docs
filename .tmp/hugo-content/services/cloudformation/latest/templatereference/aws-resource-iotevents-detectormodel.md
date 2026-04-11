This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel

The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a _detector_
_model_ (a model of your equipment or process) using _states_. For each
state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant
events. When an event is detected, it can change the state or trigger custom-built or predefined
actions using other AWS services. You can define additional events that trigger actions when entering
or exiting a state and, optionally, when a condition is met. For more information, see
[How to Use AWS IoT Events](../../../iotevents/latest/developerguide/how-to-use-iotevents.md) in the _AWS IoT Events Developer Guide_.

###### Note

When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI
commands, or CloudFormation) all detector instances created by the model are reset to their initial
states. (The detector's `state`, and the values of any variables and timers are reset.)

When you successfully update a detector model (using the AWS IoT Events console, AWS IoT Events API or CLI
commands, or CloudFormation) the version number of the detector model is incremented. (A detector model
with version number 1 before the update has version number 2 after the update succeeds.)

If you attempt to update a detector model using CloudFormation and the update does not succeed,
the system may, in some cases, restore the original detector model. When this occurs, the detector model's
version is incremented twice (for example, from version 1 to version 3) and the detector instances are
reset.

Also, be aware that if you attempt to update several detector models at once using CloudFormation,
some updates may succeed and others fail. In this case, the effects on each detector model's detector
instances and version number depend on whether the update succeeded or failed, with the results
as stated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTEvents::DetectorModel",
  "Properties" : {
      "DetectorModelDefinition" : DetectorModelDefinition,
      "DetectorModelDescription" : String,
      "DetectorModelName" : String,
      "EvaluationMethod" : String,
      "Key" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTEvents::DetectorModel
Properties:
  DetectorModelDefinition:
    DetectorModelDefinition
  DetectorModelDescription: String
  DetectorModelName: String
  EvaluationMethod: String
  Key: String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`DetectorModelDefinition`

Information that defines how a detector operates.

_Required_: Yes

_Type_: [DetectorModelDefinition](aws-properties-iotevents-detectormodel-detectormodeldefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorModelDescription`

A brief description of the detector model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorModelName`

The name of the detector model.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EvaluationMethod`

Information about the order in which events are evaluated and how actions are executed.

_Required_: No

_Type_: String

_Allowed values_: `BATCH | SERIAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The value used to identify a detector instance. When a device or system sends input, a new
detector instance with a unique key value is created. AWS IoT Events can continue to route input to its
corresponding detector instance based on this identifying information.

This parameter uses a JSON-path expression to select the attribute-value pair in the
message payload that is used for identification. To route the message to the correct detector
instance, the device must send a message payload that contains the same
attribute-value.

_Required_: No

_Type_: String

_Pattern_: ``^((`[\w\- ]+`)|([\w\-]+))(\.((`[\w\- ]+`)|([\w\-]+)))*$``

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the role that grants permission to AWS IoT Events to perform its operations.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-iotevents-detectormodel-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the detector model. For example:

`{"Ref": "myDetectorModel"}`

For the AWS IoT Events detector model `myDetectorModel`, `Ref` returns the name of the
detector model.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Simple Detector Model](#aws-resource-iotevents-detectormodel--examples--Simple_Detector_Model)

- [Full Detector Model](#aws-resource-iotevents-detectormodel--examples--Full_Detector_Model)

### Simple Detector Model

The following example creates a simple detector model with only one state.

#### JSON

```json

{
  "Description": "Simple Detector Model Template",
  "Resources": {
    "MyDetectorModel": {
      "Type": "AWS::IoTEvents::DetectorModel",
      "Properties": {
        "DetectorModelName": "myDetectorModel",
        "DetectorModelDescription": "My Detector Model created by CloudFormation",
        "Key": "myKey",
        "RoleArn": { "Fn::GetAtt" : [ "myRole", "Arn" ] },
        "DetectorModelDefinition": {
          "InitialStateName": "myInitialState",
          "States": [
            {
              "StateName": "myInitialState",
              "OnInput": {
                "Events": [
                  {
                    "EventName": "onInputPublishEvent",
                    "Condition": { "Fn::Join" : [ ".", ["$input", {"Ref": "myInput"}, "foo > 1"] ] },
                    "Actions": [
                      {
                        "IotTopicPublish": {
                          "MqttTopic": "myMqttTopic"
                        }
                      }
                    ]
                  }
                ]
              }
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "Simple Detector Model Template"
Resources:
  MyDetectorModel:
    Type: "AWS::IoTEvents::DetectorModel"
    Properties:
      DetectorModelName: "myDetectorModel"
      DetectorModelDescription: "My Detector Model created by CloudFormation"
      Key: "myKey"
      RoleArn: !GetAtt myRole.Arn
      DetectorModelDefinition:
        InitialStateName: "myInitialState"
        States:
         - StateName: "myInitialState"
           OnInput:
             Events:
               - EventName: "onInputPublishEvent"
                 Condition: !Join [".", ["$input", {'Ref': myInput}, "foo > 1"]]
                 Actions:
                   - IotTopicPublish:
                       MqttTopic: "myMqttTopic"

```

### Full Detector Model

The following example creates a more complete example of a detector model with two states.

#### JSON

```json

{
  "Description": "Detector Model Template for CloudFormation",
  "Resources": {
    "MyDetectorModel": {
      "Type": "AWS::IoTEvents::DetectorModel",
      "Properties": {
        "DetectorModelName": "myDetectorModel",
        "DetectorModelDescription": "My Detector Model created by CloudFormation",
        "Key": "myKey",
        "RoleArn": "arn:aws:iam::123456789012:role/myIotEventsRole",
        "DetectorModelDefinition": {
          "InitialStateName": "myInitialState",
          "States": [
            {
              "StateName": "myInitialState",
              "OnEnter": {
                "Events": [
                  {
                    "EventName": "onEnterEvent",
                    "Actions": [
                      {
                        "SetVariable": {
                          "VariableName": "Variable",
                          "Value": "0"
                        }
                      }
                    ]
                  },
                  {
                    "EventName": "onEnter Event 2",
                    "Condition": "true",
                    "Actions": [
                      {
                        "SetTimer": {
                          "TimerName": "myTimer",
                          "Seconds": 60
                        }
                      }
                    ]
                  }
                ]
              },
              "OnInput": {
                "Events": [
                  {
                    "EventName": "onInputEvent",
                    "Condition": { "Fn::Join" : [ ".", ["$input", {"Ref": "myInput"}, "foo > 1"] ] },
                    "Actions": [
                      {
                        "IotTopicPublish": {
                          "MqttTopic": "myMqttTopic"
                        }
                      },
                      {
                        "ResetTimer": {
                          "TimerName": "myTimer"
                        }
                      }
                    ]
                  }
                ],
                "TransitionEvents": [
                  {
                    "EventName": "Transit to other state",
                    "Condition": "true",
                    "Actions": [
                      {
                        "Sns": {
                          "TargetArn": "arn:aws:sns:123456789012:mySnsTopic"
                        }
                      },
                      {
                        "Lambda": {
                          "FunctionArn": "arn:aws:lambda:123456789012:function:myLambdaFunction"
                        }
                      },
                      {
                        "Firehose": {
                          "DeliveryStreamName": "myStreamName",
                          "Separator": ","
                        }
                      },
                      {
                        "Sqs": {
                          "QueueUrl": "myQueueUrl",
                          "UseBase64": true
                        }
                      },
                      {
                        "IotEvents": {
                          "InputName": "myInputName"
                        }
                      }
                    ],
                    "NextState": "myOtherState"
                  }
                ]
              },
              "OnExit": {
                "Events": [
                  {
                    "EventName": "Clear timers",
                    "Condition": "1 == 1",
                    "Actions": [
                      {
                        "ClearTimer": {
                          "TimerName": "myTimer"
                        }
                      }
                    ]
                  }
                ]
              }
            },
            {
              "StateName": "myOtherState",
              "OnEnter": {
                "Events": [
                  {
                    "EventName": "onEnterEvent",
                    "Actions": [
                      {
                        "SetVariable": {
                          "VariableName": "Variable",
                          "Value": "0"
                        }
                      }
                    ]
                  },
                  {
                    "EventName": "onEnter Event 2",
                    "Condition": "true",
                    "Actions": [
                      {
                        "SetTimer": {
                          "TimerName": "myTimer",
                          "Seconds": 60
                        }
                      }
                    ]
                  }
                ]
              },
              "OnExit": {
                "Events": [
                  {
                    "EventName": "Clear timers",
                    "Condition": "1 == 1",
                    "Actions": [
                      {
                        "ClearTimer": {
                          "TimerName": "myTimer"
                        }
                      }
                    ]
                  }
                ]
              }
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "Detector Model Template for CloudFormation"
Resources:
  MyDetectorModel:
    Type: "AWS::IoTEvents::DetectorModel"
    Properties:
      DetectorModelName: "myDetectorModel"
      DetectorModelDescription: "My Detector Model created by CloudFormation"
      Key: "myKey"
      RoleArn: "arn:aws:iam::123456789012:role/myIotEventsRole"
      DetectorModelDefinition:
        InitialStateName: "myInitialState"
        States:
          - StateName: "myInitialState"
            OnEnter:
              Events:
                - EventName: "onEnterEvent"
                  Actions:
                    - SetVariable:
                        VariableName: "Variable"
                        Value: "0"
                - EventName: "onEnter Event 2"
                  Condition: "true"
                  Actions:
                    - SetTimer:
                        TimerName: "myTimer"
                        Seconds: 60
            OnInput:
              Events:
                - EventName: "onInputEvent"
                  Condition: !Join [".", ["$input", {'Ref': myInput}, "foo > 1"]]
                  Actions:
                    - IotTopicPublish:
                        MqttTopic: "myMqttTopic"
                    - ResetTimer:
                        TimerName: "myTimer"
              TransitionEvents:
                - EventName: "Transit to other state"
                  Condition: "true"
                  Actions:
                    - Sns:
                        TargetArn: "arn:aws:sns:123456789012:mySnsTopic"
                    - Lambda:
                        FunctionArn: "arn:aws:lambda:123456789012:function:myLambdaFunction"
                    - Firehose:
                        DeliveryStreamName: "myStreamName"
                        Separator: ","
                    - Sqs:
                        QueueUrl: "myQueueUrl"
                        UseBase64: true
                    - IotEvents:
                        InputName: "myInputName"
                  NextState: "myOtherState"
            OnExit:
              Events:
                - EventName: "Clear timers"
                  Condition: "1 == 1"
                  Actions:
                    - ClearTimer:
                        TimerName: "myTimer"
          - StateName: "myOtherState"
            OnEnter:
              Events:
                - EventName: "onEnterEvent"
                  Actions:
                    - SetVariable:
                        VariableName: "Variable"
                        Value: "0"
                - EventName: "onEnter Event 2"
                  Condition: "true"
                  Actions:
                    - SetTimer:
                        TimerName: "myTimer"
                        Seconds: 60
            OnExit:
              Events:
                - EventName: "Clear timers"
                  Condition: "1 == 1"
                  Actions:
                    - ClearTimer:
                        TimerName: "myTimer"

```

## See also

- [How to Use AWS IoT Events](../../../iotevents/latest/developerguide/how-to-use-iotevents.md) in the _AWS IoT Events Developer Guide_

- [CreateDetectorModel](../../../../reference/iotevents/latest/apireference/api-createdetectormodel.md) in the _AWS IoT Events API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Action

All content copied from https://docs.aws.amazon.com/.
