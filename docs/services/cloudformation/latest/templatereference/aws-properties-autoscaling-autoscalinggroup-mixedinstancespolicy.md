This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy

Use this structure to launch multiple instance types and On-Demand Instances and Spot
Instances within a single Auto Scaling group.

A mixed instances policy contains information that Amazon EC2 Auto Scaling can use to
launch instances and help optimize your costs. For more information, see [Auto Scaling\
groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the _Amazon EC2_
_Auto Scaling User Guide_.

You can create a mixed instances policy for new and existing Auto Scaling groups. You must
use a launch template to configure the policy. You cannot use a launch configuration.

There are key differences between Spot Instances and On-Demand Instances:

- The price for Spot Instances varies based on demand

- Amazon EC2 can terminate an individual Spot Instance as the availability of, or price
for, Spot Instances changes

When a Spot Instance is terminated, Amazon EC2 Auto Scaling group attempts to launch a
replacement instance to maintain the desired capacity for the group.

`MixedInstancesPolicy` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstancesDistribution" : InstancesDistribution,
  "LaunchTemplate" : LaunchTemplate
}

```

### YAML

```yaml

  InstancesDistribution:
    InstancesDistribution
  LaunchTemplate:
    LaunchTemplate

```

## Properties

`InstancesDistribution`

The instances distribution.

_Required_: No

_Type_: [InstancesDistribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-autoscalinggroup-instancesdistribution.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplate`

One or more launch templates and the instance types (overrides) that are used to
launch EC2 instances to fulfill On-Demand and Spot capacities.

_Required_: Yes

_Type_: [LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricsCollection

NetworkBandwidthGbpsRequest
