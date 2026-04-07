This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan ApplicationSource

`ApplicationSource` is a property of [ScalingPlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html) that specifies the application source to use with a scaling plan. You can create one scaling plan per application source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudFormationStackARN" : String,
  "TagFilters" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  CloudFormationStackARN: String
  TagFilters:
    - TagFilter

```

## Properties

`CloudFormationStackARN`

The Amazon Resource Name (ARN) of a CloudFormation stack.

You must specify either a `CloudFormationStackARN` or
`TagFilters`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

A set of tag filters (keys and values). Each tag filter specified must contain a key
with values as optional. Each scaling plan can include up to 50 keys, and each key can
include up to 20 values.

Tags are part of the syntax that you use to specify the resources you want returned when
configuring a scaling plan from the AWS Auto Scaling console. You do not need to
specify valid tag filter values when you create a scaling plan with CloudFormation. The
`Key` and `Values` properties can accept any value as long as the
combination of values is unique across scaling plans. However, if you also want to use the
AWS Auto Scaling console to edit the scaling plan, then you must specify valid
values.

You must specify either a `CloudFormationStackARN` or
`TagFilters`.

_Required_: No

_Type_: Array of [TagFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscalingplans-scalingplan-tagfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AutoScalingPlans::ScalingPlan

CustomizedLoadMetricSpecification
