This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides

Use this structure to let Amazon EC2 Auto Scaling do the following when the Auto Scaling
group has a mixed instances policy:

- Override the instance type that is specified in the launch template.

- Use multiple instance types.

Specify the instance types that you want, or define your instance requirements instead and
let Amazon EC2 Auto Scaling provision the available instance types that meet your
requirements. This can provide Amazon EC2 Auto Scaling with a larger selection of instance
types to choose from when fulfilling Spot and On-Demand capacities. You can view which
instance types are matched before you apply the instance requirements to your Auto Scaling
group.

After you define your instance requirements, you don't have to keep updating these
settings to get new EC2 instance types automatically. Amazon EC2 Auto Scaling uses the
instance requirements of the Auto Scaling group to determine whether a new EC2 instance type
can be used.

`LaunchTemplateOverrides` is a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageId" : String,
  "InstanceRequirements" : InstanceRequirements,
  "InstanceType" : String,
  "LaunchTemplateSpecification" : LaunchTemplateSpecification,
  "WeightedCapacity" : String
}

```

### YAML

```yaml

  ImageId: String
  InstanceRequirements:
    InstanceRequirements
  InstanceType: String
  LaunchTemplateSpecification:
    LaunchTemplateSpecification
  WeightedCapacity: String

```

## Properties

`ImageId`

The ID of the Amazon Machine Image (AMI) to use for instances launched with this override. When
using Instance Refresh with `ReplaceRootVolume` strategy, this specifies the AMI for root volume
replacement operations.

For `ReplaceRootVolume` operations:

- All overrides in the `MixedInstancesPolicy` must specify an ImageId

- The AMI must contain only a single root volume

- Root volume replacement doesn't support multi-volume AMIs

_Required_: No

_Type_: String

_Pattern_: `^ami-[a-z0-9]{1,17}$`

_Minimum_: `5`

_Maximum_: `21`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceRequirements`

The instance requirements. Amazon EC2 Auto Scaling uses your specified requirements to identify
instance types. Then, it uses your On-Demand and Spot allocation strategies to launch
instances from these instance types.

You can specify up to four separate sets of instance requirements per Auto Scaling group. This
is useful for provisioning instances from different Amazon Machine Images (AMIs) in the
same Auto Scaling group. To do this, create the AMIs and create a new launch template for each
AMI. Then, create a compatible set of instance requirements for each launch template.

###### Note

If you specify `InstanceRequirements`, you can't specify
`InstanceType`.

_Required_: No

_Type_: [InstanceRequirements](aws-properties-autoscaling-autoscalinggroup-instancerequirements.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceType`

The instance type, such as `m3.xlarge`. You must specify an instance type
that is supported in your requested Region and Availability Zones. For more information,
see [Instance types](../../../ec2/latest/userguide/instance-types.md) in the _Amazon EC2 User Guide_.

You can specify up to 40 instance types per Auto Scaling group.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplateSpecification`

Provides a launch template for the specified instance type or set of instance
requirements. For example, some instance types might require a launch template with a
different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's specified in
the `LaunchTemplate` definition. For more information, see [Specifying a different launch template for an instance type](../../../autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.md) in the
_Amazon EC2 Auto Scaling User Guide_.

You can specify up to 20 launch templates per Auto Scaling group. The launch templates
specified in the overrides and in the `LaunchTemplate` definition count
towards this limit.

_Required_: No

_Type_: [LaunchTemplateSpecification](aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WeightedCapacity`

If you provide a list of instance types to use, you can specify the number of capacity
units provided by each instance type in terms of virtual CPUs, memory, storage, throughput, or
other relative performance characteristic. When a Spot or On-Demand Instance is launched, the
capacity units count toward the desired capacity. Amazon EC2 Auto Scaling launches instances
until the desired capacity is totally fulfilled, even if this results in an overage. For
example, if there are two units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can
only launch an instance with a `WeightedCapacity` of five units, the instance is
launched, and the desired capacity is exceeded by three units. For more information, see
[Configure instance weighting for Amazon EC2 Auto Scaling](../../../autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.md) in the _Amazon_
_EC2 Auto Scaling User Guide_. Value must be in the range of 1-999.

If you specify a value for `WeightedCapacity` for one instance type, you must
specify a value for `WeightedCapacity` for all of them.

###### Important

Every Auto Scaling group has three size parameters ( `DesiredCapacity`,
`MaxSize`, and `MinSize`). Usually, you set these sizes based on a
specific number of instances. However, if you configure a mixed instances policy that
defines weights for the instance types, you must specify these sizes with the same units
that you use for weighting instances.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplate

LaunchTemplateSpecification

All content copied from https://docs.aws.amazon.com/.
