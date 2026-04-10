This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel

Represents an alarm model to monitor an AWS IoT Events input attribute. You can use the alarm to get
notified when the value is outside a specified range. For more information, see [Create an\
alarm model](../../../iotevents/latest/developerguide/create-alarms.md) in the _AWS IoT Events Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTEvents::AlarmModel",
  "Properties" : {
      "AlarmCapabilities" : AlarmCapabilities,
      "AlarmEventActions" : AlarmEventActions,
      "AlarmModelDescription" : String,
      "AlarmModelName" : String,
      "AlarmRule" : AlarmRule,
      "Key" : String,
      "RoleArn" : String,
      "Severity" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTEvents::AlarmModel
Properties:
  AlarmCapabilities:
    AlarmCapabilities
  AlarmEventActions:
    AlarmEventActions
  AlarmModelDescription: String
  AlarmModelName: String
  AlarmRule:
    AlarmRule
  Key: String
  RoleArn: String
  Severity: Integer
  Tags:
    - Tag

```

## Properties

`AlarmCapabilities`

Contains the configuration information of alarm state changes.

_Required_: No

_Type_: [AlarmCapabilities](aws-properties-iotevents-alarmmodel-alarmcapabilities.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmEventActions`

Contains information about one or more alarm actions.

_Required_: No

_Type_: [AlarmEventActions](aws-properties-iotevents-alarmmodel-alarmeventactions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmModelDescription`

The description of the alarm model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmModelName`

The name of the alarm model.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AlarmRule`

Defines when your alarm is invoked.

_Required_: Yes

_Type_: [AlarmRule](aws-properties-iotevents-alarmmodel-alarmrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

An input attribute used as a key to create an alarm.
AWS IoT Events routes [inputs](../../../../reference/iotevents/latest/apireference/api-input.md)
associated with this key to the alarm.

_Required_: No

_Type_: String

_Pattern_: ``^((`[\w\- ]+`)|([\w\-]+))(\.((`[\w\- ]+`)|([\w\-]+)))*$``

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM role that allows the alarm to perform actions and access AWS resources. For more information, see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the _AWS General Reference_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Severity`

A non-negative integer that reflects the severity level of the alarm.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the alarm model. The tags help you
manage the alarm model. For more information, see [Tagging your AWS IoT Events\
resources](../../../iotevents/latest/developerguide/tagging-iotevents.md) in the _AWS IoT Events Developer Guide_.

You can create up to 50 tags for one alarm model.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotevents-alarmmodel-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the alarm model. For example:

`{"Ref": "myAlarmModel"}`

For the AWS IoT Events alarm model `myAlarmModel`, `Ref` returns the name of the
alarm model.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT Events

AcknowledgeFlow

All content copied from https://docs.aws.amazon.com/.
