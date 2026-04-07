This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider AutoScalingGroupProvider

The details of the Auto Scaling group for the capacity provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingGroupArn" : String,
  "ManagedDraining" : String,
  "ManagedScaling" : ManagedScaling,
  "ManagedTerminationProtection" : String
}

```

### YAML

```yaml

  AutoScalingGroupArn: String
  ManagedDraining: String
  ManagedScaling:
    ManagedScaling
  ManagedTerminationProtection: String

```

## Properties

`AutoScalingGroupArn`

The Amazon Resource Name (ARN) that identifies the Auto Scaling group, or the Auto
Scaling group name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedDraining`

The managed draining option for the Auto Scaling group capacity provider. When you
enable this, Amazon ECS manages and gracefully drains the EC2 container instances that
are in the Auto Scaling group capacity provider.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedScaling`

The managed scaling settings for the Auto Scaling group capacity provider.

_Required_: No

_Type_: [ManagedScaling](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-capacityprovider-managedscaling.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedTerminationProtection`

The managed termination protection setting to use for the Auto Scaling group capacity
provider. This determines whether the Auto Scaling group has managed termination
protection. The default is off.

###### Important

When using managed termination protection, managed scaling must also be used
otherwise managed termination protection doesn't work.

When managed termination protection is on, Amazon ECS prevents the Amazon EC2
instances in an Auto Scaling group that contain tasks from being terminated during a
scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must
have instance protection from scale-in actions on as well. For more information, see
[Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the _AWS Auto Scaling User_
_Guide_.

When managed termination protection is off, your Amazon EC2 instances aren't protected
from termination when the Auto Scaling group scales in.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Defining an Amazon ECS capacity provider](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-capacityprovider.html#aws-resource-ecs-capacityprovider--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AcceleratorTotalMemoryMiBRequest

BaselineEbsBandwidthMbpsRequest
