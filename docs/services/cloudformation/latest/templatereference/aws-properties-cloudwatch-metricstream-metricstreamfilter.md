This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::MetricStream MetricStreamFilter

This structure contains a metric namespace and optionally, a list of metric names, to either include in a metric '
stream or exclude from a metric stream.

A metric stream's filters can include up to 1000 total names. This limit applies to the sum of namespace names
and metric names in the filters. For example, this could include 10 metric namespace filters with
99 metrics each, or 20 namespace filters with 49 metrics specified in each filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricNames" : [ String, ... ],
  "Namespace" : String
}

```

### YAML

```yaml

  MetricNames:
    - String
  Namespace: String

```

## Properties

`MetricNames`

The names of the metrics to either include or exclude from the metric stream.

If you omit this parameter, all metrics in the namespace are included or excluded, depending on whether this
filter is specified as an exclude filter or an include filter.

Each metric name can contain only ASCII printable characters (ASCII range 32 through 126). Each metric name
must contain at least one non-whitespace character.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255 | 999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The name of the metric namespace in the filter.

The namespace can contain only ASCII printable characters (ASCII range 32 through 126). It must
contain at least one non-whitespace character.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudWatch::MetricStream

MetricStreamStatisticsConfiguration
