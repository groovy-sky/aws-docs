This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective CalendarInterval

If the interval for this service level objective is a calendar interval, this structure contains the interval specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Duration" : Integer,
  "DurationUnit" : String,
  "StartTime" : Integer
}

```

### YAML

```yaml

  Duration: Integer
  DurationUnit: String
  StartTime: Integer

```

## Properties

`Duration`

Specifies the duration of each calendar interval. For example, if `Duration` is `1` and
`DurationUnit` is `MONTH`, each interval is one month, aligned with the calendar.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DurationUnit`

Specifies the calendar interval unit.

_Required_: Yes

_Type_: String

_Allowed values_: `MINUTE | HOUR | DAY | MONTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The date and time when you want the first interval to start. Be sure to choose a time that configures the
intervals the way that you want. For example, if you want weekly intervals
starting on Mondays at 6 a.m., be sure to specify a start time that is a Monday at 6 a.m.

When used in a raw HTTP Query API, it is formatted as
be epoch time in seconds. For example: `1698778057`

As soon as one calendar interval ends, another automatically begins.

_Required_: Yes

_Type_: Integer

_Minimum_: `946684800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BurnRateConfiguration

DependencyConfig

All content copied from https://docs.aws.amazon.com/.
