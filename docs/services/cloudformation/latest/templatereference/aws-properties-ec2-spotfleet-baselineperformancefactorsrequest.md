This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet BaselinePerformanceFactorsRequest

The baseline performance to consider, using an instance family as a baseline reference.
The instance family establishes the lowest acceptable level of performance. Amazon EC2 uses this
baseline to guide instance type selection, but there is no guarantee that the selected
instance types will always exceed the baseline for every application.

Currently, this parameter only supports CPU performance as a baseline performance
factor. For example, specifying `c6i` would use the CPU performance of the
`c6i` family as the baseline reference.

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

_Type_: [CpuPerformanceFactorRequest](aws-properties-ec2-spotfleet-cpuperformancefactorrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaselineEbsBandwidthMbpsRequest

BlockDeviceMapping

All content copied from https://docs.aws.amazon.com/.
