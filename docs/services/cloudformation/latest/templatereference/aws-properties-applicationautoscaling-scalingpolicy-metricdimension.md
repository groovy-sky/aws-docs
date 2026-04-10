This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy MetricDimension

`MetricDimension` specifies a name/value pair that is part of the identity of a
CloudWatch metric for the `Dimensions` property of the [AWS::ApplicationAutoScaling::ScalingPolicy CustomizedMetricSpecification](../userguide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.md) property
type. Duplicate dimensions are not allowed.

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

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomizedMetricSpecification

PredefinedMetricSpecification

All content copied from https://docs.aws.amazon.com/.
