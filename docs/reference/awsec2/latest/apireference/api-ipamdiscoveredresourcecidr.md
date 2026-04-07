# IpamDiscoveredResourceCidr

An IPAM discovered resource CIDR. A discovered resource is a resource CIDR monitored under a resource discovery. The following resources can be discovered: VPCs, Public IPv4 pools, VPC subnets, and Elastic IP addresses. The discovered resource CIDR is the IP address range in CIDR notation that is associated with the resource.

## Contents

**availabilityZoneId**

The Availability Zone ID.

Type: String

Required: No

**ipamResourceDiscoveryId**

The resource discovery ID.

Type: String

Required: No

**ipSource**

The source that allocated the IP address space. `byoip` or `amazon` indicates public IP address space allocated by Amazon or space that you have allocated with Bring your own IP (BYOIP). `none` indicates private space.

Type: String

Valid Values: `amazon | byoip | none`

Required: No

**ipUsage**

The percentage of IP address space in use. To convert the decimal to a percentage, multiply the decimal by 100. Note the following:

- For resources that are VPCs, this is the percentage of IP address space in the VPC that's taken up by subnet CIDRs.

- For resources that are subnets, if the subnet has an IPv4 CIDR provisioned to it, this is the percentage of IPv4 address space in the subnet that's in use. If the subnet has an IPv6 CIDR provisioned to it, the percentage of IPv6 address space in use is not represented. The percentage of IPv6 address space in use cannot currently be calculated.

- For resources that are public IPv4 pools, this is the percentage of IP address space in the pool that's been allocated to Elastic IP addresses (EIPs).

Type: Double

Required: No

**networkInterfaceAttachmentStatus**

For elastic network interfaces, this is the status of whether or not the elastic network interface is attached.

Type: String

Valid Values: `available | in-use`

Required: No

**resourceCidr**

The resource CIDR.

Type: String

Required: No

**resourceId**

The resource ID.

Type: String

Required: No

**resourceOwnerId**

The resource owner ID.

Type: String

Required: No

**resourceRegion**

The resource Region.

Type: String

Required: No

**ResourceTagSet.N**

The resource tags.

Type: Array of [IpamResourceTag](api-ipamresourcetag.md) objects

Required: No

**resourceType**

The resource type.

Type: String

Valid Values: `vpc | subnet | eip | public-ipv4-pool | ipv6-pool | eni | anycast-ip-list`

Required: No

**sampleTime**

The last successful resource discovery time.

Type: Timestamp

Required: No

**subnetId**

The subnet ID.

Type: String

Required: No

**vpcId**

The VPC ID.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipamdiscoveredresourcecidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipamdiscoveredresourcecidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipamdiscoveredresourcecidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamDiscoveredPublicAddress

IpamDiscoveryFailureReason
