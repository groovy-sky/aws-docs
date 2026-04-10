This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::MetricStream MetricStreamStatisticsMetric

A structure that specifies the
metric name and namespace for one metric that is going to have additional statistics included in the stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  MetricName: String
  Namespace: String

```

## Properties

`MetricName`

The name of the metric.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricStreamStatisticsConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
