# CreateTransitGatewayConnect

Creates a Connect attachment from a specified transit gateway attachment. A Connect attachment is a GRE-based tunnel attachment that you can use to establish a connection between a transit gateway and an appliance.

A Connect attachment uses an existing VPC or AWS Direct Connect attachment as the underlying transport mechanism.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Options**

The Connect attachment options.

Type: [CreateTransitGatewayConnectRequestOptions](api-createtransitgatewayconnectrequestoptions.md) object

Required: Yes

**TagSpecification.N**

The tags to apply to the Connect attachment.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TransportTransitGatewayAttachmentId**

The ID of the transit gateway attachment. You can specify a VPC attachment or AWS Direct Connect attachment.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayConnect**

Information about the Connect attachment.

Type: [TransitGatewayConnect](api-transitgatewayconnect.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createtransitgatewayconnect.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgatewayconnect.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGateway

CreateTransitGatewayConnectPeer
