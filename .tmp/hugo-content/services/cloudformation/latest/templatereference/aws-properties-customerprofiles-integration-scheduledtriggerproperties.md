This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration ScheduledTriggerProperties

Specifies the configuration details of a scheduled-trigger flow that you define.
Currently, these settings only apply to the scheduled-trigger type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPullMode" : String,
  "FirstExecutionFrom" : Number,
  "ScheduleEndTime" : Number,
  "ScheduleExpression" : String,
  "ScheduleOffset" : Integer,
  "ScheduleStartTime" : Number,
  "Timezone" : String
}

```

### YAML

```yaml

  DataPullMode: String
  FirstExecutionFrom: Number
  ScheduleEndTime: Number
  ScheduleExpression: String
  ScheduleOffset: Integer
  ScheduleStartTime: Number
  Timezone: String

```

## Properties

`DataPullMode`

Specifies whether a scheduled flow has an incremental data transfer or a complete data
transfer for each flow run.

_Required_: No

_Type_: String

_Allowed values_: `Incremental | Complete`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstExecutionFrom`

Specifies the date range for the records to import from the connector in the first
flow run.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleEndTime`

Specifies the scheduled end time for a scheduled-trigger flow.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

The scheduling expression that determines the rate at which the schedule will run, for
example rate (5 minutes).

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleOffset`

Specifies the optional offset that is added to the time interval for a
schedule-triggered flow.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `36000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleStartTime`

Specifies the scheduled start time for a scheduled-trigger flow. The value must be a
date/time value in EPOCH format.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

Specifies the time zone used when referring to the date and time of a
scheduled-triggered flow, such as America/New\_York.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceSourceProperties

ServiceNowSourceProperties

All content copied from https://docs.aws.amazon.com/.
