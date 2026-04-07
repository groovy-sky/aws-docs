# TrafficMirrorFilterRule

Describes the Traffic Mirror rule.

## Contents

**description**

The description of the Traffic Mirror rule.

Type: String

Required: No

**destinationCidrBlock**

The destination CIDR block assigned to the Traffic Mirror rule.

Type: String

Required: No

**destinationPortRange**

The destination port range assigned to the Traffic Mirror rule.

Type: [TrafficMirrorPortRange](api-trafficmirrorportrange.md) object

Required: No

**protocol**

The protocol assigned to the Traffic Mirror rule.

Type: Integer

Required: No

**ruleAction**

The action assigned to the Traffic Mirror rule.

Type: String

Valid Values: `accept | reject`

Required: No

**ruleNumber**

The rule number of the Traffic Mirror rule.

Type: Integer

Required: No

**sourceCidrBlock**

The source CIDR block assigned to the Traffic Mirror rule.

Type: String

Required: No

**sourcePortRange**

The source port range assigned to the Traffic Mirror rule.

Type: [TrafficMirrorPortRange](api-trafficmirrorportrange.md) object

Required: No

**TagSet.N**

Tags on Traffic Mirroring filter rules.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trafficDirection**

The traffic direction assigned to the Traffic Mirror rule.

Type: String

Valid Values: `ingress | egress`

Required: No

**trafficMirrorFilterId**

The ID of the Traffic Mirror filter that the rule is associated with.

Type: String

Required: No

**trafficMirrorFilterRuleId**

The ID of the Traffic Mirror rule.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/trafficmirrorfilterrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/trafficmirrorfilterrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/trafficmirrorfilterrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrafficMirrorFilter

TrafficMirrorPortRange
