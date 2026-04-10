This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet CustomerManagedWorkerCapabilities

The worker capabilities for a customer managed workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceleratorCount" : AcceleratorCountRange,
  "AcceleratorTotalMemoryMiB" : AcceleratorTotalMemoryMiBRange,
  "AcceleratorTypes" : [ String, ... ],
  "CpuArchitectureType" : String,
  "CustomAmounts" : [ FleetAmountCapability, ... ],
  "CustomAttributes" : [ FleetAttributeCapability, ... ],
  "MemoryMiB" : MemoryMiBRange,
  "OsFamily" : String,
  "VCpuCount" : VCpuCountRange
}

```

### YAML

```yaml

  AcceleratorCount:
    AcceleratorCountRange
  AcceleratorTotalMemoryMiB:
    AcceleratorTotalMemoryMiBRange
  AcceleratorTypes:
    - String
  CpuArchitectureType: String
  CustomAmounts:
    - FleetAmountCapability
  CustomAttributes:
    - FleetAttributeCapability
  MemoryMiB:
    MemoryMiBRange
  OsFamily: String
  VCpuCount:
    VCpuCountRange

```

## Properties

`AcceleratorCount`

The range of the accelerator.

_Required_: No

_Type_: [AcceleratorCountRange](aws-properties-deadline-fleet-acceleratorcountrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorTotalMemoryMiB`

The total memory (MiB) for the customer managed worker capabilities.

_Required_: No

_Type_: [AcceleratorTotalMemoryMiBRange](aws-properties-deadline-fleet-acceleratortotalmemorymibrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorTypes`

The accelerator types for the customer managed worker capabilities.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CpuArchitectureType`

The CPU architecture type for the customer managed worker capabilities.

_Required_: Yes

_Type_: String

_Allowed values_: `x86_64 | arm64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAmounts`

Custom requirement ranges for customer managed worker capabilities.

_Required_: No

_Type_: Array of [FleetAmountCapability](aws-properties-deadline-fleet-fleetamountcapability.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAttributes`

Custom attributes for the customer manged worker capabilities.

_Required_: No

_Type_: Array of [FleetAttributeCapability](aws-properties-deadline-fleet-fleetattributecapability.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryMiB`

The memory (MiB).

_Required_: Yes

_Type_: [MemoryMiBRange](aws-properties-deadline-fleet-memorymibrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OsFamily`

The operating system (OS) family.

_Required_: Yes

_Type_: String

_Allowed values_: `WINDOWS | LINUX | MACOS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VCpuCount`

The vCPU count for the customer manged worker capabilities.

_Required_: Yes

_Type_: [VCpuCountRange](aws-properties-deadline-fleet-vcpucountrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerManagedFleetConfiguration

Ec2EbsVolume

All content copied from https://docs.aws.amazon.com/.
