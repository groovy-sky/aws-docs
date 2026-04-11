# SubnetConfiguration

Describes the configuration of a subnet for a VPC endpoint.

## Contents

**Ipv4**

The IPv4 address to assign to the endpoint network interface in the subnet. You must provide
an IPv4 address if the VPC endpoint supports IPv4.

If you specify an IPv4 address when modifying a VPC endpoint, we replace the existing
endpoint network interface with a new endpoint network interface with this IP address.
This process temporarily disconnects the subnet and the VPC endpoint.

Type: String

Required: No

**Ipv6**

The IPv6 address to assign to the endpoint network interface in the subnet. You must provide
an IPv6 address if the VPC endpoint supports IPv6.

If you specify an IPv6 address when modifying a VPC endpoint, we replace the existing
endpoint network interface with a new endpoint network interface with this IP address.
This process temporarily disconnects the subnet and the VPC endpoint.

Type: String

Required: No

**SubnetId**

The ID of the subnet.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/subnetconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/subnetconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/subnetconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SubnetCidrReservation

SubnetIpPrefixes
