This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet SpotPlacement

Describes Spot Instance placement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "AvailabilityZoneId" : String,
  "GroupName" : String,
  "Tenancy" : String
}

```

### YAML

```yaml

  AvailabilityZone: String
  AvailabilityZoneId: String
  GroupName: String
  Tenancy: String

```

## Properties

`AvailabilityZone`

The Availability Zone.

To specify multiple Availability Zones, separate them using commas; for example,
"us-west-2a, us-west-2b".

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneId`

The ID of the Availability Zone. For example, `use2-az1`.

\[Spot Fleet only\] To specify multiple Availability Zones, separate them using commas;
for example, " `use2-az1`, `use2-bz1`".

Either `AvailabilityZone` or `AvailabilityZoneId` must be
specified in the request, but not both.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupName`

The name of the placement group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tenancy`

The tenancy of the instance (if the instance is running in a VPC). An instance with a
tenancy of `dedicated` runs on single-tenant hardware. The `host`
tenancy is not supported for Spot Instances.

_Required_: No

_Type_: String

_Allowed values_: `dedicated | default | host`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpotMaintenanceStrategies

Tag

All content copied from https://docs.aws.amazon.com/.
