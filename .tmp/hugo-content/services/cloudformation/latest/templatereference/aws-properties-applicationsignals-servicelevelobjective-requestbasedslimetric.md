This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective RequestBasedSliMetric

This structure contains the information about the metric that is used for a request-based SLO.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependencyConfig" : DependencyConfig,
  "KeyAttributes" : String,
  "MetricType" : String,
  "MonitoredRequestCountMetric" : MonitoredRequestCountMetric,
  "OperationName" : String,
  "TotalRequestCountMetric" : [ MetricDataQuery, ... ]
}

```

### YAML

```yaml

  DependencyConfig:
    DependencyConfig
  KeyAttributes: String
  MetricType: String
  MonitoredRequestCountMetric:
    MonitoredRequestCountMetric
  OperationName: String
  TotalRequestCountMetric:
    - MetricDataQuery

```

## Properties

`DependencyConfig`

Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName`.

_Required_: No

_Type_: [DependencyConfig](aws-properties-applicationsignals-servicelevelobjective-dependencyconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyAttributes`

This is a string-to-string map that contains information about the type of object that this SLO is related to. It can
include the following fields.

- `Type` designates the type of object that this SLO is related to.

- `ResourceType` specifies the type of the resource. This field is used only
when the value of the `Type` field is `Resource` or `AWS::Resource`.

- `Name` specifies the name of the object. This is used only if the value of the `Type` field
is `Service`, `RemoteService`, or `AWS::Service`.

- `Identifier` identifies the resource objects of this resource.
This is used only if the value of the `Type` field
is `Resource` or `AWS::Resource`.

- `Environment` specifies the location where this object is hosted, or what it belongs to.

- `AwsAccountId` allows you to create an SLO for an object that exists in another account.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricType`

If the SLO monitors either the `LATENCY` or `AVAILABILITY` metric that Application Signals
collects, this field displays which of those metrics is used.

_Required_: No

_Type_: String

_Allowed values_: `LATENCY | AVAILABILITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoredRequestCountMetric`

Use this structure to define the metric that you want to use as the "good request" or "bad request"
value for a request-based SLO.
This value observed for the metric defined in
`TotalRequestCountMetric` will be divided by the number found for
`MonitoredRequestCountMetric` to determine the percentage of successful requests that
this SLO tracks.

_Required_: No

_Type_: [MonitoredRequestCountMetric](aws-properties-applicationsignals-servicelevelobjective-monitoredrequestcountmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperationName`

If the SLO monitors a specific operation of the service, this field displays that operation name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalRequestCountMetric`

This structure defines the metric that is used as the "total requests" number for a request-based SLO.
The number observed for this metric is divided by the number of "good requests" or "bad requests" that is
observed for the metric defined in
`MonitoredRequestCountMetric`.

_Required_: No

_Type_: Array of [MetricDataQuery](aws-properties-applicationsignals-servicelevelobjective-metricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestBasedSli

RollingInterval

All content copied from https://docs.aws.amazon.com/.
