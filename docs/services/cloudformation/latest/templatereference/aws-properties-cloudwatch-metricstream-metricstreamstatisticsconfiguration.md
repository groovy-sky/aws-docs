This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::MetricStream MetricStreamStatisticsConfiguration

This structure specifies a list of additional statistics to stream, and the metrics to stream those
additional statistics for.

All metrics that match the combination of metric name and namespace will be streamed
with the additional statistics, no matter their dimensions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalStatistics" : [ String, ... ],
  "IncludeMetrics" : [ MetricStreamStatisticsMetric, ... ]
}

```

### YAML

```yaml

  AdditionalStatistics:
    - String
  IncludeMetrics:
    - MetricStreamStatisticsMetric

```

## Properties

`AdditionalStatistics`

The
additional statistics to stream for the metrics listed in `IncludeMetrics`.

_Required_: Yes

_Type_: Array of String

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeMetrics`

An array that
defines the metrics that are to have additional statistics streamed.

_Required_: Yes

_Type_: Array of [MetricStreamStatisticsMetric](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-metricstream-metricstreamstatisticsmetric.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricStreamFilter

MetricStreamStatisticsMetric
