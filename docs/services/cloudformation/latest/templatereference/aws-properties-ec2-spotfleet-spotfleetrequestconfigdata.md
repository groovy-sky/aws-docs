This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet SpotFleetRequestConfigData

Specifies the configuration of a Spot Fleet request. For more information, see [Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the _Amazon EC2 User Guide_.

You must specify either `LaunchSpecifications` or
`LaunchTemplateConfigs`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "Context" : String,
  "ExcessCapacityTerminationPolicy" : String,
  "IamFleetRole" : String,
  "InstanceInterruptionBehavior" : String,
  "InstancePoolsToUseCount" : Integer,
  "LaunchSpecifications" : [ SpotFleetLaunchSpecification, ... ],
  "LaunchTemplateConfigs" : [ LaunchTemplateConfig, ... ],
  "LoadBalancersConfig" : LoadBalancersConfig,
  "OnDemandAllocationStrategy" : String,
  "OnDemandMaxTotalPrice" : String,
  "OnDemandTargetCapacity" : Integer,
  "ReplaceUnhealthyInstances" : Boolean,
  "SpotMaintenanceStrategies" : SpotMaintenanceStrategies,
  "SpotMaxTotalPrice" : String,
  "SpotPrice" : String,
  "TagSpecifications" : [ SpotFleetTagSpecification, ... ],
  "TargetCapacity" : Integer,
  "TargetCapacityUnitType" : String,
  "TerminateInstancesWithExpiration" : Boolean,
  "Type" : String,
  "ValidFrom" : String,
  "ValidUntil" : String
}

```

### YAML

```yaml

  AllocationStrategy: String
  Context: String
  ExcessCapacityTerminationPolicy: String
  IamFleetRole: String
  InstanceInterruptionBehavior: String
  InstancePoolsToUseCount: Integer
  LaunchSpecifications:
    - SpotFleetLaunchSpecification
  LaunchTemplateConfigs:
    - LaunchTemplateConfig
  LoadBalancersConfig:
    LoadBalancersConfig
  OnDemandAllocationStrategy: String
  OnDemandMaxTotalPrice: String
  OnDemandTargetCapacity: Integer
  ReplaceUnhealthyInstances: Boolean
  SpotMaintenanceStrategies:
    SpotMaintenanceStrategies
  SpotMaxTotalPrice: String
  SpotPrice: String
  TagSpecifications:
    - SpotFleetTagSpecification
  TargetCapacity: Integer
  TargetCapacityUnitType: String
  TerminateInstancesWithExpiration: Boolean
  Type: String
  ValidFrom: String
  ValidUntil: String

```

## Properties

`AllocationStrategy`

The strategy that determines how to allocate the target Spot Instance capacity across the Spot Instance
pools specified by the Spot Fleet launch configuration. For more information, see [Allocation\
strategies for Spot Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-allocation-strategy.html) in the _Amazon EC2 User Guide_.

priceCapacityOptimized (recommended)

Spot Fleet identifies the pools with
the highest capacity availability for the number of instances that are launching. This means
that we will request Spot Instances from the pools that we believe have the lowest chance of interruption
in the near term. Spot Fleet then requests Spot Instances from the lowest priced of these pools.

capacityOptimized

Spot Fleet identifies the pools with
the highest capacity availability for the number of instances that are launching. This means
that we will request Spot Instances from the pools that we believe have the lowest chance of interruption
in the near term. To give certain
instance types a higher chance of launching first, use
`capacityOptimizedPrioritized`. Set a priority for each instance type by
using the `Priority` parameter for `LaunchTemplateOverrides`. You can
assign the same priority to different `LaunchTemplateOverrides`. EC2 implements
the priorities on a best-effort basis, but optimizes for capacity first.
`capacityOptimizedPrioritized` is supported only if your Spot Fleet uses a
launch template. Note that if the `OnDemandAllocationStrategy` is set to
`prioritized`, the same priority is applied when fulfilling On-Demand
capacity.

diversified

Spot Fleet requests instances from all of the Spot Instance pools that you
specify.

lowestPrice (not recommended)

###### Important

We don't recommend the `lowestPrice` allocation strategy because
it has the highest risk of interruption for your Spot Instances.

Spot Fleet requests instances from the lowest priced Spot Instance pool that has available
capacity. If the lowest priced pool doesn't have available capacity, the Spot Instances
come from the next lowest priced pool that has available capacity. If a pool runs
out of capacity before fulfilling your desired capacity, Spot Fleet will continue to
fulfill your request by drawing from the next lowest priced pool. To ensure that
your desired capacity is met, you might receive Spot Instances from several pools. Because
this strategy only considers instance price and not capacity availability, it
might lead to high interruption rates.

Default: `lowestPrice`

_Required_: No

_Type_: String

_Allowed values_: `capacityOptimized | capacityOptimizedPrioritized | diversified | lowestPrice | priceCapacityOptimized`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Context`

Reserved.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcessCapacityTerminationPolicy`

Indicates whether running Spot Instances should be terminated if you decrease the
target capacity of the Spot Fleet request below the current size of the Spot
Fleet.

Supported only for fleets of type `maintain`.

_Required_: No

_Type_: String

_Allowed values_: `Default | NoTermination | default | noTermination`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamFleetRole`

The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that grants the
Spot Fleet the permission to request, launch, terminate, and tag instances on your behalf.
For more information, see [Spot\
Fleet Prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html#spot-fleet-prerequisites) in the _Amazon EC2 User Guide_.
Spot Fleet can terminate Spot Instances on your behalf when you
cancel its Spot Fleet request or when the Spot Fleet request expires, if you set
`TerminateInstancesWithExpiration`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceInterruptionBehavior`

The behavior when a Spot Instance is interrupted. The default is
`terminate`.

_Required_: No

_Type_: String

_Allowed values_: `hibernate | stop | terminate`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstancePoolsToUseCount`

The number of Spot pools across which to allocate your target Spot capacity. Valid
only when Spot **AllocationStrategy** is set to
`lowest-price`. Spot Fleet selects the cheapest Spot pools and evenly
allocates your target Spot capacity across the number of Spot pools that you
specify.

Note that Spot Fleet attempts to draw Spot Instances from the number of pools that you specify on a
best effort basis. If a pool runs out of Spot capacity before fulfilling your target
capacity, Spot Fleet will continue to fulfill your request by drawing from the next cheapest
pool. To ensure that your target capacity is met, you might receive Spot Instances from more than
the number of pools that you specified. Similarly, if most of the pools have no Spot
capacity, you might receive your full target capacity from fewer than the number of
pools that you specified.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchSpecifications`

The launch specifications for the Spot Fleet request. If you specify
`LaunchSpecifications`, you can't specify
`LaunchTemplateConfigs`.

_Required_: Conditional

_Type_: Array of [SpotFleetLaunchSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchTemplateConfigs`

The launch template and overrides. If you specify `LaunchTemplateConfigs`,
you can't specify `LaunchSpecifications`.

_Required_: Conditional

_Type_: Array of [LaunchTemplateConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-launchtemplateconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoadBalancersConfig`

One or more Classic Load Balancers and target groups to attach to the Spot Fleet
request. Spot Fleet registers the running Spot Instances with the specified Classic Load
Balancers and target groups.

With Network Load Balancers, Spot Fleet cannot register instances that have the
following instance types: C1, CC1, CC2, CG1, CG2, CR1, CS1, G1, G2, HI1, HS1, M1, M2,
M3, and T1.

_Required_: No

_Type_: [LoadBalancersConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-loadbalancersconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandAllocationStrategy`

The order of the launch template overrides to use in fulfilling On-Demand capacity. If
you specify `lowestPrice`, Spot Fleet uses price to determine the order, launching
the lowest price first. If you specify `prioritized`, Spot Fleet uses the priority
that you assign to each Spot Fleet launch template override, launching the highest priority
first. If you do not specify a value, Spot Fleet defaults to `lowestPrice`.

_Required_: No

_Type_: String

_Allowed values_: `lowestPrice | prioritized`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandMaxTotalPrice`

The maximum amount per hour for On-Demand Instances that you're willing to pay. You
can use the `onDemandMaxTotalPrice` parameter, the
`spotMaxTotalPrice` parameter, or both parameters to ensure that your
fleet cost does not exceed your budget. If you set a maximum price per hour for the
On-Demand Instances and Spot Instances in your request, Spot Fleet will launch instances until it reaches the
maximum amount you're willing to pay. When the maximum amount you're willing to pay is
reached, the fleet stops launching instances even if it hasn’t met the target
capacity.

###### Note

If your fleet includes T instances that are configured as `unlimited`,
and if their average CPU usage exceeds the baseline utilization, you will incur a charge
for surplus credits. The `onDemandMaxTotalPrice` does not account for surplus
credits, and, if you use surplus credits, your final cost might be higher than what you
specified for `onDemandMaxTotalPrice`. For more information, see [Surplus credits can incur charges](../../../ec2/latest/userguide/burstable-performance-instances-unlimited-mode-concepts.md#unlimited-mode-surplus-credits) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandTargetCapacity`

The number of On-Demand units to request. You can choose to set the target capacity in
terms of instances or a performance characteristic that is important to your application
workload, such as vCPUs, memory, or I/O. If the request type is `maintain`,
you can specify a target capacity of 0 and add capacity later.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplaceUnhealthyInstances`

Indicates whether Spot Fleet should replace unhealthy instances.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotMaintenanceStrategies`

The strategies for managing your Spot Instances that are at an elevated risk of being
interrupted.

_Required_: No

_Type_: [SpotMaintenanceStrategies](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-spotmaintenancestrategies.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotMaxTotalPrice`

The maximum amount per hour for Spot Instances that you're willing to pay. You can use
the `spotMaxTotalPrice` parameter, the `onDemandMaxTotalPrice`
parameter, or both parameters to ensure that your fleet cost does not exceed your budget.
If you set a maximum price per hour for the On-Demand Instances and Spot Instances in your request, Spot Fleet will
launch instances until it reaches the maximum amount you're willing to pay. When the
maximum amount you're willing to pay is reached, the fleet stops launching instances even
if it hasn’t met the target capacity.

###### Note

If your fleet includes T instances that are configured as `unlimited`,
and if their average CPU usage exceeds the baseline utilization, you will incur a charge
for surplus credits. The `spotMaxTotalPrice` does not account for surplus
credits, and, if you use surplus credits, your final cost might be higher than what you
specified for `spotMaxTotalPrice`. For more information, see [Surplus credits can incur charges](../../../ec2/latest/userguide/burstable-performance-instances-unlimited-mode-concepts.md#unlimited-mode-surplus-credits) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotPrice`

The maximum price per unit hour that you are willing to pay for a Spot Instance. We do not recommend
using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.

###### Important

If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify this parameter.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagSpecifications`

The key-value pair for tagging the Spot Fleet request on creation. The value for
`ResourceType` must be `spot-fleet-request`, otherwise the
Spot Fleet request fails. To tag instances at launch, specify the tags in the [launch\
template](../../../ec2/latest/userguide/ec2-launch-templates.md#create-launch-template) (valid only if you use `LaunchTemplateConfigs`) or in
the `
                            SpotFleetTagSpecification
                        ` (valid only if you use
`LaunchSpecifications`). For information about tagging after launch, see
[Tag your resources](../../../ec2/latest/userguide/using-tags.md#tag-resources).

_Required_: No

_Type_: Array of [SpotFleetTagSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-spotfleettagspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetCapacity`

The number of units to request for the Spot Fleet. You can choose to set the target
capacity in terms of instances or a performance characteristic that is important to your
application workload, such as vCPUs, memory, or I/O. If the request type is
`maintain`, you can specify a target capacity of 0 and add capacity
later.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetCapacityUnitType`

The unit for the target capacity. You can specify this parameter only when
using attribute-based instance type selection.

Default: `units` (the number of instances)

_Required_: No

_Type_: String

_Allowed values_: `vcpu | memory-mib | units`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminateInstancesWithExpiration`

Indicates whether running Spot Instances are terminated when the Spot Fleet request
expires.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of request. Indicates whether the Spot Fleet only requests the target
capacity or also attempts to maintain it. When this value is `request`, the
Spot Fleet only places the required requests. It does not attempt to replenish Spot
Instances if capacity is diminished, nor does it submit requests in alternative Spot
pools if capacity is not available. When this value is `maintain`, the Spot
Fleet maintains the target capacity. The Spot Fleet places the required requests to meet
capacity and automatically replenishes any interrupted instances. Default:
`maintain`. `instant` is listed but is not used by Spot
Fleet.

_Required_: No

_Type_: String

_Allowed values_: `maintain | request`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidFrom`

The start date and time of the request, in UTC format
( _YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).
By default, Amazon EC2 starts fulfilling the request immediately.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidUntil`

The end date and time of the request, in UTC format
( _YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).
After the end date and time, no new Spot Instance requests are placed or able to fulfill
the request. If no value is specified, the Spot Fleet request remains until you cancel
it.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SpotFleetMonitoring

SpotFleetTagSpecification
