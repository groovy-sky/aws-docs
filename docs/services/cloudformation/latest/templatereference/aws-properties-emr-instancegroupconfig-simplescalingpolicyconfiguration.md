This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig SimpleScalingPolicyConfiguration

`SimpleScalingPolicyConfiguration` is a subproperty of the `ScalingAction` property type. `SimpleScalingPolicyConfiguration` determines how an automatic scaling action adds or removes instances, the cooldown period, and the number of EC2 instances that are added each time the CloudWatch metric alarm condition is satisfied.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdjustmentType" : String,
  "CoolDown" : Integer,
  "ScalingAdjustment" : Integer
}

```

### YAML

```yaml

  AdjustmentType: String
  CoolDown: Integer
  ScalingAdjustment: Integer

```

## Properties

`AdjustmentType`

The way in which Amazon EC2 instances are added (if
`ScalingAdjustment` is a positive number) or terminated (if
`ScalingAdjustment` is a negative number) each time the scaling activity is
triggered. `CHANGE_IN_CAPACITY` is the default. `CHANGE_IN_CAPACITY`
indicates that the Amazon EC2 instance count increments or decrements by
`ScalingAdjustment`, which should be expressed as an integer.
`PERCENT_CHANGE_IN_CAPACITY` indicates the instance count increments or
decrements by the percentage specified by `ScalingAdjustment`, which should be
expressed as an integer. For example, 20 indicates an increase in 20% increments of cluster
capacity. `EXACT_CAPACITY` indicates the scaling activity results in an instance
group with the number of Amazon EC2 instances specified by
`ScalingAdjustment`, which should be expressed as a positive integer.

_Required_: No

_Type_: String

_Allowed values_: `CHANGE_IN_CAPACITY | PERCENT_CHANGE_IN_CAPACITY | EXACT_CAPACITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoolDown`

The amount of time, in seconds, after a scaling activity completes before any further
trigger-related scaling activities can start. The default value is 0.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingAdjustment`

The amount by which to scale in or scale out, based on the specified
`AdjustmentType`. A positive value adds to the instance group's Amazon EC2 instance count while a negative number removes instances. If
`AdjustmentType` is set to `EXACT_CAPACITY`, the number should
only be a positive integer. If `AdjustmentType` is set to
`PERCENT_CHANGE_IN_CAPACITY`, the value should express the percentage as an
integer. For example, -20 indicates a decrease in 20% increments of cluster
capacity.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingTrigger

VolumeSpecification

All content copied from https://docs.aws.amazon.com/.
