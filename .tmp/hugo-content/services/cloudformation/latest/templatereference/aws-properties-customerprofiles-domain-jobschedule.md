This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain JobSchedule

The day and time when do you want to start the Identity Resolution Job every
week.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DayOfTheWeek" : String,
  "Time" : String
}

```

### YAML

```yaml

  DayOfTheWeek: String
  Time: String

```

## Properties

`DayOfTheWeek`

The day when the Identity Resolution Job should run every week.

_Required_: Yes

_Type_: String

_Allowed values_: `SUNDAY | MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time when the Identity Resolution Job should run every week.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`

_Minimum_: `3`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportingConfig

Matching

All content copied from https://docs.aws.amazon.com/.
