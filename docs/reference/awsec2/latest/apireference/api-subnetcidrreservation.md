# SubnetCidrReservation

Describes a subnet CIDR reservation.

## Contents

**cidr**

The CIDR that has been reserved.

Type: String

Required: No

**description**

The description assigned to the subnet CIDR reservation.

Type: String

Required: No

**ownerId**

The ID of the account that owns the subnet CIDR reservation.

Type: String

Required: No

**reservationType**

The type of reservation.

Type: String

Valid Values: `prefix | explicit`

Required: No

**subnetCidrReservationId**

The ID of the subnet CIDR reservation.

Type: String

Required: No

**subnetId**

The ID of the subnet.

Type: String

Required: No

**TagSet.N**

The tags assigned to the subnet CIDR reservation.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/subnetcidrreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/subnetcidrreservation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/subnetcidrreservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SubnetCidrBlockState

SubnetConfiguration
