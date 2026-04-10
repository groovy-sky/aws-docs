This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective MetricDataQuery

Use this structure to define a metric or metric math expression that you want to use as for a service level objective.

Each `MetricDataQuery` in the `MetricDataQueries` array specifies either a metric to retrieve, or a metric math expression
to be performed on retrieved metrics. A single `MetricDataQueries` array can include as many as 20 `MetricDataQuery` structures in the array.
The 20 structures can include as many as 10 structures that contain a `MetricStat` parameter to retrieve a metric, and as many as 10 structures that
contain the `Expression` parameter to perform a math expression. Of those `Expression` structures,
exactly one must have true as the value for `ReturnData`. The result of this expression used for the SLO.

For more information about metric math expressions, see
[Use metric math](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md).

Within each `MetricDataQuery` object, you must specify either
`Expression` or `MetricStat` but not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "Expression" : String,
  "Id" : String,
  "MetricStat" : MetricStat,
  "ReturnData" : Boolean
}

```

### YAML

```yaml

  AccountId: String
  Expression: String
  Id: String
  MetricStat:
    MetricStat
  ReturnData: Boolean

```

## Properties

`AccountId`

The ID of the account where this metric is located. If you are performing this operation in a monitoring account,
use this to specify which source account to retrieve this metric from.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

This field can contain a metric math expression to be performed on the other metrics that
you are retrieving within this `MetricDataQueries` structure.

A math expression
can use the `Id` of the other metrics or queries to refer to those metrics, and can also use
the `Id` of other
expressions to use the result of those expressions. For more information about metric math expressions, see
[Metric Math Syntax and Functions](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md#metric-math-syntax) in the
_Amazon CloudWatch User Guide_.

Within each `MetricDataQuery` object, you must specify either
`Expression` or `MetricStat` but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A short name used to tie this object to the results in the response. This `Id` must be unique
within a `MetricDataQueries` array. If you are performing math expressions on this set of data,
this name represents that data and can serve as a variable in the metric math expression. The valid characters
are letters, numbers, and underscore. The first character must be a lowercase letter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricStat`

A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO.

Within one `MetricDataQuery` object, you must specify either
`Expression` or `MetricStat` but not both.

_Required_: No

_Type_: [MetricStat](aws-properties-applicationsignals-servicelevelobjective-metricstat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReturnData`

Use this only if you are using a metric math expression for the SLO.
Specify `true` for `ReturnData` for only the one expression result to use as the alarm. For all
other metrics and expressions in the same `CreateServiceLevelObjective` operation, specify `ReturnData` as `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Metric

MetricStat

All content copied from https://docs.aws.amazon.com/.
