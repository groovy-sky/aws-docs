# IpamAddressHistoryRecord

The historical record of a CIDR within an IPAM scope. For more information, see [View the history of IP addresses](../../../../services/vpc/latest/ipam/view-history-cidr-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Contents

**resourceCidr**

The CIDR of the resource.

Type: String

Required: No

**resourceComplianceStatus**

The compliance status of a resource. For more information on compliance statuses, see [Monitor CIDR usage by resource](../../../../services/vpc/latest/ipam/monitor-cidr-compliance-ipam.md) in the _Amazon VPC IPAM User Guide_.

Type: String

Valid Values: `compliant | noncompliant | unmanaged | ignored`

Required: No

**resourceId**

The ID of the resource.

Type: String

Required: No

**resourceName**

The name of the resource.

Type: String

Required: No

**resourceOverlapStatus**

The overlap status of an IPAM resource. The overlap status tells you if the CIDR for a resource overlaps with another CIDR in the scope. For more information on overlap statuses, see [Monitor CIDR usage by resource](../../../../services/vpc/latest/ipam/monitor-cidr-compliance-ipam.md) in the _Amazon VPC IPAM User Guide_.

Type: String

Valid Values: `overlapping | nonoverlapping | ignored`

Required: No

**resourceOwnerId**

The ID of the resource owner.

Type: String

Required: No

**resourceRegion**

The AWS Region of the resource.

Type: String

Required: No

**resourceType**

The type of the resource.

Type: String

Valid Values: `eip | vpc | subnet | network-interface | instance`

Required: No

**sampledEndTime**

Sampled end time of the resource-to-CIDR association within the IPAM scope. Changes are picked up in periodic snapshots, so the end time may have occurred before this specific time.

Type: Timestamp

Required: No

**sampledStartTime**

Sampled start time of the resource-to-CIDR association within the IPAM scope. Changes are picked up in periodic snapshots, so the start time may have occurred before this specific time.

Type: Timestamp

Required: No

**vpcId**

The VPC ID of the resource.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipamaddresshistoryrecord.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipamaddresshistoryrecord.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipamaddresshistoryrecord.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Ipam

IpamCidrAuthorizationContext
