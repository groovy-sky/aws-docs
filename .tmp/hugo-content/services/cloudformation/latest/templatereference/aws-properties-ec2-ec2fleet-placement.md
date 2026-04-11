This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet Placement

Describes the placement of an instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Affinity" : String,
  "AvailabilityZone" : String,
  "GroupName" : String,
  "HostId" : String,
  "HostResourceGroupArn" : String,
  "PartitionNumber" : Integer,
  "SpreadDomain" : String,
  "Tenancy" : String
}

```

### YAML

```yaml

  Affinity: String
  AvailabilityZone: String
  GroupName: String
  HostId: String
  HostResourceGroupArn: String
  PartitionNumber: Integer
  SpreadDomain: String
  Tenancy: String

```

## Properties

`Affinity`

The affinity setting for the instance on the Dedicated Host.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md) or [ImportInstance](../../../../reference/awsec2/latest/apireference/api-importinstance.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZone`

The Availability Zone of the instance.

On input, you can specify `AvailabilityZone` or `AvailabilityZoneId`,
but not both. If you specify neither one, Amazon EC2 automatically selects an Availability Zone
for you.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupName`

The name of the placement group that the instance is in.

On input, you can specify `GroupId` or `GroupName`,
but not both.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostId`

The ID of the Dedicated Host on which the instance resides.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md) or [ImportInstance](../../../../reference/awsec2/latest/apireference/api-importinstance.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostResourceGroupArn`

The ARN of the host resource group in which to launch the instances.

On input, if you specify this parameter, either omit the **Tenancy** parameter or set it to `host`.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartitionNumber`

The number of the partition that the instance is in. Valid only if the placement group
strategy is set to `partition`.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md).

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpreadDomain`

Reserved for future use.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tenancy`

The tenancy of the instance. An instance with a
tenancy of `dedicated` runs on single-tenant hardware.

This parameter is not supported for [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md). The
`host` tenancy is not supported for [ImportInstance](../../../../reference/awsec2/latest/apireference/api-importinstance.md) or
for T3 instances that are configured for the `unlimited` CPU credit
option.

_Required_: No

_Type_: String

_Allowed values_: `default | dedicated | host`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PerformanceFactorReferenceRequest

ReservedCapacityOptionsRequest

All content copied from https://docs.aws.amazon.com/.
