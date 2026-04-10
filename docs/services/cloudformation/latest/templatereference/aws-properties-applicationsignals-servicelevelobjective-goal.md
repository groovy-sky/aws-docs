This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective Goal

This structure contains the attributes that determine the goal of an SLO. This includes
the time period for evaluation and the attainment threshold.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttainmentGoal" : Number,
  "Interval" : Interval,
  "WarningThreshold" : Number
}

```

### YAML

```yaml

  AttainmentGoal: Number
  Interval:
    Interval
  WarningThreshold: Number

```

## Properties

`AttainmentGoal`

The threshold that determines if the goal is being met.

If this is a period-based SLO, the attainment goal is the
percentage of good periods that meet the threshold requirements to the total periods within the interval.
For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the
periods to be in healthy state.

If this is a request-based SLO, the attainment goal is the percentage of requests that must be
successful to meet the attainment goal.

If you omit this parameter, 99 is used
to represent 99% as the attainment goal.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The time period used to evaluate the SLO. It can be either a calendar interval or rolling interval.

If you omit this parameter, a rolling interval of 7 days is used.

_Required_: No

_Type_: [Interval](aws-properties-applicationsignals-servicelevelobjective-interval.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarningThreshold`

The percentage of remaining budget over total budget that you want to get warnings for.
If you omit this parameter, the default of 50.0 is used.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExclusionWindow

Interval

All content copied from https://docs.aws.amazon.com/.
