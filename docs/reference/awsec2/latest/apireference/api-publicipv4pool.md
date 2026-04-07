# PublicIpv4Pool

Describes an IPv4 address pool.

## Contents

**description**

A description of the address pool.

Type: String

Required: No

**networkBorderGroup**

The name of the location from which the address pool is advertised.
A network border group is a unique set of Availability Zones or Local Zones
from where AWS advertises public IP addresses.

Type: String

Required: No

**PoolAddressRangeSet.N**

The address ranges.

Type: Array of [PublicIpv4PoolRange](api-publicipv4poolrange.md) objects

Required: No

**poolId**

The ID of the address pool.

Type: String

Required: No

**TagSet.N**

Any tags for the address pool.

Type: Array of [Tag](api-tag.md) objects

Required: No

**totalAddressCount**

The total number of addresses.

Type: Integer

Required: No

**totalAvailableAddressCount**

The total number of available addresses.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/publicipv4pool.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/publicipv4pool.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/publicipv4pool.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PublicIpDnsNameOptions

PublicIpv4PoolRange
