# CreateTrafficMirrorSession

Creates a Traffic Mirror session.

A Traffic Mirror session actively copies packets from a Traffic Mirror source to a Traffic Mirror target. Create a filter, and then assign it
to the session to define a subset of the traffic to mirror, for example all TCP
traffic.

The Traffic Mirror source and the Traffic Mirror target (monitoring appliances) can be in the same VPC, or in a different VPC connected via VPC peering or a transit gateway.

By default, no traffic is mirrored. Use [CreateTrafficMirrorFilter](api-createtrafficmirrorfilter.md) to
create filter rules that specify the traffic to mirror.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

The description of the Traffic Mirror session.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**NetworkInterfaceId**

The ID of the source network interface.

Type: String

Required: Yes

**PacketLength**

The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do
not specify this parameter when you want to mirror the entire packet. To mirror a subset of
the packet, set this to the length (in bytes) that you want to mirror. For example, if you
set this value to 100, then the first 100 bytes that meet the filter criteria are copied to
the target.

If you do not want to mirror the entire packet, use the `PacketLength` parameter to specify the number of bytes in each packet to mirror.

For sessions with Network Load Balancer (NLB) Traffic Mirror targets the default `PacketLength` will be set to 8500. Valid values are 1-8500. Setting a `PacketLength` greater than 8500 will result in an error response.

Type: Integer

Required: No

**SessionNumber**

The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.

Valid values are 1-32766.

Type: Integer

Required: Yes

**TagSpecification.N**

The tags to assign to a Traffic Mirror session.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TrafficMirrorFilterId**

The ID of the Traffic Mirror filter.

Type: String

Required: Yes

**TrafficMirrorTargetId**

The ID of the Traffic Mirror target.

Type: String

Required: Yes

**VirtualNetworkId**

The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN
protocol, see [RFC 7348](https://datatracker.ietf.org/doc/html/rfc7348). If you do
not specify a `VirtualNetworkId`, an account-wide unique ID is chosen at
random.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

**requestId**

The ID of the request.

Type: String

**trafficMirrorSession**

Information about the Traffic Mirror session.

Type: [TrafficMirrorSession](api-trafficmirrorsession.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateTrafficMirrorSession)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateTrafficMirrorSession)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtrafficmirrorsession.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateTrafficMirrorSession)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtrafficmirrorsession.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTrafficMirrorFilterRule

CreateTrafficMirrorTarget
