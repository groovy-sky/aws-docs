This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesThinClient::Environment MaintenanceWindow

###### Important

End of support notice: On March 31, 2027, AWS will end support for Amazon WorkSpaces Thin Client. After March 31, 2027,
you will no longer be able to access the WorkSpaces Thin Client console or WorkSpaces Thin Client resources. For more information, see
[Amazon WorkSpaces Thin Client end of support](https://docs.aws.amazon.com/workspaces-thin-client/latest/ug/workspacesthinclient-end-of-support.html).

Describes the maintenance window for a thin client device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplyTimeOf" : String,
  "DaysOfTheWeek" : [ String, ... ],
  "EndTimeHour" : Integer,
  "EndTimeMinute" : Integer,
  "StartTimeHour" : Integer,
  "StartTimeMinute" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  ApplyTimeOf: String
  DaysOfTheWeek:
    - String
  EndTimeHour: Integer
  EndTimeMinute: Integer
  StartTimeHour: Integer
  StartTimeMinute: Integer
  Type: String

```

## Properties

`ApplyTimeOf`

The option to set the maintenance window during the device local time or Universal
Coordinated Time (UTC).

_Required_: No

_Type_: String

_Allowed values_: `UTC | DEVICE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DaysOfTheWeek`

The days of the week during which the maintenance window is open.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTimeHour`

The hour for the maintenance window end ( `00`- `23`).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTimeMinute`

The minutes for the maintenance window end ( `00`- `59`).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `59`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTimeHour`

The hour for the maintenance window start ( `00`- `23`).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTimeMinute`

The minutes past the hour for the maintenance window start
( `00`- `59`).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `59`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

An option to select the default or custom maintenance window.

_Required_: Yes

_Type_: String

_Allowed values_: `SYSTEM | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::WorkSpacesThinClient::Environment

Tag
