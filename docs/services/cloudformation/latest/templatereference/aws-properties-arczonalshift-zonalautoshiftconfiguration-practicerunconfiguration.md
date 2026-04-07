This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCZonalShift::ZonalAutoshiftConfiguration PracticeRunConfiguration

A practice run configuration for a resource includes the Amazon CloudWatch alarms that you've specified for a practice
run, as well as any blocked dates or blocked windows for the practice run.

When a resource has a practice run configuation, ARC
starts weekly zonal shifts for the resource, to shift traffic away from an Availability Zone. Weekly practice
runs help you to make sure that your application can continue to operate normally with the loss of one
Availability Zone.

You can update or delete a practice run configuration. When you delete a practice run configuration, zonal
autoshift is disabled for the resource. A practice run configuration is required when zonal autoshift is enabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockedDates" : [ String, ... ],
  "BlockedWindows" : [ String, ... ],
  "BlockingAlarms" : [ ControlCondition, ... ],
  "OutcomeAlarms" : [ ControlCondition, ... ]
}

```

### YAML

```yaml

  BlockedDates:
    - String
  BlockedWindows:
    - String
  BlockingAlarms:
    - ControlCondition
  OutcomeAlarms:
    - ControlCondition

```

## Properties

`BlockedDates`

An array of one or more dates that you can specify when AWS does not start practice runs for a resource.
Dates are in UTC.

Specify blocked dates in the format `YYYY-MM-DD`, separated by spaces.

_Required_: No

_Type_: Array of String

_Minimum_: `10 | 0`

_Maximum_: `10 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedWindows`

An array of one or more days and times that you can specify when ARC does not start practice runs for a resource. Days
and times are in UTC.

Specify blocked windows in the format `DAY:HH:MM-DAY:HH:MM`, separated by spaces. For example,
`MON:18:30-MON:19:30 TUE:18:30-TUE:19:30`.

###### Important

Blocked windows have to start and end on the same day. Windows that span multiple days aren't supported.

_Required_: No

_Type_: Array of String

_Minimum_: `19 | 0`

_Maximum_: `19 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockingAlarms`

An optional alarm that you can specify that blocks practice runs when the alarm is in an `ALARM` state.
When a blocking alarm goes into an `ALARM` state, it prevents practice runs from being started, and ends practice runs that are in progress.

_Required_: No

_Type_: Array of [ControlCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-arczonalshift-zonalautoshiftconfiguration-controlcondition.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutcomeAlarms`

The alarm that you specify to monitor the health of your application during practice runs. When the outcome alarm
goes into an `ALARM` state, the practice run is ended and the outcome is set to `FAILED`.

_Required_: Yes

_Type_: Array of [ControlCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-arczonalshift-zonalautoshiftconfiguration-controlcondition.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ControlCondition

Next
