# IpamPoolCidr

A CIDR provisioned to an IPAM pool.

## Contents

**cidr**

The CIDR provisioned to the IPAM pool. A CIDR is a representation of an IP address and its associated network mask (or netmask)
and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23`. An IPv6 CIDR example is `2001:DB8::/32`.

Type: String

Required: No

**failureReason**

Details related to why an IPAM pool CIDR failed to be provisioned.

Type: [IpamPoolCidrFailureReason](api-ipampoolcidrfailurereason.md) object

Required: No

**ipamPoolCidrId**

The IPAM pool CIDR ID.

Type: String

Required: No

**netmaskLength**

The netmask length of the CIDR you'd like to provision to a pool. Can be used for provisioning Amazon-provided IPv6 CIDRs to top-level pools and for provisioning CIDRs to pools with source pools. Cannot be used to provision BYOIP CIDRs to top-level pools. "NetmaskLength" or "Cidr" is required.

Type: Integer

Required: No

**state**

The state of the CIDR.

Type: String

Valid Values: `pending-provision | provisioned | failed-provision | pending-deprovision | deprovisioned | failed-deprovision | pending-import | failed-import`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipampoolcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipampoolcidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipampoolcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamPoolAllocation

IpamPoolCidrFailureReason
