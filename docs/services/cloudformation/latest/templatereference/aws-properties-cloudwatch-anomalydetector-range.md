This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector Range

Each `Range` specifies one range of days or times to exclude from use for training or updating an
anomaly detection model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  EndTime: String
  StartTime: String

```

## Properties

`EndTime`

The end time of the range to exclude. The format is `yyyy-MM-dd'T'HH:mm:ss`. For example,
`2019-07-01T23:59:59`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time of the range to exclude. The format is `yyyy-MM-dd'T'HH:mm:ss`. For example,
`2019-07-01T23:59:59`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricStat

SingleMetricAnomalyDetector

All content copied from https://docs.aws.amazon.com/.
