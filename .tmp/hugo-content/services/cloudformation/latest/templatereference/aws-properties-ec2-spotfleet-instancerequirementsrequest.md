This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet InstanceRequirementsRequest

The attributes for the instance types. When you specify instance attributes, Amazon EC2 will
identify instance types with these attributes.

You must specify `VCpuCount` and `MemoryMiB`. All other attributes
are optional. Any unspecified optional attribute is set to its default.

When you specify multiple attributes, you get instance types that satisfy all of the
specified attributes. If you specify multiple values for an attribute, you get instance
types that satisfy any of the specified values.

To limit the list of instance types from which Amazon EC2 can identify matching instance types,
you can use one of the following parameters, but not both in the same request:

- `AllowedInstanceTypes` \- The instance types to include in the list. All
other instance types are ignored, even if they match your specified attributes.

- `ExcludedInstanceTypes` \- The instance types to exclude from the list,
even if they match your specified attributes.

###### Note

If you specify `InstanceRequirements`, you can't specify
`InstanceType`.

Attribute-based instance type selection is only supported when using Auto Scaling
groups, EC2 Fleet, and Spot Fleet to launch instances. If you plan to use the launch template in
the [launch instance\
wizard](../../../ec2/latest/userguide/ec2-launch-instance-wizard.md), or with the [RunInstances](../../../../reference/awsec2/latest/apireference/api-runinstances.md) API or
[AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) AWS CloudFormation resource, you can't specify
`InstanceRequirements`.

For more information, see [Specify attributes for instance type selection for EC2 Fleet or Spot Fleet](../../../ec2/latest/userguide/ec2-fleet-attribute-based-instance-type-selection.md) and [Spot\
placement score](../../../ec2/latest/userguide/spot-placement-score.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceleratorCount" : AcceleratorCountRequest,
  "AcceleratorManufacturers" : [ String, ... ],
  "AcceleratorNames" : [ String, ... ],
  "AcceleratorTotalMemoryMiB" : AcceleratorTotalMemoryMiBRequest,
  "AcceleratorTypes" : [ String, ... ],
  "AllowedInstanceTypes" : [ String, ... ],
  "BareMetal" : String,
  "BaselineEbsBandwidthMbps" : BaselineEbsBandwidthMbpsRequest,
  "BaselinePerformanceFactors" : BaselinePerformanceFactorsRequest,
  "BurstablePerformance" : String,
  "CpuManufacturers" : [ String, ... ],
  "ExcludedInstanceTypes" : [ String, ... ],
  "InstanceGenerations" : [ String, ... ],
  "LocalStorage" : String,
  "LocalStorageTypes" : [ String, ... ],
  "MaxSpotPriceAsPercentageOfOptimalOnDemandPrice" : Integer,
  "MemoryGiBPerVCpu" : MemoryGiBPerVCpuRequest,
  "MemoryMiB" : MemoryMiBRequest,
  "NetworkBandwidthGbps" : NetworkBandwidthGbpsRequest,
  "NetworkInterfaceCount" : NetworkInterfaceCountRequest,
  "OnDemandMaxPricePercentageOverLowestPrice" : Integer,
  "RequireEncryptionInTransit" : Boolean,
  "RequireHibernateSupport" : Boolean,
  "SpotMaxPricePercentageOverLowestPrice" : Integer,
  "TotalLocalStorageGB" : TotalLocalStorageGBRequest,
  "VCpuCount" : VCpuCountRangeRequest
}

```

### YAML

```yaml

  AcceleratorCount:
    AcceleratorCountRequest
  AcceleratorManufacturers:
    - String
  AcceleratorNames:
    - String
  AcceleratorTotalMemoryMiB:
    AcceleratorTotalMemoryMiBRequest
  AcceleratorTypes:
    - String
  AllowedInstanceTypes:
    - String
  BareMetal: String
  BaselineEbsBandwidthMbps:
    BaselineEbsBandwidthMbpsRequest
  BaselinePerformanceFactors:
    BaselinePerformanceFactorsRequest
  BurstablePerformance: String
  CpuManufacturers:
    - String
  ExcludedInstanceTypes:
    - String
  InstanceGenerations:
    - String
  LocalStorage: String
  LocalStorageTypes:
    - String
  MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: Integer
  MemoryGiBPerVCpu:
    MemoryGiBPerVCpuRequest
  MemoryMiB:
    MemoryMiBRequest
  NetworkBandwidthGbps:
    NetworkBandwidthGbpsRequest
  NetworkInterfaceCount:
    NetworkInterfaceCountRequest
  OnDemandMaxPricePercentageOverLowestPrice: Integer
  RequireEncryptionInTransit: Boolean
  RequireHibernateSupport: Boolean
  SpotMaxPricePercentageOverLowestPrice: Integer
  TotalLocalStorageGB:
    TotalLocalStorageGBRequest
  VCpuCount:
    VCpuCountRangeRequest

```

## Properties

`AcceleratorCount`

The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on
an instance.

To exclude accelerator-enabled instance types, set `Max` to `0`.

Default: No minimum or maximum limits

_Required_: No

_Type_: [AcceleratorCountRequest](aws-properties-ec2-spotfleet-acceleratorcountrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AcceleratorManufacturers`

Indicates whether instance types must have accelerators by specific manufacturers.

- For instance types with AWS devices, specify `amazon-web-services`.

- For instance types with AMD devices, specify `amd`.

- For instance types with Habana devices, specify `habana`.

- For instance types with NVIDIA devices, specify `nvidia`.

- For instance types with Xilinx devices, specify `xilinx`.

Default: Any manufacturer

_Required_: No

_Type_: Array of String

_Allowed values_: `amazon-web-services | amd | habana | nvidia | xilinx`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AcceleratorNames`

The accelerators that must be on the instance type.

- For instance types with NVIDIA A10G GPUs, specify `a10g`.

- For instance types with NVIDIA A100 GPUs, specify `a100`.

- For instance types with NVIDIA H100 GPUs, specify `h100`.

- For instance types with AWS Inferentia chips, specify `inferentia`.

- For instance types with AWS Inferentia2 chips, specify `inferentia2`.

- For instance types with Habana Gaudi HL-205 GPUs, specify `gaudi-hl-205`.

- For instance types with NVIDIA GRID K520 GPUs, specify `k520`.

- For instance types with NVIDIA K80 GPUs, specify `k80`.

- For instance types with NVIDIA L4 GPUs, specify `l4`.

- For instance types with NVIDIA L40S GPUs, specify `l40s`.

- For instance types with NVIDIA M60 GPUs, specify `m60`.

- For instance types with AMD Radeon Pro V520 GPUs, specify `radeon-pro-v520`.

- For instance types with AWS Trainium chips, specify `trainium`.

- For instance types with AWS Trainium2 chips, specify `trainium2`.

- For instance types with NVIDIA T4 GPUs, specify `t4`.

- For instance types with NVIDIA T4G GPUs, specify `t4g`.

- For instance types with Xilinx U30 cards, specify `u30`.

- For instance types with Xilinx VU9P FPGAs, specify `vu9p`.

- For instance types with NVIDIA V100 GPUs, specify `v100`.

Default: Any accelerator

_Required_: No

_Type_: Array of String

_Allowed values_: `a10g | a100 | h100 | inferentia | k520 | k80 | m60 | radeon-pro-v520 | t4 | t4g | vu9p | v100 | l40s | l4 | gaudi-hl-205 | inferentia2 | trainium | trainium2 | u30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AcceleratorTotalMemoryMiB`

The minimum and maximum amount of total accelerator memory, in MiB.

Default: No minimum or maximum limits

_Required_: No

_Type_: [AcceleratorTotalMemoryMiBRequest](aws-properties-ec2-spotfleet-acceleratortotalmemorymibrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AcceleratorTypes`

The accelerator types that must be on the instance type.

- For instance types with FPGA accelerators, specify `fpga`.

- For instance types with GPU accelerators, specify `gpu`.

- For instance types with Inference accelerators, specify `inference`.

- For instance types with Media accelerators, specify `media`.

Default: Any accelerator type

_Required_: No

_Type_: Array of String

_Allowed values_: `gpu | fpga | inference | media`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AllowedInstanceTypes`

The instance types to apply your specified attributes against. All other instance types
are ignored, even if they match your specified attributes.

You can use strings with one or more wild cards, represented by
an asterisk ( `*`), to allow an instance type, size, or generation. The
following are examples: `m5.8xlarge`, `c5*.*`, `m5a.*`,
`r*`, `*3*`.

For example, if you specify `c5*`,Amazon EC2 will allow the entire C5 instance
family, which includes all C5a and C5n instance types. If you specify
`m5a.*`, Amazon EC2 will allow all the M5a instance types, but not the M5n
instance types.

###### Note

If you specify `AllowedInstanceTypes`, you can't specify `ExcludedInstanceTypes`.

Default: All instance types

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BareMetal`

Indicates whether bare metal instance types must be included, excluded, or required.

- To include bare metal instance types, specify `included`.

- To require only bare metal instance types, specify `required`.

- To exclude bare metal instance types, specify `excluded`.

Default: `excluded`

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BaselineEbsBandwidthMbps`

The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps. For more information, see
[Amazon\
EBS–optimized instances](../../../ec2/latest/userguide/ebs-optimized.md) in the _Amazon EC2 User Guide_.

Default: No minimum or maximum limits

_Required_: No

_Type_: [BaselineEbsBandwidthMbpsRequest](aws-properties-ec2-spotfleet-baselineebsbandwidthmbpsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BaselinePerformanceFactors`

The baseline performance to consider, using an instance family as a baseline reference.
The instance family establishes the lowest acceptable level of performance. Amazon EC2 uses this
baseline to guide instance type selection, but there is no guarantee that the selected
instance types will always exceed the baseline for every application. Currently, this
parameter only supports CPU performance as a baseline performance factor. For more
information, see [Performance protection](../../../ec2/latest/userguide/ec2-fleet-attribute-based-instance-type-selection.md#ec2fleet-abis-performance-protection) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: [BaselinePerformanceFactorsRequest](aws-properties-ec2-spotfleet-baselineperformancefactorsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BurstablePerformance`

Indicates whether burstable performance T instance types are included, excluded, or required. For more information, see
[Burstable performance instances](../../../ec2/latest/userguide/burstable-performance-instances.md).

- To include burstable performance instance types, specify `included`.

- To require only burstable performance instance types, specify `required`.

- To exclude burstable performance instance types, specify `excluded`.

Default: `excluded`

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CpuManufacturers`

The CPU manufacturers to include.

- For instance types with Intel CPUs, specify `intel`.

- For instance types with AMD CPUs, specify `amd`.

- For instance types with AWS CPUs, specify `amazon-web-services`.

- For instance types with Apple CPUs, specify `apple`.

###### Note

Don't confuse the CPU manufacturer with the CPU architecture. Instances will
be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you
specify in your launch template.

Default: Any manufacturer

_Required_: No

_Type_: Array of String

_Allowed values_: `intel | amd | amazon-web-services | apple`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExcludedInstanceTypes`

The instance types to exclude.

You can use strings with one or more wild cards, represented by
an asterisk ( `*`), to exclude an instance family, type, size, or generation. The
following are examples: `m5.8xlarge`, `c5*.*`, `m5a.*`,
`r*`, `*3*`.

For example, if you specify `c5*`,Amazon EC2 will exclude the entire C5 instance
family, which includes all C5a and C5n instance types. If you specify
`m5a.*`, Amazon EC2 will exclude all the M5a instance types, but not the M5n
instance types.

###### Note

If you specify `ExcludedInstanceTypes`, you can't specify `AllowedInstanceTypes`.

Default: No excluded instance types

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceGenerations`

Indicates whether current or previous generation instance types are included. The
current generation instance types are recommended for use. Current generation instance types are
typically the latest two to three generations in each instance family. For more
information, see [Instance types](../../../ec2/latest/userguide/instance-types.md) in the
_Amazon EC2 User Guide_.

For current generation instance types, specify `current`.

For previous generation instance types, specify `previous`.

Default: Current and previous generation instance types

_Required_: No

_Type_: Array of String

_Allowed values_: `current | previous`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalStorage`

Indicates whether instance types with instance store volumes are included, excluded, or required. For more information,
[Amazon\
EC2 instance store](../../../ec2/latest/userguide/instancestorage.md) in the _Amazon EC2 User Guide_.

- To include instance types with instance store volumes, specify
`included`.

- To require only instance types with instance store volumes, specify
`required`.

- To exclude instance types with instance store volumes, specify
`excluded`.

Default: `included`

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalStorageTypes`

The type of local storage that is required.

- For instance types with hard disk drive (HDD) storage, specify `hdd`.

- For instance types with solid state drive (SSD) storage, specify
`ssd`.

Default: `hdd` and `ssd`

_Required_: No

_Type_: Array of String

_Allowed values_: `hdd | ssd`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`

\[Price protection\] The price protection threshold for Spot Instances, as a percentage of an
identified On-Demand price. The identified On-Demand price is the price of the lowest
priced current generation C, M, or R instance type with your specified attributes. If no
current generation C, M, or R instance type matches your attributes, then the identified
price is from the lowest priced current generation instance types, and failing that, from
the lowest priced previous generation instance types that match your attributes. When Amazon EC2
selects instance types with your attributes, it will exclude instance types whose price
exceeds your specified threshold.

The parameter accepts an integer, which Amazon EC2 interprets as a percentage.

If you set `TargetCapacityUnitType` to `vcpu` or
`memory-mib`, the price protection threshold is based on the per vCPU or per
memory price instead of the per instance price.

###### Note

Only one of `SpotMaxPricePercentageOverLowestPrice` or
`MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` can be specified. If you
don't specify either, Amazon EC2 will automatically apply optimal price protection to
consistently select from a wide range of instance types. To indicate no price protection
threshold for Spot Instances, meaning you want to consider all instance types that match your
attributes, include one of these parameters and specify a high value, such as
`999999`.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemoryGiBPerVCpu`

The minimum and maximum amount of memory per vCPU, in GiB.

Default: No minimum or maximum limits

_Required_: No

_Type_: [MemoryGiBPerVCpuRequest](aws-properties-ec2-spotfleet-memorygibpervcpurequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemoryMiB`

The minimum and maximum amount of memory, in MiB.

_Required_: No

_Type_: [MemoryMiBRequest](aws-properties-ec2-spotfleet-memorymibrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkBandwidthGbps`

The minimum and maximum amount of baseline network bandwidth, in gigabits per second
(Gbps). For more information, see [Amazon EC2 instance network bandwidth](../../../ec2/latest/userguide/ec2-instance-network-bandwidth.md) in the _Amazon EC2 User Guide_.

Default: No minimum or maximum limits

_Required_: No

_Type_: [NetworkBandwidthGbpsRequest](aws-properties-ec2-spotfleet-networkbandwidthgbpsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceCount`

The minimum and maximum number of network interfaces.

Default: No minimum or maximum limits

_Required_: No

_Type_: [NetworkInterfaceCountRequest](aws-properties-ec2-spotfleet-networkinterfacecountrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandMaxPricePercentageOverLowestPrice`

\[Price protection\] The price protection threshold for On-Demand Instances, as a percentage higher than
an identified On-Demand price. The identified On-Demand price is the price of the lowest
priced current generation C, M, or R instance type with your specified attributes. When
Amazon EC2 selects instance types with your attributes, it will exclude instance types whose
price exceeds your specified threshold.

The parameter accepts an integer, which Amazon EC2 interprets as a percentage.

To indicate no price protection threshold, specify a high value, such as
`999999`.

This parameter is not supported for [GetSpotPlacementScores](../../../../reference/awsec2/latest/apireference/api-getspotplacementscores.md) and [GetInstanceTypesFromInstanceRequirements](../../../../reference/awsec2/latest/apireference/api-getinstancetypesfrominstancerequirements.md).

###### Note

If you set `TargetCapacityUnitType` to `vcpu` or
`memory-mib`, the price protection threshold is applied based on the
per-vCPU or per-memory price instead of the per-instance price.

Default: `20`

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequireEncryptionInTransit`

Specifies whether instance types must support encrypting in-transit traffic between
instances. For more information, including the supported instance types, see [Encryption in\
transit](../../../ec2/latest/userguide/data-protection.md#encryption-transit) in the _Amazon EC2 User Guide_.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequireHibernateSupport`

Indicates whether instance types must support hibernation for On-Demand Instances.

This parameter is not supported for [GetSpotPlacementScores](../../../../reference/awsec2/latest/apireference/api-getspotplacementscores.md).

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotMaxPricePercentageOverLowestPrice`

\[Price protection\] The price protection threshold for Spot Instances, as a percentage higher than
an identified Spot price. The identified Spot price is the Spot price of the lowest priced
current generation C, M, or R instance type with your specified attributes. If no current
generation C, M, or R instance type matches your attributes, then the identified Spot price
is from the lowest priced current generation instance types, and failing that, from the
lowest priced previous generation instance types that match your attributes. When Amazon EC2
selects instance types with your attributes, it will exclude instance types whose Spot
price exceeds your specified threshold.

The parameter accepts an integer, which Amazon EC2 interprets as a percentage.

If you set `TargetCapacityUnitType` to `vcpu` or
`memory-mib`, the price protection threshold is applied based on the
per-vCPU or per-memory price instead of the per-instance price.

This parameter is not supported for [GetSpotPlacementScores](../../../../reference/awsec2/latest/apireference/api-getspotplacementscores.md) and [GetInstanceTypesFromInstanceRequirements](../../../../reference/awsec2/latest/apireference/api-getinstancetypesfrominstancerequirements.md).

###### Note

Only one of `SpotMaxPricePercentageOverLowestPrice` or
`MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` can be specified. If you
don't specify either, Amazon EC2 will automatically apply optimal price protection to
consistently select from a wide range of instance types. To indicate no price protection
threshold for Spot Instances, meaning you want to consider all instance types that match your
attributes, include one of these parameters and specify a high value, such as
`999999`.

Default: `100`

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TotalLocalStorageGB`

The minimum and maximum amount of total local storage, in GB.

Default: No minimum or maximum limits

_Required_: No

_Type_: [TotalLocalStorageGBRequest](aws-properties-ec2-spotfleet-totallocalstoragegbrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VCpuCount`

The minimum and maximum number of vCPUs.

_Required_: No

_Type_: [VCpuCountRangeRequest](aws-properties-ec2-spotfleet-vcpucountrangerequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceNetworkInterfaceSpecification

LaunchTemplateConfig

All content copied from https://docs.aws.amazon.com/.
