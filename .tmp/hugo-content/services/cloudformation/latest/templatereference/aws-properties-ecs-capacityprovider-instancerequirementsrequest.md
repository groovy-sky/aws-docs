This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider InstanceRequirementsRequest

The instance requirements for attribute-based instance type selection. Instead of
specifying exact instance types, you define requirements such as vCPU count, memory
size, network performance, and accelerator specifications. Amazon ECS automatically
selects Amazon EC2 instance types that match these requirements, providing flexibility
and helping to mitigate capacity constraints.

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
  RequireHibernateSupport: Boolean
  SpotMaxPricePercentageOverLowestPrice: Integer
  TotalLocalStorageGB:
    TotalLocalStorageGBRequest
  VCpuCount:
    VCpuCountRangeRequest

```

## Properties

`AcceleratorCount`

The minimum and maximum number of accelerators for the instance types. This is used
when you need instances with specific numbers of GPUs or other accelerators.

_Required_: No

_Type_: [AcceleratorCountRequest](aws-properties-ecs-capacityprovider-acceleratorcountrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorManufacturers`

The accelerator manufacturers to include. You can specify `nvidia`,
`amd`, `amazon-web-services`, or `xilinx` depending
on your accelerator requirements.

_Required_: No

_Type_: Array of String

_Allowed values_: `amazon-web-services | amd | habana | nvidia | xilinx`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorNames`

The specific accelerator names to include. For example, you can specify
`a100`, `v100`, `k80`, or other specific
accelerator models.

_Required_: No

_Type_: Array of String

_Allowed values_: `a10g | a100 | h100 | inferentia | k520 | k80 | m60 | radeon-pro-v520 | t4 | t4g | vu9p | v100 | l40s`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorTotalMemoryMiB`

The minimum and maximum total accelerator memory in mebibytes (MiB). This is important
for GPU workloads that require specific amounts of video memory.

_Required_: No

_Type_: [AcceleratorTotalMemoryMiBRequest](aws-properties-ecs-capacityprovider-acceleratortotalmemorymibrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceleratorTypes`

The accelerator types to include. You can specify `gpu` for graphics
processing units, `fpga` for field programmable gate arrays, or
`inference` for machine learning inference accelerators.

_Required_: No

_Type_: Array of String

_Allowed values_: `gpu | fpga | inference`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedInstanceTypes`

The instance types to include in the selection. When specified, Amazon ECS only
considers these instance types, subject to the other requirements specified.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BareMetal`

Indicates whether to include bare metal instance types. Set to `included`
to allow bare metal instances, `excluded` to exclude them, or
`required` to use only bare metal instances.

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaselineEbsBandwidthMbps`

The minimum and maximum baseline Amazon EBS bandwidth in megabits per second (Mbps).
This is important for workloads with high storage I/O requirements.

_Required_: No

_Type_: [BaselineEbsBandwidthMbpsRequest](aws-properties-ecs-capacityprovider-baselineebsbandwidthmbpsrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BurstablePerformance`

Indicates whether to include burstable performance instance types (T2, T3, T3a, T4g).
Set to `included` to allow burstable instances, `excluded` to
exclude them, or `required` to use only burstable instances.

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CpuManufacturers`

The CPU manufacturers to include or exclude. You can specify `intel`,
`amd`, or `amazon-web-services` to control which CPU types are
used for your workloads.

_Required_: No

_Type_: Array of String

_Allowed values_: `intel | amd | amazon-web-services`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedInstanceTypes`

The instance types to exclude from selection. Use this to prevent Amazon ECS from
selecting specific instance types that may not be suitable for your workloads.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceGenerations`

The instance generations to include. You can specify `current` to use the
latest generation instances, or `previous` to include previous generation
instances for cost optimization.

_Required_: No

_Type_: Array of String

_Allowed values_: `current | previous`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalStorage`

Indicates whether to include instance types with local storage. Set to
`included` to allow local storage, `excluded` to exclude it,
or `required` to use only instances with local storage.

_Required_: No

_Type_: String

_Allowed values_: `included | required | excluded`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalStorageTypes`

The local storage types to include. You can specify `hdd` for hard disk
drives, `ssd` for solid state drives, or both.

_Required_: No

_Type_: Array of String

_Allowed values_: `hdd | ssd`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`

The maximum price for Spot instances as a percentage of the optimal On-Demand price.
This provides more precise cost control for Spot instance selection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryGiBPerVCpu`

The minimum and maximum amount of memory per vCPU in gibibytes (GiB). This helps
ensure that instance types have the appropriate memory-to-CPU ratio for your
workloads.

_Required_: No

_Type_: [MemoryGiBPerVCpuRequest](aws-properties-ecs-capacityprovider-memorygibpervcpurequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryMiB`

The minimum and maximum amount of memory in mebibytes (MiB) for the instance types.
Amazon ECS selects instance types that have memory within this range.

_Required_: Yes

_Type_: [MemoryMiBRequest](aws-properties-ecs-capacityprovider-memorymibrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkBandwidthGbps`

The minimum and maximum network bandwidth in gigabits per second (Gbps). This is
crucial for network-intensive workloads that require high throughput.

_Required_: No

_Type_: [NetworkBandwidthGbpsRequest](aws-properties-ecs-capacityprovider-networkbandwidthgbpsrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceCount`

The minimum and maximum number of network interfaces for the instance types. This is
useful for workloads that require multiple network interfaces.

_Required_: No

_Type_: [NetworkInterfaceCountRequest](aws-properties-ecs-capacityprovider-networkinterfacecountrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnDemandMaxPricePercentageOverLowestPrice`

The price protection threshold for On-Demand Instances, as a percentage higher than an
identified On-Demand price. The identified On-Demand price is the price of the lowest
priced current generation C, M, or R instance type with your specified attributes. If no
current generation C, M, or R instance type matches your attributes, then the identified
price is from either the lowest priced current generation instance types or, failing
that, the lowest priced previous generation instance types that match your attributes.
When Amazon ECS selects instance types with your attributes, we will exclude instance
types whose price exceeds your specified threshold.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireHibernateSupport`

Indicates whether the instance types must support hibernation. When set to
`true`, only instance types that support hibernation are selected.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotMaxPricePercentageOverLowestPrice`

The maximum price for Spot instances as a percentage over the lowest priced On-Demand
instance. This helps control Spot instance costs while maintaining access to
capacity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalLocalStorageGB`

The minimum and maximum total local storage in gigabytes (GB) for instance types with
local storage.

_Required_: No

_Type_: [TotalLocalStorageGBRequest](aws-properties-ecs-capacityprovider-totallocalstoragegbrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VCpuCount`

The minimum and maximum number of vCPUs for the instance types. Amazon ECS selects
instance types that have vCPU counts within this range.

_Required_: Yes

_Type_: [VCpuCountRangeRequest](aws-properties-ecs-capacityprovider-vcpucountrangerequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceLaunchTemplate

ManagedInstancesLocalStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
