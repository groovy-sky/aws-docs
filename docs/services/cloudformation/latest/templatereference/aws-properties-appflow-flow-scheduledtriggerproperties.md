This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow ScheduledTriggerProperties

Specifies the configuration details of a schedule-triggered flow as defined by the user.
Currently, these settings only apply to the `Scheduled` trigger type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPullMode" : String,
  "FirstExecutionFrom" : Number,
  "FlowErrorDeactivationThreshold" : Integer,
  "ScheduleEndTime" : Number,
  "ScheduleExpression" : String,
  "ScheduleOffset" : Number,
  "ScheduleStartTime" : Number,
  "TimeZone" : String
}

```

### YAML

```yaml

  DataPullMode: String
  FirstExecutionFrom: Number
  FlowErrorDeactivationThreshold: Integer
  ScheduleEndTime: Number
  ScheduleExpression: String
  ScheduleOffset: Number
  ScheduleStartTime: Number
  TimeZone: String

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

Specifies the date range for the records to import from the connector in the first flow
run.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowErrorDeactivationThreshold`

Defines how many times a scheduled flow fails consecutively before Amazon AppFlow
deactivates it.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleEndTime`

The time at which the scheduled flow ends. The time is formatted as a timestamp that
follows the ISO 8601 standard, such as `2022-04-27T13:00:00-07:00`.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

The scheduling expression that determines the rate at which the schedule will run, for
example `rate(5minutes)`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleOffset`

Specifies the optional offset that is added to the time interval for a schedule-triggered
flow.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `36000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleStartTime`

The time at which the scheduled flow starts. The time is formatted as a timestamp that
follows the ISO 8601 standard, such as `2022-04-26T13:00:00-07:00`.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

Specifies the time zone used when referring to the dates and times of a scheduled flow,
such as `America/New_York`. This time zone is only a descriptive label. It doesn't
affect how Amazon AppFlow interprets the timestamps that you specify to schedule the
flow.

If you want to schedule a flow by using times in a particular time zone, indicate the time
zone as a UTC offset in your timestamps. For example, the UTC offsets for the
`America/New_York` timezone are `-04:00` EDT and `-05:00
        EST`.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ScheduledTriggerProperties](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ScheduledTriggerProperties.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SAPODataSourceProperties

ServiceNowSourceProperties
