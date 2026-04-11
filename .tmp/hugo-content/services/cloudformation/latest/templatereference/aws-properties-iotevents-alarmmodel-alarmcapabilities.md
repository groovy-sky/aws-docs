This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AlarmCapabilities

Contains the configuration information of alarm state changes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcknowledgeFlow" : AcknowledgeFlow,
  "InitializationConfiguration" : InitializationConfiguration
}

```

### YAML

```yaml

  AcknowledgeFlow:
    AcknowledgeFlow
  InitializationConfiguration:
    InitializationConfiguration

```

## Properties

`AcknowledgeFlow`

Specifies whether to get notified for alarm state changes.

_Required_: No

_Type_: [AcknowledgeFlow](aws-properties-iotevents-alarmmodel-acknowledgeflow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitializationConfiguration`

Specifies the default alarm state.
The configuration applies to all alarms that were created based on this alarm model.

_Required_: No

_Type_: [InitializationConfiguration](aws-properties-iotevents-alarmmodel-initializationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmAction

AlarmEventActions

All content copied from https://docs.aws.amazon.com/.
