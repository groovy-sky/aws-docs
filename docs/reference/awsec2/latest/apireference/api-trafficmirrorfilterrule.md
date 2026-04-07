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

Type: [TrafficMirrorPortRange](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TrafficMirrorPortRange.html) object

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

Type: [TrafficMirrorPortRange](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TrafficMirrorPortRange.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TrafficMirrorFilterRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TrafficMirrorFilterRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TrafficMirrorFilterRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrafficMirrorFilter

TrafficMirrorPortRange
