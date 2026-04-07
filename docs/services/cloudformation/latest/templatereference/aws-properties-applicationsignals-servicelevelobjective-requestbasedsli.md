This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective RequestBasedSli

This structure contains information about the performance metric that a request-based SLO monitors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "MetricThreshold" : Number,
  "RequestBasedSliMetric" : RequestBasedSliMetric
}

```

### YAML

```yaml

  ComparisonOperator: String
  MetricThreshold: Number
  RequestBasedSliMetric:
    RequestBasedSliMetric

```

## Properties

`ComparisonOperator`

The arithmetic operation used when comparing the specified metric to the
threshold.

_Required_: No

_Type_: String

_Allowed values_: `GreaterThanOrEqualTo | LessThanOrEqualTo | LessThan | GreaterThan`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricThreshold`

This value is the threshold that
the observed metric values of the SLI metric are compared to.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestBasedSliMetric`

A structure that contains information about the metric that the SLO monitors.

_Required_: Yes

_Type_: [RequestBasedSliMetric](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationsignals-servicelevelobjective-requestbasedslimetric.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecurrenceRule

RequestBasedSliMetric
