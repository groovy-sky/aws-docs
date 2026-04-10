This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup BaselineEbsBandwidthMbpsRequest

`BaselineEbsBandwidthMbpsRequest` is a property of the
`InstanceRequirements` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.md) property type that
describes the minimum and maximum baseline bandwidth performance for an instance type, in
Mbps.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Integer,
  "Min" : Integer
}

```

### YAML

```yaml

  Max: Integer
  Min: Integer

```

## Properties

`Max`

The maximum value in Mbps.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Min`

The minimum value in Mbps.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailabilityZoneImpairmentPolicy

BaselinePerformanceFactorsRequest

All content copied from https://docs.aws.amazon.com/.
