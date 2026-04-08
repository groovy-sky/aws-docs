# DisassociateTrunkInterface

Removes an association between a branch network interface with a trunk network interface.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId**

The ID of the association

Type: String

Required: Yes

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. For more information, see [Ensuring\
idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. For more information, see [Ensuring\
idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disassociatetrunkinterface.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociatetrunkinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateTransitGatewayRouteTable

DisassociateVpcCidrBlock
