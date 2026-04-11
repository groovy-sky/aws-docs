This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindow

The `AWS::SSM::MaintenanceWindow` resource represents general information
about a maintenance window for AWS Systems Manager. Maintenance windows let you
define a schedule for when to perform potentially disruptive actions on your instances,
such as patching an operating system (OS), updating drivers, or installing software.
Each maintenance window has a schedule, a duration, a set of registered targets, and a
set of registered tasks.

For more information, see [Systems Manager Maintenance Windows](../../../systems-manager/latest/userguide/systems-manager-maintenance.md) in the _AWS Systems Manager User Guide_ and [CreateMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-createmaintenancewindow.md) in the _AWS Systems Manager API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::MaintenanceWindow",
  "Properties" : {
      "AllowUnassociatedTargets" : Boolean,
      "Cutoff" : Integer,
      "Description" : String,
      "Duration" : Integer,
      "EndDate" : String,
      "Name" : String,
      "Schedule" : String,
      "ScheduleOffset" : Integer,
      "ScheduleTimezone" : String,
      "StartDate" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSM::MaintenanceWindow
Properties:
  AllowUnassociatedTargets: Boolean
  Cutoff: Integer
  Description: String
  Duration: Integer
  EndDate: String
  Name: String
  Schedule: String
  ScheduleOffset: Integer
  ScheduleTimezone: String
  StartDate: String
  Tags:
    - Tag

```

## Properties

`AllowUnassociatedTargets`

Enables a maintenance window task to run on managed instances, even if you have not
registered those instances as targets. If enabled, then you must specify the
unregistered instances (by instance ID) when you register a task with the maintenance
window.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cutoff`

The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling
new tasks for execution.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the maintenance window.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Duration`

The duration of the maintenance window in hours.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndDate`

The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled
to become inactive.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the maintenance window.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule of the maintenance window in the form of a cron or rate expression.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleOffset`

The number of days to wait to run a maintenance window after the scheduled cron expression
date and time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleTimezone`

The time zone that the scheduled maintenance window executions are based on, in Internet
Assigned Numbers Authority (IANA) format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartDate`

The date and time, in ISO-8601 Extended format, for when the maintenance window is
scheduled to become active. `StartDate` allows you to delay activation of the
maintenance window until the specified future date.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you assign to a resource in the form of an arbitrary set of
tags (key-value pairs). Tags enable you to categorize a resource in different ways, such
as by purpose, owner, or environment. For example, you might want to tag a maintenance
window to identify the type of tasks it will run, the types of targets, and the
environment it will run in.

_Required_: No

_Type_: Array of [Tag](aws-properties-ssm-maintenancewindow-tag.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the maintenance window ID, such as
`mw-abcde1234567890yz`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`WindowId`

The ID of the maintenance window.

## Examples

### Create a maintenance window that does not allow unregistered targets

The following example creates a Systems Manager maintenance window that
runs for two hours with a one hour cutoff every Sunday at 04:00 AM US Eastern
Time. The maintenance window doesn't allow unregistered targets.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindow": {
            "Type": "AWS::SSM::MaintenanceWindow",
            "Properties": {
                "AllowUnassociatedTargets": false,
                "Cutoff": 1,
                "Description": "Maintenance Window to update SSM Agent",
                "Duration": 2,
                "Name": "UpdateSSMAgentMaintenanceWindow",
                "Schedule": "cron(0 4 ? * SUN *)",
                "ScheduleTimezone": "US/Eastern"
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindow:
    Type: AWS::SSM::MaintenanceWindow
    Properties:
      AllowUnassociatedTargets: false
      Cutoff: 1
      Description: Maintenance Window to update SSM Agent
      Duration: 2
      Name: UpdateSSMAgentMaintenanceWindow
      Schedule: cron(0 4 ? * SUN *)
      ScheduleTimezone: US/Eastern
```

## See also

- [AWS::SSM::MaintenanceWindowTarget](../userguide/aws-resource-ssm-maintenancewindowtarget.md)

- [AWS::SSM::MaintenanceWindowTask](../userguide/aws-resource-ssm-maintenancewindowtask.md)

- [CreateMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-createmaintenancewindow.md) in the _AWS Systems Manager API Reference._

- [Reference: Cron and Rate Expressions for Systems Manager](../../../systems-manager/latest/userguide/reference-cron-and-rate-expressions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
