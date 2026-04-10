This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective MonitoredRequestCountMetric

This structure defines the metric that is used as the "good request" or "bad request"
value for a request-based SLO.
This value observed for the metric defined in
`TotalRequestCountMetric` is divided by the number found for
`MonitoredRequestCountMetric` to determine the percentage of successful requests that
this SLO tracks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BadCountMetric" : [ MetricDataQuery, ... ],
  "GoodCountMetric" : [ MetricDataQuery, ... ]
}

```

### YAML

```yaml

  BadCountMetric:
    - MetricDataQuery
  GoodCountMetric:
    - MetricDataQuery

```

## Properties

`BadCountMetric`

If you want to count "bad requests" to determine the percentage of successful requests for this
request-based SLO, specify the metric to use as "bad requests" in this structure.

_Required_: No

_Type_: Array of [MetricDataQuery](aws-properties-applicationsignals-servicelevelobjective-metricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoodCountMetric`

If you want to count "good requests" to determine the percentage of successful requests for this
request-based SLO, specify the metric to use as "good requests" in this structure.

_Required_: No

_Type_: Array of [MetricDataQuery](aws-properties-applicationsignals-servicelevelobjective-metricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricStat

RecurrenceRule

All content copied from https://docs.aws.amazon.com/.
