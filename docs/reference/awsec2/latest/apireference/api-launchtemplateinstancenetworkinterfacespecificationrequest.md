# LaunchTemplateInstanceNetworkInterfaceSpecificationRequest

The parameters for a network interface.

## Contents

**AssociateCarrierIpAddress**

Associates a Carrier IP address with eth0 for a new network interface.

Use this option when you launch an instance in a Wavelength Zone and want to associate
a Carrier IP address with the network interface. For more information about Carrier IP
addresses, see [Carrier IP addresses](../../../../services/wavelength/latest/developerguide/how-wavelengths-work.md#provider-owned-ip) in the _AWS Wavelength Developer_
_Guide_.

Type: Boolean

Required: No

**AssociatePublicIpAddress**

Associates a public IPv4 address with eth0 for a new network interface.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab on the [Amazon VPC pricing page](http://aws.amazon.com/vpc/pricing).

Type: Boolean

Required: No

**ConnectionTrackingSpecification**

A security group connection tracking specification that enables you to set the timeout
for connection tracking on an Elastic network interface. For more information, see
[Idle connection tracking timeout](../../../../services/ec2/latest/userguide/security-group-connection-tracking.md#connection-tracking-timeouts) in the
_Amazon EC2 User Guide_.

Type: [ConnectionTrackingSpecificationRequest](api-connectiontrackingspecificationrequest.md) object

Required: No

**DeleteOnTermination**

Indicates whether the network interface is deleted when the instance is
terminated.

Type: Boolean

Required: No

**Description**

A description for the network interface.

Type: String

Required: No

**DeviceIndex**

The device index for the network interface attachment. The primary network interface
has a device index of 0. Each network interface is of type `interface`, you
must specify a device index. If you create a launch template that includes secondary
network interfaces but not a primary network interface, then you must add a primary
network interface as a launch parameter when you launch an instance from the
template.

Type: Integer

Required: No

**EnaQueueCount**

The number of ENA queues to be created with the instance.

Type: Integer

Required: No

**EnaSrdSpecification**

Configure ENA Express settings for your launch template.

Type: [EnaSrdSpecificationRequest](api-enasrdspecificationrequest.md) object

Required: No

**InterfaceType**

The type of network interface. To create an Elastic Fabric Adapter (EFA), specify
`efa` or `efa`. For more information, see [Elastic Fabric Adapter for AI/ML\
and HPC workloads on Amazon EC2](../../../../services/ec2/latest/userguide/efa.md) in the
_Amazon EC2 User Guide_.

If you are not creating an EFA, specify `interface` or omit this
parameter.

If you specify `efa-only`, do not assign any IP addresses to the network
interface. EFA-only network interfaces do not support IP addresses.

Valid values: `interface` \| `efa` \| `efa-only`

Type: String

Required: No

**Ipv4Prefix.N**

One or more IPv4 prefixes to be assigned to the network interface. You cannot use this
option if you use the `Ipv4PrefixCount` option.

Type: Array of [Ipv4PrefixSpecificationRequest](api-ipv4prefixspecificationrequest.md) objects

Required: No

**Ipv4PrefixCount**

The number of IPv4 prefixes to be automatically assigned to the network interface. You
cannot use this option if you use the `Ipv4Prefix` option.

Type: Integer

Required: No

**Ipv6AddressCount**

The number of IPv6 addresses to assign to a network interface. Amazon EC2
automatically selects the IPv6 addresses from the subnet range. You can't use this
option if specifying specific IPv6 addresses.

Type: Integer

Required: No

**Ipv6Addresses.N**

One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You
can't use this option if you're specifying a number of IPv6 addresses.

Type: Array of [InstanceIpv6AddressRequest](api-instanceipv6addressrequest.md) objects

Required: No

**Ipv6Prefix.N**

One or more IPv6 prefixes to be assigned to the network interface. You cannot use this
option if you use the `Ipv6PrefixCount` option.

Type: Array of [Ipv6PrefixSpecificationRequest](api-ipv6prefixspecificationrequest.md) objects

Required: No

**Ipv6PrefixCount**

The number of IPv6 prefixes to be automatically assigned to the network interface. You
cannot use this option if you use the `Ipv6Prefix` option.

Type: Integer

Required: No

**NetworkCardIndex**

The index of the network card. Some instance types support multiple network cards. The
primary network interface must be assigned to network card index 0. The default is
network card index 0.

Type: Integer

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: No

**PrimaryIpv6**

The primary IPv6 address of the network interface. When you enable an IPv6 GUA address
to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address until the
instance is terminated or the network interface is detached. For more information about
primary IPv6 addresses, see [RunInstances](api-runinstances.md).

Type: Boolean

Required: No

**PrivateIpAddress**

The primary private IPv4 address of the network interface.

Type: String

Required: No

**PrivateIpAddresses.N**

One or more private IPv4 addresses.

Type: Array of [PrivateIpAddressSpecification](api-privateipaddressspecification.md) objects

Required: No

**SecondaryPrivateIpAddressCount**

The number of secondary private IPv4 addresses to assign to a network
interface.

Type: Integer

Required: No

**SecurityGroupId.N**

The IDs of one or more security groups.

Type: Array of strings

Required: No

**SubnetId**

The ID of the subnet for the network interface.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateinstancenetworkinterfacespecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateinstancenetworkinterfacespecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateinstancenetworkinterfacespecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateInstanceNetworkInterfaceSpecification

LaunchTemplateInstanceSecondaryInterfaceSpecification
