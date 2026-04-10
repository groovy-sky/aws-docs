This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask NotificationConfig

The `NotificationConfig` property type specifies configurations for sending
notifications for a maintenance window task in AWS Systems Manager.

`NotificationConfig` is a property of the [MaintenanceWindowRunCommandParameters](../userguide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NotificationArn" : String,
  "NotificationEvents" : [ String, ... ],
  "NotificationType" : String
}

```

### YAML

```yaml

  NotificationArn: String
  NotificationEvents:
    - String
  NotificationType: String

```

## Properties

`NotificationArn`

An Amazon Resource Name (ARN) for an Amazon Simple Notification Service (Amazon SNS) topic. Run
Command pushes notifications about command status changes to this topic.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationEvents`

The different events that you can receive notifications for. These events include the
following: `All` (events), `InProgress`, `Success`,
`TimedOut`, `Cancelled`, `Failed`. To learn more
about these events, see [Configuring\
Amazon SNS Notifications for AWS Systems Manager](../../../systems-manager/latest/userguide/monitoring-sns-notifications.md) in the
_AWS Systems Manager User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationType`

The notification type.

- `Command`: Receive notification when the status of a command
changes.

- `Invocation`: For commands sent to multiple instances, receive
notification on a per-instance basis when the status of a command
changes.

_Required_: No

_Type_: String

_Allowed values_: `Command | Invocation`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceWindowStepFunctionsParameters

Target

All content copied from https://docs.aws.amazon.com/.
