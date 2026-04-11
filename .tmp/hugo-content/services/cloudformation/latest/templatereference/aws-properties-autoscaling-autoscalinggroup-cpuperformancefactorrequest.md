This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup CpuPerformanceFactorRequest

The CPU performance to consider, using an instance family as the baseline reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "References" : [ PerformanceFactorReferenceRequest, ... ]
}

```

### YAML

```yaml

  References:
    - PerformanceFactorReferenceRequest

```

## Properties

`References`

Specify an instance family to use as the baseline reference for CPU performance. All instance types that match your specified attributes will be compared against the CPU performance of the
referenced instance family, regardless of CPU manufacturer or architecture differences.

###### Note

Currently only one instance family can be specified in the list.

_Required_: No

_Type_: Array of [PerformanceFactorReferenceRequest](aws-properties-autoscaling-autoscalinggroup-performancefactorreferencerequest.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityReservationTarget

InstanceLifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
