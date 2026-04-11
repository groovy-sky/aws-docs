# TrafficMirrorFilter

Describes the Traffic Mirror filter.

## Contents

**description**

The description of the Traffic Mirror filter.

Type: String

Required: No

**EgressFilterRuleSet.N**

Information about the egress rules that are associated with the Traffic Mirror filter.

Type: Array of [TrafficMirrorFilterRule](api-trafficmirrorfilterrule.md) objects

Required: No

**IngressFilterRuleSet.N**

Information about the ingress rules that are associated with the Traffic Mirror filter.

Type: Array of [TrafficMirrorFilterRule](api-trafficmirrorfilterrule.md) objects

Required: No

**NetworkServiceSet.N**

The network service traffic that is associated with the Traffic Mirror filter.

Type: Array of strings

Valid Values: `amazon-dns`

Required: No

**TagSet.N**

The tags assigned to the Traffic Mirror filter.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trafficMirrorFilterId**

The ID of the Traffic Mirror filter.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/trafficmirrorfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/trafficmirrorfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/trafficmirrorfilter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TotalLocalStorageGBRequest

TrafficMirrorFilterRule
