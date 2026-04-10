This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet ServiceManagedEc2InstanceCapabilities

The Amazon EC2 instance capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceleratorCapabilities" : AcceleratorCapabilities,
  "AllowedInstanceTypes" : [ String, ... ],
  "CpuArchitectureType" : String,
  "CustomAmounts" : [ FleetAmountCapability, ... ],
  "CustomAttributes" : [ FleetAttributeCapability, ... ],
  "ExcludedInstanceTypes" : [ String, ... ],
  "MemoryMiB" : MemoryMiBRange,
  "OsFamily" : String,
  "RootEbsVolume" : Ec2EbsVolume,
  "VCpuCount" : VCpuCountRange
}

```

### YAML

```yaml

  AcceleratorCapabilities:
    AcceleratorCapabilities
  AllowedInstanceTypes:
    - String
  CpuArchitectureType: String
  CustomAmounts:
    - FleetAmountCapability
  CustomAttributes:
    - FleetAttributeCapability
  ExcludedInstanceTypes:
    - String
  MemoryMiB:
    MemoryMiBRange
  OsFamily: String
  RootEbsVolume:
    Ec2EbsVolume
  VCpuCount:
    VCpuCountRange

```

## Properties

`AcceleratorCapabilities`

Describes the GPU accelerator capabilities required for worker host instances in this
fleet.

_Required_: No

_Type_: [AcceleratorCapabilities](aws-properties-deadline-fleet-acceleratorcapabilities.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedInstanceTypes`

The allowable Amazon EC2 instance types.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CpuArchitectureType`

The CPU architecture type.

_Required_: Yes

_Type_: String

_Allowed values_: `x86_64 | arm64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAmounts`

The custom capability amounts to require for instances in this fleet.

_Required_: No

_Type_: Array of [FleetAmountCapability](aws-properties-deadline-fleet-fleetamountcapability.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAttributes`

The custom capability attributes to require for instances in this fleet.

_Required_: No

_Type_: Array of [FleetAttributeCapability](aws-properties-deadline-fleet-fleetattributecapability.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedInstanceTypes`

The instance types to exclude from the fleet.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryMiB`

The memory, as MiB, for the Amazon EC2 instance type.

_Required_: Yes

_Type_: [MemoryMiBRange](aws-properties-deadline-fleet-memorymibrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OsFamily`

The operating system (OS) family.

_Required_: Yes

_Type_: String

_Allowed values_: `LINUX | WINDOWS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootEbsVolume`

The root EBS volume.

_Required_: No

_Type_: [Ec2EbsVolume](aws-properties-deadline-fleet-ec2ebsvolume.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VCpuCount`

The amount of vCPU to require for instances in this fleet.

_Required_: Yes

_Type_: [VCpuCountRange](aws-properties-deadline-fleet-vcpucountrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceManagedEc2FleetConfiguration

ServiceManagedEc2InstanceMarketOptions

All content copied from https://docs.aws.amazon.com/.
