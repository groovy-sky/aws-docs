# ModifyTrafficMirrorFilterRule

Modifies the specified Traffic Mirror rule.

`DestinationCidrBlock` and `SourceCidrBlock` must both be an IPv4
range or an IPv6 range.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Description**

The description to assign to the Traffic Mirror rule.

Type: String

Required: No

**DestinationCidrBlock**

The destination CIDR block to assign to the Traffic Mirror rule.

Type: String

Required: No

**DestinationPortRange**

The destination ports that are associated with the Traffic Mirror rule.

Type: [TrafficMirrorPortRangeRequest](api-trafficmirrorportrangerequest.md) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Protocol**

The protocol, for example TCP, to assign to the Traffic Mirror rule.

Type: Integer

Required: No

**RemoveField.N**

The properties that you want to remove from the Traffic Mirror filter rule.

When you remove a property from a Traffic Mirror filter rule, the property is set to the default.

Type: Array of strings

Valid Values: `destination-port-range | source-port-range | protocol | description`

Required: No

**RuleAction**

The action to assign to the rule.

Type: String

Valid Values: `accept | reject`

Required: No

**RuleNumber**

The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given
direction. The rules are processed in ascending order by rule number.

Type: Integer

Required: No

**SourceCidrBlock**

The source CIDR block to assign to the Traffic Mirror rule.

Type: String

Required: No

**SourcePortRange**

The port range to assign to the Traffic Mirror rule.

Type: [TrafficMirrorPortRangeRequest](api-trafficmirrorportrangerequest.md) object

Required: No

**TrafficDirection**

The type of traffic to assign to the rule.

Type: String

Valid Values: `ingress | egress`

Required: No

**TrafficMirrorFilterRuleId**

The ID of the Traffic Mirror rule.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**trafficMirrorFilterRule**

###### Note

Tags are not returned for ModifyTrafficMirrorFilterRule.

A Traffic Mirror rule.

Type: [TrafficMirrorFilterRule](api-trafficmirrorfilterrule.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifytrafficmirrorfilterrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyTrafficMirrorFilterNetworkServices

ModifyTrafficMirrorSession
