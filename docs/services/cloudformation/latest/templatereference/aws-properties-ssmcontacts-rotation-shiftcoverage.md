This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Rotation ShiftCoverage

Information about the days of the week that the on-call rotation coverage
includes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CoverageTimes" : [ CoverageTime, ... ],
  "DayOfWeek" : String
}

```

### YAML

```yaml

  CoverageTimes:
    - CoverageTime
  DayOfWeek: String

```

## Properties

`CoverageTimes`

The start and end times of the shift.

_Required_: Yes

_Type_: Array of [CoverageTime](aws-properties-ssmcontacts-rotation-coveragetime.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DayOfWeek`

A list of days on which the schedule is active.

_Required_: Yes

_Type_: String

_Allowed values_: `MON | TUE | WED | THU | FRI | SAT | SUN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecurrenceSettings

Tag

All content copied from https://docs.aws.amazon.com/.
