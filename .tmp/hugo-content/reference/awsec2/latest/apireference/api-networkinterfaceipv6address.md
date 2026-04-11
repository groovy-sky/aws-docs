# NetworkInterfaceIpv6Address

Describes an IPv6 address associated with a network interface.

## Contents

**ipv6Address**

The IPv6 address.

Type: String

Required: No

**isPrimaryIpv6**

Determines if an IPv6 address associated with a network interface is the primary IPv6
address. When you enable an IPv6 GUA address to be a primary IPv6, the first IPv6 GUA
will be made the primary IPv6 address until the instance is terminated or the network
interface is detached. For more information, see [ModifyNetworkInterfaceAttribute](api-modifynetworkinterfaceattribute.md).

Type: Boolean

Required: No

**publicIpv6DnsName**

An IPv6-enabled public hostname for a network interface. Requests from within the VPC or from the internet resolve to the IPv6 GUA of the network interface. For more information, see [EC2 instance hostnames, DNS names, and domains](../../../../services/ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon EC2 User Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinterfaceipv6address.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinterfaceipv6address.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinterfaceipv6address.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInterfaceCountRequest

NetworkInterfacePermission
