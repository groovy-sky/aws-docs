This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup MemoryGiBPerVCpuRequest

`MemoryGiBPerVCpuRequest` is a property of the `InstanceRequirements`
property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.md) property type that
describes the minimum and maximum amount of memory per vCPU for an instance type, in
GiB.

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

The memory maximum in GiB.

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Min`

The memory minimum in GiB.

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleHookSpecification

MemoryMiBRequest

All content copied from https://docs.aws.amazon.com/.
