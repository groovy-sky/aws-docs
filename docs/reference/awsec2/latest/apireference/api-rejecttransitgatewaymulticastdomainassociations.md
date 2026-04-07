# RejectTransitGatewayMulticastDomainAssociations

Rejects a request to associate cross-account subnets with a transit gateway multicast domain.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SubnetIds.N**

The IDs of the subnets to associate with the transit gateway multicast domain.

Type: Array of strings

Required: No

**TransitGatewayAttachmentId**

The ID of the transit gateway attachment.

Type: String

Required: No

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**associations**

Information about the multicast domain associations.

Type: [TransitGatewayMulticastDomainAssociations](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomainAssociations.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/RejectTransitGatewayMulticastDomainAssociations)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RejectCapacityReservationBillingOwnership

RejectTransitGatewayPeeringAttachment
