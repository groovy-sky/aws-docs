This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration Time

The time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimeOfDay" : String,
  "TimeZone" : String
}

```

### YAML

```yaml

  TimeOfDay: String
  TimeZone: String

```

## Properties

`TimeOfDay`

The time of day in 24-hour format (00:00).

_Required_: Yes

_Type_: String

_Pattern_: `^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The timezone.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

WeeklySchedule

All content copied from https://docs.aws.amazon.com/.
