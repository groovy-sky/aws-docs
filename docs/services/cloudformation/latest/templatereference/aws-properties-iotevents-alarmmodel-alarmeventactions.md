This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AlarmEventActions

Contains information about one or more alarm actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmActions" : [ AlarmAction, ... ]
}

```

### YAML

```yaml

  AlarmActions:
    - AlarmAction

```

## Properties

`AlarmActions`

Specifies one or more supported actions to receive notifications when the alarm state
changes.

_Required_: No

_Type_: Array of [AlarmAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotevents-alarmmodel-alarmaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AlarmCapabilities

AlarmRule
