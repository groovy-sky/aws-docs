# ScheduledInstancesNetworkInterface

Describes a network interface for a Scheduled Instance.

## Contents

**AssociatePublicIpAddress**

Indicates whether to assign a public IPv4 address to instances launched in a VPC. The
public IPv4 address can only be assigned to a network interface for eth0, and can only be
assigned to a new network interface, not an existing one. You cannot specify more than one
network interface in the request. If launching into a default subnet, the default value is
`true`.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab on the [Amazon VPC pricing page](http://aws.amazon.com/vpc/pricing).

Type: Boolean

Required: No

**DeleteOnTermination**

Indicates whether to delete the interface when the instance is terminated.

Type: Boolean

Required: No

**Description**

The description.

Type: String

Required: No

**DeviceIndex**

The index of the device for the network interface attachment.

Type: Integer

Required: No

**Group.N**

The IDs of the security groups.

Type: Array of strings

Required: No

**Ipv6Address.N**

The specific IPv6 addresses from the subnet range.

Type: Array of [ScheduledInstancesIpv6Address](api-scheduledinstancesipv6address.md) objects

Required: No

**Ipv6AddressCount**

The number of IPv6 addresses to assign to the network interface. The IPv6 addresses are automatically selected from the subnet range.

Type: Integer

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: No

**PrivateIpAddress**

The IPv4 address of the network interface within the subnet.

Type: String

Required: No

**PrivateIpAddressConfig.N**

The private IPv4 addresses.

Type: Array of [ScheduledInstancesPrivateIpAddressConfig](api-scheduledinstancesprivateipaddressconfig.md) objects

Required: No

**SecondaryPrivateIpAddressCount**

The number of secondary private IPv4 addresses.

Type: Integer

Required: No

**SubnetId**

The ID of the subnet.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/scheduledinstancesnetworkinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/scheduledinstancesnetworkinterface.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/scheduledinstancesnetworkinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScheduledInstancesMonitoring

ScheduledInstancesPlacement
