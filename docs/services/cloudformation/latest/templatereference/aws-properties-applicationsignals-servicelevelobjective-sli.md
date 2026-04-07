This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective Sli

This structure specifies the information about the service and the performance metric that an SLO is to monitor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "MetricThreshold" : Number,
  "SliMetric" : SliMetric
}

```

### YAML

```yaml

  ComparisonOperator: String
  MetricThreshold: Number
  SliMetric:
    SliMetric

```

## Properties

`ComparisonOperator`

The arithmetic operation to use when comparing the specified metric to the threshold.

_Required_: Yes

_Type_: String

_Allowed values_: `GreaterThanOrEqualTo | LessThanOrEqualTo | LessThan | GreaterThan`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricThreshold`

The value that the SLI metric is compared to.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SliMetric`

Use this structure to specify the metric to be used for the SLO.

_Required_: Yes

_Type_: [SliMetric](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationsignals-servicelevelobjective-slimetric.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RollingInterval

SliMetric
