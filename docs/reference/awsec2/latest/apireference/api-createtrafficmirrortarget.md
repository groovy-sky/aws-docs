# CreateTrafficMirrorTarget

Creates a target for your Traffic Mirror session.

A Traffic Mirror target is the destination for mirrored traffic. The Traffic Mirror source and
the Traffic Mirror target (monitoring appliances) can be in the same VPC, or in
different VPCs connected via VPC peering or a transit gateway.

A Traffic Mirror target can be a network interface, a Network Load Balancer, or a Gateway Load Balancer endpoint.

To use the target in a Traffic Mirror session, use [CreateTrafficMirrorSession](api-createtrafficmirrorsession.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

The description of the Traffic Mirror target.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GatewayLoadBalancerEndpointId**

The ID of the Gateway Load Balancer endpoint.

Type: String

Required: No

**NetworkInterfaceId**

The network interface ID that is associated with the target.

Type: String

Required: No

**NetworkLoadBalancerArn**

The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.

Type: String

Required: No

**TagSpecification.N**

The tags to assign to the Traffic Mirror target.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

**requestId**

The ID of the request.

Type: String

**trafficMirrorTarget**

Information about the Traffic Mirror target.

Type: [TrafficMirrorTarget](api-trafficmirrortarget.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateTrafficMirrorTarget)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateTrafficMirrorTarget)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtrafficmirrortarget.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateTrafficMirrorTarget)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtrafficmirrortarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTrafficMirrorSession

CreateTransitGateway
