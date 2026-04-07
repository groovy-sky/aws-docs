# ModifyVerifiedAccessEndpoint

Modifies the configuration of the specified AWS Verified Access endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CidrOptions**

The CIDR options.

Type: [ModifyVerifiedAccessEndpointCidrOptions](api-modifyverifiedaccessendpointcidroptions.md) object

Required: No

**ClientToken**

A unique, case-sensitive token that you provide to ensure idempotency of your
modification request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

A description for the Verified Access endpoint.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**LoadBalancerOptions**

The load balancer details if creating the Verified Access endpoint as
`load-balancer` type.

Type: [ModifyVerifiedAccessEndpointLoadBalancerOptions](api-modifyverifiedaccessendpointloadbalanceroptions.md) object

Required: No

**NetworkInterfaceOptions**

The network interface options.

Type: [ModifyVerifiedAccessEndpointEniOptions](api-modifyverifiedaccessendpointenioptions.md) object

Required: No

**RdsOptions**

The RDS options.

Type: [ModifyVerifiedAccessEndpointRdsOptions](api-modifyverifiedaccessendpointrdsoptions.md) object

Required: No

**VerifiedAccessEndpointId**

The ID of the Verified Access endpoint.

Type: String

Required: Yes

**VerifiedAccessGroupId**

The ID of the Verified Access group.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**verifiedAccessEndpoint**

Details about the Verified Access endpoint.

Type: [VerifiedAccessEndpoint](api-verifiedaccessendpoint.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVerifiedAccessEndpoint)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVerifiedAccessEndpoint)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVerifiedAccessEndpoint)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyverifiedaccessendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyTransitGatewayVpcAttachment

ModifyVerifiedAccessEndpointPolicy
