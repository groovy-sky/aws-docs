This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster SpotProvisioningSpecification

`SpotProvisioningSpecification` is a subproperty of the `InstanceFleetProvisioningSpecifications` property type. `SpotProvisioningSpecification` determines the launch specification for Spot instances in the instance fleet, which includes the defined duration and provisioning timeout behavior.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "BlockDurationMinutes" : Integer,
  "TimeoutAction" : String,
  "TimeoutDurationMinutes" : Integer
}

```

### YAML

```yaml

  AllocationStrategy: String
  BlockDurationMinutes: Integer
  TimeoutAction: String
  TimeoutDurationMinutes: Integer

```

## Properties

`AllocationStrategy`

Specifies one of the following strategies to launch Spot Instance fleets:
`capacity-optimized`, `price-capacity-optimized`, `lowest-price`, or
`diversified`, and `capacity-optimized-prioritized`. For more information on the provisioning strategies, see [Allocation strategies for Spot Instances](../../../ec2/latest/userguide/ec2-fleet-allocation-strategy.md) in the _Amazon EC2 User Guide for Linux Instances_.

###### Note

When you launch a Spot Instance fleet with the old console, it automatically launches with the `capacity-optimized` strategy. You can't change the allocation strategy from the old console.

_Required_: No

_Type_: String

_Allowed values_: `capacity-optimized | price-capacity-optimized | lowest-price | diversified | capacity-optimized-prioritized`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockDurationMinutes`

The defined duration for Spot Instances (also known as Spot blocks) in minutes. When
specified, the Spot Instance does not terminate before the defined duration expires, and
defined duration pricing for Spot Instances applies. Valid values are 60, 120, 180, 240,
300, or 360. The duration period starts as soon as a Spot Instance receives its instance
ID. At the end of the duration, Amazon EC2 marks the Spot Instance for termination
and provides a Spot Instance termination notice, which gives the instance a two-minute
warning before it terminates.

###### Note

Spot Instances with a defined duration (also known as Spot blocks) are no longer
available to new customers from July 1, 2021. For customers who have previously used the
feature, we will continue to support Spot Instances with a defined duration until
December 31, 2022.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutAction`

The action to take when `TargetSpotCapacity` has not been fulfilled when the
`TimeoutDurationMinutes` has expired; that is, when all Spot Instances could
not be provisioned within the Spot provisioning timeout. Valid values are
`TERMINATE_CLUSTER` and `SWITCH_TO_ON_DEMAND`. SWITCH\_TO\_ON\_DEMAND
specifies that if no Spot Instances are available, On-Demand Instances should be
provisioned to fulfill any remaining Spot capacity.

_Required_: Yes

_Type_: String

_Allowed values_: `SWITCH_TO_ON_DEMAND | TERMINATE_CLUSTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutDurationMinutes`

The Spot provisioning timeout period in minutes. If Spot Instances are not provisioned
within this time period, the `TimeOutAction` is taken. Minimum value is 5 and
maximum value is 1440. The timeout applies only during initial provisioning, when the
cluster is first created.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SimpleScalingPolicyConfiguration

SpotResizingSpecification

All content copied from https://docs.aws.amazon.com/.
