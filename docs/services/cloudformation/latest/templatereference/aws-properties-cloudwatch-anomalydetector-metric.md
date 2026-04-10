This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector Metric

Represents a specific metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ Dimension, ... ],
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  Dimensions:
    - Dimension
  MetricName: String
  Namespace: String

```

## Properties

`Dimensions`

The dimensions for the metric.

_Required_: No

_Type_: Array of [Dimension](aws-properties-cloudwatch-anomalydetector-dimension.md)

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricName`

The name of the metric. This is a required field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

The namespace of the metric.

_Required_: Yes

_Type_: String

_Pattern_: `[^:].*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

MetricCharacteristics

All content copied from https://docs.aws.amazon.com/.
