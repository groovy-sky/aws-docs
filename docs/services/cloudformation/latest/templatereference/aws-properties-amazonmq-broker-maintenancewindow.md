This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker MaintenanceWindow

The parameters that determine the WeeklyStartTime.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DayOfWeek" : String,
  "TimeOfDay" : String,
  "TimeZone" : String
}

```

### YAML

```yaml

  DayOfWeek: String
  TimeOfDay: String
  TimeZone: String

```

## Properties

`DayOfWeek`

Required. The day of the week.

_Required_: Yes

_Type_: String

_Allowed values_: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeOfDay`

Required. The time, in 24-hour format.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{1,2}:[0-9]{2}(:[0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The time zone, UTC by default, in either the Country/City format, or the UTC
offset format.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogList

TagsEntry

All content copied from https://docs.aws.amazon.com/.
