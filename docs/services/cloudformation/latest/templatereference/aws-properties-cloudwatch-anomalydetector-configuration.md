This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector Configuration

Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude
when training and updating the model. The configuration can also include the time zone to use for the metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludedTimeRanges" : [ Range, ... ],
  "MetricTimeZone" : String
}

```

### YAML

```yaml

  ExcludedTimeRanges:
    - Range
  MetricTimeZone: String

```

## Properties

`ExcludedTimeRanges`

Specifies an array of time ranges to exclude from use when the anomaly detection model is trained and updated.
Use this to make sure that events that could cause unusual values for the metric, such as deployments, aren't used when
CloudWatch creates or updates the model.

_Required_: No

_Type_: Array of [Range](aws-properties-cloudwatch-anomalydetector-range.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricTimeZone`

The time zone to use for the metric. This is useful to enable the model to automatically account for daylight savings
time changes if the metric is sensitive to such time changes.

To specify a time zone, use the name of the time zone as specified in the standard tz database. For more information,
see [tz database](https://en.wikipedia.org/wiki/Tz_database).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudWatch::AnomalyDetector

Dimension

All content copied from https://docs.aws.amazon.com/.
