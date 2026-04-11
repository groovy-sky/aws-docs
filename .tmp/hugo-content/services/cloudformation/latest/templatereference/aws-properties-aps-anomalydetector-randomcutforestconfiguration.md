This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::AnomalyDetector RandomCutForestConfiguration

Configuration for the Random Cut Forest algorithm used for anomaly detection in
time-series data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IgnoreNearExpectedFromAbove" : IgnoreNearExpected,
  "IgnoreNearExpectedFromBelow" : IgnoreNearExpected,
  "Query" : String,
  "SampleSize" : Integer,
  "ShingleSize" : Integer
}

```

### YAML

```yaml

  IgnoreNearExpectedFromAbove:
    IgnoreNearExpected
  IgnoreNearExpectedFromBelow:
    IgnoreNearExpected
  Query: String
  SampleSize: Integer
  ShingleSize: Integer

```

## Properties

`IgnoreNearExpectedFromAbove`

Configuration for ignoring values that are near expected values from above during
anomaly detection.

_Required_: No

_Type_: [IgnoreNearExpected](aws-properties-aps-anomalydetector-ignorenearexpected.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnoreNearExpectedFromBelow`

Configuration for ignoring values that are near expected values from below during
anomaly detection.

_Required_: No

_Type_: [IgnoreNearExpected](aws-properties-aps-anomalydetector-ignorenearexpected.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Query`

The Prometheus query used to retrieve the time-series data for anomaly
detection.

###### Important

Random Cut Forest queries must be wrapped by a supported PromQL aggregation
operator. For more information, see [Aggregation operators](https://prometheus.io/docs/prometheus/latest/querying/operators) on the _Prometheus docs_
website.

**Supported PromQL aggregation operators**:
`avg`, `count`, `group`, `max`,
`min`, `quantile`, `stddev`,
`stdvar`, and `sum`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleSize`

The number of data points sampled from the input stream for the Random Cut Forest
algorithm. The default number is 256 consecutive data points.

_Required_: No

_Type_: Integer

_Minimum_: `256`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShingleSize`

The number of consecutive data points used to create a shingle for the Random Cut
Forest algorithm. The default number is 8 consecutive data points.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MissingDataAction

Tag

All content copied from https://docs.aws.amazon.com/.
