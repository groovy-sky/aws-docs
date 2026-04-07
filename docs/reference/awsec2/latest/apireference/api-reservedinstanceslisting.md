# ReservedInstancesListing

Describes a Reserved Instance listing.

## Contents

**clientToken**

A unique, case-sensitive key supplied by the client to ensure that the request is
idempotent. For more information, see [Ensuring\
Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**createDate**

The time the listing was created.

Type: Timestamp

Required: No

**InstanceCounts.N**

The number of instances in this state.

Type: Array of [InstanceCount](api-instancecount.md) objects

Required: No

**PriceSchedules.N**

The price of the Reserved Instance listing.

Type: Array of [PriceSchedule](api-priceschedule.md) objects

Required: No

**reservedInstancesId**

The ID of the Reserved Instance.

Type: String

Required: No

**reservedInstancesListingId**

The ID of the Reserved Instance listing.

Type: String

Required: No

**status**

The status of the Reserved Instance listing.

Type: String

Valid Values: `active | pending | cancelled | closed`

Required: No

**statusMessage**

The reason for the current status of the Reserved Instance listing. The response can be
blank.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the resource.

Type: Array of [Tag](api-tag.md) objects

Required: No

**updateDate**

The last modified timestamp of the listing.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/reservedinstanceslisting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/reservedinstanceslisting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/reservedinstanceslisting.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReservedInstancesId

ReservedInstancesModification
