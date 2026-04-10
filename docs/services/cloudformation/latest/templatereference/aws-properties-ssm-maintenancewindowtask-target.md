This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask Target

The `Target` property type specifies targets (either instances or window
target IDs). You specify instances by using `Key=InstanceIds,Values=<instanceid1>,<instanceid2>`. You specify window target IDs using
`Key=WindowTargetIds,Values=<window-target-id-1>,<window-target-id-2>` for a maintenance window task in
AWS Systems Manager.

`Target` is a property of the [AWS::SSM::MaintenanceWindowTask](../userguide/aws-resource-ssm-maintenancewindowtask.md) property type.

###### Note

To use `resource-groups:Name` as the key for a maintenance window
target, specify the resource group as a
`AWS::SSM::MaintenanceWindowTarget` type, and use the
`Ref` function to specify the target for
`AWS::SSM::MaintenanceWindowTask`. For an example, see **Create a Run Command task that targets instances using a resource**
**group name** in [AWS::SSM::MaintenanceWindowTask Examples](../userguide/aws-resource-ssm-maintenancewindowtask.md#aws-resource-ssm-maintenancewindowtask--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

User-defined criteria for sending commands that target instances that meet the
criteria. `Key` can be `InstanceIds` or
`WindowTargetIds`. For more information about how to target instances
within a maintenance window task, see [About\
'register-task-with-maintenance-window' Options and Values](../../../systems-manager/latest/userguide/register-tasks-options.md) in the
_AWS Systems Manager User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}_.:/=\-@]*$|resource-groups:ResourceTypeFilters|resource-groups:Name`

_Minimum_: `1`

_Maximum_: `163`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

User-defined criteria that maps to `Key`. For example, if you specify
`InstanceIds`, you can specify
`i-1234567890abcdef0,i-9876543210abcdef0` to run a command on two EC2
instances. For more information about how to target instances within a maintenance
window task, see [About\
'register-task-with-maintenance-window' Options and Values](../../../systems-manager/latest/userguide/register-tasks-options.md) in the
_AWS Systems Manager User Guide_.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [RegisterTargetWithMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-registertargetwithmaintenancewindow.md) in the _AWS Systems Manager API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationConfig

TaskInvocationParameters

All content copied from https://docs.aws.amazon.com/.
