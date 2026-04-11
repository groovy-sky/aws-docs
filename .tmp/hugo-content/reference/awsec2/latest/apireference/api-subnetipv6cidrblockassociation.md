# SubnetIpv6CidrBlockAssociation

Describes an association between a subnet and an IPv6 CIDR block.

## Contents

**associationId**

The ID of the association.

Type: String

Required: No

**ipSource**

The source that allocated the IP address space. `byoip` or `amazon` indicates public IP address space allocated by Amazon or space that you have allocated with Bring your own IP (BYOIP). `none` indicates private space.

Type: String

Valid Values: `amazon | byoip | none`

Required: No

**ipv6AddressAttribute**

Public IPv6 addresses are those advertised on the internet from AWS. Private IP addresses are not and cannot be advertised on the internet from AWS.

Type: String

Valid Values: `public | private`

Required: No

**ipv6CidrBlock**

The IPv6 CIDR block.

Type: String

Required: No

**ipv6CidrBlockState**

The state of the CIDR block.

Type: [SubnetCidrBlockState](api-subnetcidrblockstate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/subnetipv6cidrblockassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/subnetipv6cidrblockassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/subnetipv6cidrblockassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SubnetIpPrefixes

Subscription
