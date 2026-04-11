# InstanceIpv6Address

Describes an IPv6 address.

## Contents

**Ipv6Address** (request), **ipv6Address** (response)

The IPv6 address.

Type: String

Required: No

**IsPrimaryIpv6** (request), **isPrimaryIpv6** (response)

Determines if an IPv6 address associated with a network interface is the primary IPv6 address. When you enable an IPv6 GUA address to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address until the instance is terminated or the network interface is detached.
For more information, see [RunInstances](api-runinstances.md).

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceipv6address.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceipv6address.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceipv6address.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceIpv4Prefix

InstanceIpv6AddressRequest
