# Placement

Describes the placement of an instance.

## Contents

**Affinity** (request), **affinity** (response)

The affinity setting for the instance on the Dedicated Host.

This parameter is not supported for [CreateFleet](api-createfleet.md) or [ImportInstance](api-importinstance.md).

Type: String

Required: No

**AvailabilityZone** (request), **availabilityZone** (response)

The Availability Zone of the instance.

On input, you can specify `AvailabilityZone` or `AvailabilityZoneId`,
but not both. If you specify neither one, Amazon EC2 automatically selects an Availability Zone
for you.

This parameter is not supported for [CreateFleet](api-createfleet.md).

Type: String

Required: No

**AvailabilityZoneId** (request), **availabilityZoneId** (response)

The ID of the Availability Zone of the instance.

On input, you can specify `AvailabilityZone` or `AvailabilityZoneId`,
but not both. If you specify neither one, Amazon EC2 automatically selects an Availability Zone
for you.

This parameter is not supported for [CreateFleet](api-createfleet.md).

Type: String

Required: No

**GroupId** (request), **groupId** (response)

The ID of the placement group that the instance is in.

On input, you can specify `GroupId` or `GroupName`,
but not both.

Type: String

Required: No

**GroupName** (request), **groupName** (response)

The name of the placement group that the instance is in.

On input, you can specify `GroupId` or `GroupName`,
but not both.

Type: String

Required: No

**HostId** (request), **hostId** (response)

The ID of the Dedicated Host on which the instance resides.

This parameter is not supported for [CreateFleet](api-createfleet.md) or [ImportInstance](api-importinstance.md).

Type: String

Required: No

**HostResourceGroupArn** (request), **hostResourceGroupArn** (response)

The ARN of the host resource group in which to launch the instances.

On input, if you specify this parameter, either omit the **Tenancy** parameter or set it to `host`.

This parameter is not supported for [CreateFleet](api-createfleet.md).

Type: String

Required: No

**PartitionNumber** (request), **partitionNumber** (response)

The number of the partition that the instance is in. Valid only if the placement group
strategy is set to `partition`.

This parameter is not supported for [CreateFleet](api-createfleet.md).

Type: Integer

Required: No

**SpreadDomain** (request), **spreadDomain** (response)

Reserved for future use.

Type: String

Required: No

**Tenancy** (request), **tenancy** (response)

The tenancy of the instance. An instance with a
tenancy of `dedicated` runs on single-tenant hardware.

This parameter is not supported for [CreateFleet](api-createfleet.md). The
`host` tenancy is not supported for [ImportInstance](api-importinstance.md) or
for T3 instances that are configured for the `unlimited` CPU credit
option.

Type: String

Valid Values: `default | dedicated | host`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/placement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/placement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/placement.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Phase2IntegrityAlgorithmsRequestListValue

PlacementGroup
