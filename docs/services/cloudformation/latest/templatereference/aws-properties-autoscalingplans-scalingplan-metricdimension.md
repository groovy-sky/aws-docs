This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan MetricDimension

`MetricDimension` is a subproperty of [CustomizedScalingMetricSpecification](../userguide/aws-properties-autoscalingplans-scalingplan-customizedscalingmetricspecification.md) that specifies a dimension for a
customized metric to use with a scaling plan.
Dimensions are arbitrary name/value pairs that can be associated with a CloudWatch metric.
Duplicate dimensions are not allowed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the dimension.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the dimension.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomizedScalingMetricSpecification

PredefinedLoadMetricSpecification

All content copied from https://docs.aws.amazon.com/.
