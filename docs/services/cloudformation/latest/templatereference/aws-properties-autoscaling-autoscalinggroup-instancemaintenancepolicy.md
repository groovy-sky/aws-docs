This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup InstanceMaintenancePolicy

`InstanceMaintenancePolicy` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.

For more information, see [Instance\
maintenance policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html) in the _Amazon EC2 Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxHealthyPercentage" : Integer,
  "MinHealthyPercentage" : Integer
}

```

### YAML

```yaml

  MaxHealthyPercentage: Integer
  MinHealthyPercentage: Integer

```

## Properties

`MaxHealthyPercentage`

Specifies the upper threshold as a percentage of the desired capacity of the Auto Scaling
group. It represents the maximum percentage of the group that can be in service and
healthy, or pending, to support your workload when replacing instances. Value range is
100 to 200. To clear a previously set value, specify a value of `-1`.

Both `MinHealthyPercentage` and `MaxHealthyPercentage` must be
specified, and the difference between them cannot be greater than 100. A large range
increases the number of instances that can be replaced at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinHealthyPercentage`

Specifies the lower threshold as a percentage of the desired capacity of the Auto Scaling
group. It represents the minimum percentage of the group to keep in service, healthy,
and ready to use to support your workload when replacing instances. Value range is 0 to
100\. To clear a previously set value, specify a value of `-1`.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceLifecyclePolicy

InstanceRequirements
