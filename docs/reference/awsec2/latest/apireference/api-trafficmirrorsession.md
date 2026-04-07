# TrafficMirrorSession

Describes a Traffic Mirror session.

## Contents

**description**

The description of the Traffic Mirror session.

Type: String

Required: No

**networkInterfaceId**

The ID of the Traffic Mirror session's network interface.

Type: String

Required: No

**ownerId**

The ID of the account that owns the Traffic Mirror session.

Type: String

Required: No

**packetLength**

The number of bytes in each packet to mirror. These are the bytes after the VXLAN header. To mirror a subset, set this to the length (in bytes) to mirror. For example, if you set this value to 100, then the first 100 bytes that meet the filter criteria are copied to the target. Do not specify this parameter when you want to mirror the entire packet

Type: Integer

Required: No

**sessionNumber**

The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.

Valid values are 1-32766.

Type: Integer

Required: No

**TagSet.N**

The tags assigned to the Traffic Mirror session.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trafficMirrorFilterId**

The ID of the Traffic Mirror filter.

Type: String

Required: No

**trafficMirrorSessionId**

The ID for the Traffic Mirror session.

Type: String

Required: No

**trafficMirrorTargetId**

The ID of the Traffic Mirror target.

Type: String

Required: No

**virtualNetworkId**

The virtual network ID associated with the Traffic Mirror session.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/trafficmirrorsession.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/trafficmirrorsession.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/trafficmirrorsession.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrafficMirrorPortRangeRequest

TrafficMirrorTarget
