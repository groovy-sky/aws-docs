This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup NetworkBandwidthGbpsRequest

`NetworkBandwidthGbpsRequest` is a property of the
`InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.md) property type that
describes the minimum and maximum network bandwidth for an instance type, in Gbps.

###### Note

Setting the minimum bandwidth does not guarantee that your instance will achieve the
minimum bandwidth. Amazon EC2 will identify instance types that support the specified
minimum bandwidth, but the actual bandwidth of your instance might go below the specified
minimum at times. For more information, see [Available instance bandwidth](../../../ec2/latest/userguide/ec2-instance-network-bandwidth.md#available-instance-bandwidth) in the _Amazon EC2 User Guide for Linux_
_Instances_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Number,
  "Min" : Number
}

```

### YAML

```yaml

  Max: Number
  Min: Number

```

## Properties

`Max`

The maximum amount of network bandwidth, in gigabits per second (Gbps).

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Min`

The minimum amount of network bandwidth, in gigabits per second (Gbps).

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MixedInstancesPolicy

NetworkInterfaceCountRequest

All content copied from https://docs.aws.amazon.com/.
