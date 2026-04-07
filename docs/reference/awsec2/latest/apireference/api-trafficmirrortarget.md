# TrafficMirrorTarget

Describes a Traffic Mirror target.

## Contents

**description**

Information about the Traffic Mirror target.

Type: String

Required: No

**gatewayLoadBalancerEndpointId**

The ID of the Gateway Load Balancer endpoint.

Type: String

Required: No

**networkInterfaceId**

The network interface ID that is attached to the target.

Type: String

Required: No

**networkLoadBalancerArn**

The Amazon Resource Name (ARN) of the Network Load Balancer.

Type: String

Required: No

**ownerId**

The ID of the account that owns the Traffic Mirror target.

Type: String

Required: No

**TagSet.N**

The tags assigned to the Traffic Mirror target.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trafficMirrorTargetId**

The ID of the Traffic Mirror target.

Type: String

Required: No

**type**

The type of Traffic Mirror target.

Type: String

Valid Values: `network-interface | network-load-balancer | gateway-load-balancer-endpoint`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/trafficmirrortarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/trafficmirrortarget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/trafficmirrortarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrafficMirrorSession

TransitGateway
