This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet SpotCapacityRebalance

The Spot Instance replacement strategy to use when Amazon EC2 emits a signal that your
Spot Instance is at an elevated risk of being interrupted. For more information, see
[Capacity\
rebalancing](../../../ec2/latest/userguide/spot-fleet-capacity-rebalance.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplacementStrategy" : String,
  "TerminationDelay" : Integer
}

```

### YAML

```yaml

  ReplacementStrategy: String
  TerminationDelay: Integer

```

## Properties

`ReplacementStrategy`

The replacement strategy to use. Only available for fleets of type
`maintain`.

`launch` \- Spot Fleet launches a new replacement Spot Instance when a
rebalance notification is emitted for an existing Spot Instance in the fleet. Spot Fleet
does not terminate the instances that receive a rebalance notification. You can
terminate the old instances, or you can leave them running. You are charged for all
instances while they are running.

`launch-before-terminate` \- Spot Fleet launches a new replacement Spot
Instance when a rebalance notification is emitted for an existing Spot Instance in the
fleet, and then, after a delay that you specify (in `TerminationDelay`),
terminates the instances that received a rebalance notification.

_Required_: No

_Type_: String

_Allowed values_: `launch | launch-before-terminate`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TerminationDelay`

The amount of time (in seconds) that Amazon EC2 waits before terminating the old Spot
Instance after launching a new replacement Spot Instance.

Required when `ReplacementStrategy` is set to `launch-before-terminate`.

Not valid when `ReplacementStrategy` is set to `launch`.

Valid values: Minimum value of `120` seconds. Maximum value of `7200` seconds.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivateIpAddressSpecification

SpotFleetLaunchSpecification

All content copied from https://docs.aws.amazon.com/.
