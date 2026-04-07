# InstanceNetworkInterface

Describes a network interface.

## Contents

**association**

The association information for an Elastic IPv4 associated with the network
interface.

Type: [InstanceNetworkInterfaceAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceNetworkInterfaceAssociation.html) object

Required: No

**attachment**

The network interface attachment.

Type: [InstanceNetworkInterfaceAttachment](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceNetworkInterfaceAttachment.html) object

Required: No

**connectionTrackingConfiguration**

A security group connection tracking configuration that enables you to set the timeout
for connection tracking on an Elastic network interface. For more information, see
[Connection tracking timeouts](../../../../services/ec2/latest/userguide/security-group-connection-tracking.md#connection-tracking-timeouts) in the
_Amazon EC2 User Guide_.

Type: [ConnectionTrackingSpecificationResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ConnectionTrackingSpecificationResponse.html) object

Required: No

**description**

The description.

Type: String

Required: No

**GroupSet.N**

The security groups.

Type: Array of [GroupIdentifier](api-groupidentifier.md) objects

Required: No

**interfaceType**

The type of network interface.

Valid values: `interface` \| `efa` \| `efa-only` \| `evs` \|
`trunk`

Type: String

Required: No

**Ipv4PrefixSet.N**

The IPv4 delegated prefixes that are assigned to the network interface.

Type: Array of [InstanceIpv4Prefix](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceIpv4Prefix.html) objects

Required: No

**Ipv6AddressesSet.N**

The IPv6 addresses associated with the network interface.

Type: Array of [InstanceIpv6Address](api-instanceipv6address.md) objects

Required: No

**Ipv6PrefixSet.N**

The IPv6 delegated prefixes that are assigned to the network interface.

Type: Array of [InstanceIpv6Prefix](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceIpv6Prefix.html) objects

Required: No

**macAddress**

The MAC address.

Type: String

Required: No

**networkInterfaceId**

The ID of the network interface.

Type: String

Required: No

**operator**

The service provider that manages the network interface.

Type: [OperatorResponse](api-operatorresponse.md) object

Required: No

**ownerId**

The ID of the AWS account that created the network interface.

Type: String

Required: No

**privateDnsName**

The private DNS name.

Type: String

Required: No

**privateIpAddress**

The IPv4 address of the network interface within the subnet.

Type: String

Required: No

**PrivateIpAddressesSet.N**

The private IPv4 addresses associated with the network interface.

Type: Array of [InstancePrivateIpAddress](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstancePrivateIpAddress.html) objects

Required: No

**sourceDestCheck**

Indicates whether source/destination checking is enabled.

Type: Boolean

Required: No

**status**

The status of the network interface.

Type: String

Valid Values: `available | associated | attaching | in-use | detaching`

Required: No

**subnetId**

The ID of the subnet.

Type: String

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/InstanceNetworkInterface)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/InstanceNetworkInterface)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/InstanceNetworkInterface)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceMonitoring

InstanceNetworkInterfaceAssociation
