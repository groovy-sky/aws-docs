# DeleteCarrierGateway

Deletes a carrier gateway.

###### Important

If you do not delete the route that contains the carrier gateway as the
Target, the route is a blackhole route. For information about how to delete a route, see
[DeleteRoute](api-deleteroute.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CarrierGatewayId**

The ID of the carrier gateway.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**carrierGateway**

Information about the carrier gateway.

Type: [CarrierGateway](api-carriergateway.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletecarriergateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletecarriergateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteCapacityManagerDataExport

DeleteClientVpnEndpoint
