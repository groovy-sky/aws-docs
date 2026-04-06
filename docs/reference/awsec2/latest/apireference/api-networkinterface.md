# NetworkInterface

Describes a network interface.

## Contents

**AssociatedSubnetSet.N**

The subnets associated with this network interface.

Type: Array of strings

Required: No

**association**

The association information for an Elastic IP address (IPv4) associated with the
network interface.

Type: [NetworkInterfaceAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterfaceAssociation.html) object

Required: No

**attachment**

The network interface attachment.

Type: [NetworkInterfaceAttachment](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterfaceAttachment.html) object

Required: No

**availabilityZone**

The Availability Zone.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone.

Type: String

Required: No

**connectionTrackingConfiguration**

A security group connection tracking configuration that enables you to set the timeout
for connection tracking on an Elastic network interface. For more information, see
[Connection tracking timeouts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-connection-tracking.html#connection-tracking-timeouts) in the
_Amazon EC2 User Guide_.

Type: [ConnectionTrackingConfiguration](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ConnectionTrackingConfiguration.html) object

Required: No

**denyAllIgwTraffic**

Indicates whether a network interface with an IPv6 address is unreachable from the
public internet. If the value is `true`, inbound traffic from the internet is
dropped and you cannot assign an elastic IP address to the network interface. The
network interface is reachable from peered VPCs and resources connected through a
transit gateway, including on-premises networks.

Type: Boolean

Required: No

**description**

A description.

Type: String

Required: No

**GroupSet.N**

Any security groups for the network interface.

Type: Array of [GroupIdentifier](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GroupIdentifier.html) objects

Required: No

**interfaceType**

The type of network interface.

Type: String

Valid Values: `api_gateway_managed | aws_codestar_connections_managed | branch | ec2_instance_connect_endpoint | efa | efa-only | efs | evs | gateway_load_balancer | gateway_load_balancer_endpoint | global_accelerator_managed | interface | iot_rules_managed | lambda | load_balancer | nat_gateway | network_load_balancer | quicksight | transit_gateway | trunk | vpc_endpoint`

Required: No

**Ipv4PrefixSet.N**

The IPv4 prefixes that are assigned to the network interface.

Type: Array of [Ipv4PrefixSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Ipv4PrefixSpecification.html) objects

Required: No

**ipv6Address**

The IPv6 globally unique address associated with the network interface.

Type: String

Required: No

**Ipv6AddressesSet.N**

The IPv6 addresses associated with the network interface.

Type: Array of [NetworkInterfaceIpv6Address](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterfaceIpv6Address.html) objects

Required: No

**ipv6Native**

Indicates whether this is an IPv6 only network interface.

Type: Boolean

Required: No

**Ipv6PrefixSet.N**

The IPv6 prefixes that are assigned to the network interface.

Type: Array of [Ipv6PrefixSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Ipv6PrefixSpecification.html) objects

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

Type: [OperatorResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_OperatorResponse.html) object

Required: No

**outpostArn**

The Amazon Resource Name (ARN) of the Outpost.

Type: String

Required: No

**ownerId**

The AWS account ID of the owner of the network interface.

Type: String

Required: No

**privateDnsName**

The private hostname. For more information, see [EC2 instance hostnames, DNS names, and domains](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**privateIpAddress**

The IPv4 address of the network interface within the subnet.

Type: String

Required: No

**PrivateIpAddressesSet.N**

The private IPv4 addresses associated with the network interface.

Type: Array of [NetworkInterfacePrivateIpAddress](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterfacePrivateIpAddress.html) objects

Required: No

**publicDnsName**

A public hostname. For more information, see [EC2 instance hostnames, DNS names, and domains](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**publicIpDnsNameOptions**

Public hostname type options. For more information, see [EC2 instance hostnames, DNS names, and domains](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the _Amazon EC2 User Guide_.

Type: [PublicIpDnsNameOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PublicIpDnsNameOptions.html) object

Required: No

**requesterId**

The alias or AWS account ID of the principal or service that created
the network interface.

Type: String

Required: No

**requesterManaged**

Indicates whether the network interface is being managed by AWS.

Type: Boolean

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

**TagSet.N**

Any tags assigned to the network interface.

Type: Array of [Tag](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Tag.html) objects

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/NetworkInterface)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/NetworkInterface)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/NetworkInterface)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkInsightsPath

NetworkInterfaceAssociation
