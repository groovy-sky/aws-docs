This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup BaselinePerformanceFactorsRequest

The baseline performance to consider, using an instance family as a baseline reference. The instance family establishes the lowest acceptable level of performance. Auto Scaling uses
this baseline to guide instance type selection, but there is no guarantee that the selected instance types will always exceed the baseline for every application.

Currently, this parameter only supports CPU performance as a baseline performance factor. For example, specifying `c6i` uses the CPU performance of the `c6i`
family as the baseline reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cpu" : CpuPerformanceFactorRequest
}

```

### YAML

```yaml

  Cpu:
    CpuPerformanceFactorRequest

```

## Properties

`Cpu`

The CPU performance to consider, using an instance family as the baseline reference.

_Required_: No

_Type_: [CpuPerformanceFactorRequest](aws-properties-autoscaling-autoscalinggroup-cpuperformancefactorrequest.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaselineEbsBandwidthMbpsRequest

CapacityReservationSpecification

All content copied from https://docs.aws.amazon.com/.
