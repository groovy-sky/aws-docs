This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector MetricCharacteristics

This object includes parameters that you can use to provide information to CloudWatch to help it build more accurate anomaly detection models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PeriodicSpikes" : Boolean
}

```

### YAML

```yaml

  PeriodicSpikes: Boolean

```

## Properties

`PeriodicSpikes`

Set this parameter to true if values for this metric consistently include spikes that should not be considered to be anomalies. With this set to true,
CloudWatch will expect to see spikes that occurred consistently during the model training period, and won't flag future similar spikes as anomalies.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Metric

MetricDataQuery
