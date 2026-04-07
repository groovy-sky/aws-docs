# AcceptTransitGatewayVpcAttachment

Accepts a request to attach a VPC to a transit gateway.

The VPC attachment must be in the `pendingAcceptance` state.
Use [DescribeTransitGatewayVpcAttachments](api-describetransitgatewayvpcattachments.md) to view your pending VPC attachment requests.
Use [RejectTransitGatewayVpcAttachment](api-rejecttransitgatewayvpcattachment.md) to reject a VPC attachment request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TransitGatewayAttachmentId**

The ID of the attachment.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayVpcAttachment**

The VPC attachment.

Type: [TransitGatewayVpcAttachment](api-transitgatewayvpcattachment.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AcceptTransitGatewayVpcAttachment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AcceptTransitGatewayVpcAttachment)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AcceptTransitGatewayVpcAttachment)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/accepttransitgatewayvpcattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceptTransitGatewayPeeringAttachment

AcceptVpcEndpointConnections
