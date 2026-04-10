This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile StatisticalThreshold

A statistical ranking (percentile) that
indicates a threshold value by which a behavior is determined to be in compliance or in
violation of the behavior.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Statistic" : String
}

```

### YAML

```yaml

  Statistic: String

```

## Properties

`Statistic`

The percentile that
resolves to a threshold value by which compliance with a behavior is determined. Metrics are
collected over the specified period ( `durationSeconds`) from all reporting devices
in your account and statistical ranks are calculated. Then, the measurements from a device are
collected over the same period. If the accumulated measurements from the device fall above or
below ( `comparisonOperator`) the value associated with the percentile specified,
then the device is considered to be in compliance with the behavior, otherwise a violation
occurs.

_Required_: No

_Type_: String

_Allowed values_: `Average | p0 | p0.1 | p0.01 | p1 | p10 | p50 | p90 | p99 | p99.9 | p99.99 | p100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricValue

Tag

All content copied from https://docs.aws.amazon.com/.
