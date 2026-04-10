This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector MetricMathAnomalyDetector

Indicates the CloudWatch math expression that provides the time series the anomaly
detector uses as input. The designated math expression must return a single time
series.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricDataQueries" : [ MetricDataQuery, ... ]
}

```

### YAML

```yaml

  MetricDataQueries:
    - MetricDataQuery

```

## Properties

`MetricDataQueries`

An array of metric data query structures that enables you to create an anomaly
detector based on the result of a metric math expression. Each item in
`MetricDataQueries` gets a metric or performs a math expression. One item
in `MetricDataQueries` is the expression that provides the time series that
the anomaly detector uses as input. Designate the expression by setting
`ReturnData` to `true` for this object in the array. For all
other expressions and metrics, set `ReturnData` to `false`. The
designated expression must return a single time series.

_Required_: No

_Type_: Array of [MetricDataQuery](aws-properties-cloudwatch-anomalydetector-metricdataquery.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDataQuery

MetricStat

All content copied from https://docs.aws.amazon.com/.
