This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet SpotMaintenanceStrategies

The strategies for managing your Spot Instances that are at an elevated risk of being
interrupted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityRebalance" : SpotCapacityRebalance
}

```

### YAML

```yaml

  CapacityRebalance:
    SpotCapacityRebalance

```

## Properties

`CapacityRebalance`

The Spot Instance replacement strategy to use when Amazon EC2 emits a signal that your
Spot Instance is at an elevated risk of being interrupted. For more information, see
[Capacity\
rebalancing](../../../ec2/latest/userguide/spot-fleet-capacity-rebalance.md) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: [SpotCapacityRebalance](aws-properties-ec2-spotfleet-spotcapacityrebalance.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpotFleetTagSpecification

SpotPlacement

All content copied from https://docs.aws.amazon.com/.
