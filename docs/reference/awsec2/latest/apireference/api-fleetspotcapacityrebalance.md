# FleetSpotCapacityRebalance

The strategy to use when Amazon EC2 emits a signal that your Spot Instance is at an
elevated risk of being interrupted.

## Contents

**replacementStrategy**

The replacement strategy to use. Only available for fleets of type
`maintain`.

`launch` \- EC2 Fleet launches a new replacement Spot Instance when a
rebalance notification is emitted for an existing Spot Instance in the fleet. EC2 Fleet
does not terminate the instances that receive a rebalance notification. You can terminate
the old instances, or you can leave them running. You are charged for all instances while
they are running.

`launch-before-terminate` \- EC2 Fleet launches a new replacement Spot
Instance when a rebalance notification is emitted for an existing Spot Instance in the
fleet, and then, after a delay that you specify (in `TerminationDelay`),
terminates the instances that received a rebalance notification.

Type: String

Valid Values: `launch | launch-before-terminate`

Required: No

**terminationDelay**

The amount of time (in seconds) that Amazon EC2 waits before terminating the old Spot
Instance after launching a new replacement Spot Instance.

Required when `ReplacementStrategy` is set to `launch-before-terminate`.

Not valid when `ReplacementStrategy` is set to `launch`.

Valid values: Minimum value of `120` seconds. Maximum value of `7200` seconds.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/fleetspotcapacityrebalance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/fleetspotcapacityrebalance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/fleetspotcapacityrebalance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FleetLaunchTemplateSpecificationRequest

FleetSpotCapacityRebalanceRequest
