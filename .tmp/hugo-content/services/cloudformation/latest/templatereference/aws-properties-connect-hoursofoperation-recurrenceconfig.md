This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation RecurrenceConfig

Defines the recurrence configuration for overrides. This configuration uses a recurrence pattern to specify when and how frequently an event should repeat.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecurrencePattern" : RecurrencePattern
}

```

### YAML

```yaml

  RecurrencePattern:
    RecurrencePattern

```

## Properties

`RecurrencePattern`

The recurrence pattern that defines how the event repeats. Example: Frequency, Interval, ByMonth, ByMonthDay, ByWeekdayOccurrence

_Required_: Yes

_Type_: [RecurrencePattern](aws-properties-connect-hoursofoperation-recurrencepattern.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverrideTimeSlice

RecurrencePattern

All content copied from https://docs.aws.amazon.com/.
