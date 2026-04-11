This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective RollingInterval

If the interval for this SLO is a rolling interval, this structure contains the interval specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Duration" : Integer,
  "DurationUnit" : String
}

```

### YAML

```yaml

  Duration: Integer
  DurationUnit: String

```

## Properties

`Duration`

Specifies the duration of each rolling interval. For example, if `Duration` is `7` and
`DurationUnit` is `DAY`, each rolling interval is seven days.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DurationUnit`

Specifies the rolling interval unit.

_Required_: Yes

_Type_: String

_Allowed values_: `MINUTE | HOUR | DAY | MONTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestBasedSliMetric

Sli

All content copied from https://docs.aws.amazon.com/.
