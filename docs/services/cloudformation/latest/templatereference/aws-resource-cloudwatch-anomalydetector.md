This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector

The `AWS::CloudWatch::AnomalyDetector` type specifies an anomaly detection band for a certain metric and statistic. The band
represents the expected "normal" range for the metric values. Anomaly detection bands can be used for visualization of a metric's expected values,
and for alarms.

For more information see [Using CloudWatch anomaly detection.](../../../amazoncloudwatch/latest/monitoring/cloudwatch-anomaly-detection.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::AnomalyDetector",
  "Properties" : {
      "Configuration" : Configuration,
      "Dimensions" : [ Dimension, ... ],
      "MetricCharacteristics" : MetricCharacteristics,
      "MetricMathAnomalyDetector" : MetricMathAnomalyDetector,
      "MetricName" : String,
      "Namespace" : String,
      "SingleMetricAnomalyDetector" : SingleMetricAnomalyDetector,
      "Stat" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::AnomalyDetector
Properties:
  Configuration:
    Configuration
  Dimensions:
    - Dimension
  MetricCharacteristics:
    MetricCharacteristics
  MetricMathAnomalyDetector:
    MetricMathAnomalyDetector
  MetricName: String
  Namespace: String
  SingleMetricAnomalyDetector:
    SingleMetricAnomalyDetector
  Stat: String

```

## Properties

`Configuration`

Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude
when training and updating the model. The configuration can also include the time zone to use for the metric.

_Required_: No

_Type_: [Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-anomalydetector-configuration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

The dimensions of the metric associated with the anomaly detection band.

_Required_: No

_Type_: Array of [Dimension](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-anomalydetector-dimension.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricCharacteristics`

Use this object to include parameters to provide information about your metric to CloudWatch to help it build more accurate anomaly detection models.
Currently, it includes the `PeriodicSpikes` parameter.

_Required_: No

_Type_: [MetricCharacteristics](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-anomalydetector-metriccharacteristics.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricMathAnomalyDetector`

The CloudWatch metric math expression for this anomaly detector.

_Required_: No

_Type_: [MetricMathAnomalyDetector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-anomalydetector-metricmathanomalydetector.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricName`

The name of the metric associated with the anomaly detection band.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The namespace of the metric associated with the anomaly detection band.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleMetricAnomalyDetector`

The CloudWatch metric and statistic for this anomaly detector.

_Required_: No

_Type_: [SingleMetricAnomalyDetector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Stat`

The statistic of the metric associated with the anomaly detection band.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

## Examples

### Anomaly Detector

This example creates an anomaly detector model for the metric named `JvmMetric` with the dimension value of `UsedMemory`.
It excludes a time range from the model training.

#### JSON

```json

{
    "Description": "AnomalyDetectorOnUsedMemory",
    "Resources": {
        "AnomalyDetectorOnUsedMemory": {
            "Type": "AWS::CloudWatch::AnomalyDetector",
            "Properties": {
                "MetricName": "JvmMetric",
                "Namespace": "AWSSDK/Java",
                "Stat": "Average",
                "Dimensions": [
                    {
                        "Name": "Memory",
                        "Value": "UsedMemory"
                    }
                ],
                "Configuration": {
                    "MetricTimeZone": "UTC",
                    "ExcludedTimeRanges": [
                        {
                            "StartTime": "2019-07-01T00:00:00",
                            "EndTime": "2019-07-01T23:59:59"
                        }
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

Description: AnomalyDetectorOnUsedMemory
Resources:
  AnomalyDetectorOnUsedMemory:
    Type: AWS::CloudWatch::AnomalyDetector
    Properties:
      MetricName: JvmMetric
      Namespace: AWSSDK/Java
      Stat: Average
      Dimensions:
        - Name: Memory
          Value: UsedMemory
      Configuration:
        MetricTimeZone: UTC
        ExcludedTimeRanges:
          - StartTime: 2019-07-01T00:00:00
            EndTime: 2019-07-01T23:59:59

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Configuration
