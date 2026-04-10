This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective SliMetric

Use this structure to specify the metric to be used for the SLO.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependencyConfig" : DependencyConfig,
  "KeyAttributes" : String,
  "MetricDataQueries" : [ MetricDataQuery, ... ],
  "MetricType" : String,
  "OperationName" : String,
  "PeriodSeconds" : Integer,
  "Statistic" : String
}

```

### YAML

```yaml

  DependencyConfig:
    DependencyConfig
  KeyAttributes: String
  MetricDataQueries:
    - MetricDataQuery
  MetricType: String
  OperationName: String
  PeriodSeconds: Integer
  Statistic: String

```

## Properties

`DependencyConfig`

Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName`.

_Required_: No

_Type_: [DependencyConfig](aws-properties-applicationsignals-servicelevelobjective-dependencyconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyAttributes`

If this SLO is related to a metric collected by Application Signals, you must use this field to specify which service
the SLO metric is related to. To do so, you must specify at least the `Type`,
`Name`, and `Environment` attributes.

This is a string-to-string map. It can
include the following fields.

- `Type` designates the type of object this is.

- `ResourceType` specifies the type of the resource. This field is used only
when the value of the `Type` field is `Resource` or `AWS::Resource`.

- `Name` specifies the name of the object. This is used only if the value of the `Type` field
is `Service`, `RemoteService`, or `AWS::Service`.

- `Identifier` identifies the resource objects of this resource.
This is used only if the value of the `Type` field
is `Resource` or `AWS::Resource`.

- `Environment` specifies the location where this object is hosted, or what it belongs to.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricDataQueries`

If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, use this structure to specify that metric or expression.

_Required_: No

_Type_: Array of [MetricDataQuery](aws-properties-applicationsignals-servicelevelobjective-metricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricType`

If the SLO is to monitor either the `LATENCY` or `AVAILABILITY` metric that Application Signals collects, use this field to specify which of those metrics is used.

_Required_: No

_Type_: String

_Allowed values_: `LATENCY | AVAILABILITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperationName`

If the SLO is to monitor a specific operation of the service, use this field to specify the name of that operation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodSeconds`

The number of seconds to use as the period for SLO evaluation. Your application's performance is compared to the SLI during each period. For each period, the application is determined to have either achieved or not achieved the necessary performance.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.
For more information about statistics,
see [CloudWatch statistics definitions](../../../amazoncloudwatch/latest/monitoring/statistics-definitions.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sli

Tag

All content copied from https://docs.aws.amazon.com/.
