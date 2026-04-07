# AssociateTransitGatewayRouteTable

Associates the specified attachment with the specified transit gateway route table. You can
associate only one route table with an attachment.

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

**TransitGatewayRouteTableId**

The ID of the transit gateway route table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**association**

The ID of the association.

Type: [TransitGatewayAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAssociation.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AssociateTransitGatewayRouteTable)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssociateTransitGatewayPolicyTable

AssociateTrunkInterface
