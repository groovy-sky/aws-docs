This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet MaintenanceStrategies

The strategies for managing your Spot Instances that are at an elevated risk of being
interrupted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityRebalance" : CapacityRebalance
}

```

### YAML

```yaml

  CapacityRebalance:
    CapacityRebalance

```

## Properties

`CapacityRebalance`

The strategy to use when Amazon EC2 emits a signal that your Spot Instance is at an
elevated risk of being interrupted.

_Required_: No

_Type_: [CapacityRebalance](aws-properties-ec2-ec2fleet-capacityrebalance.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceRequirementsRequest

MemoryGiBPerVCpuRequest

All content copied from https://docs.aws.amazon.com/.
