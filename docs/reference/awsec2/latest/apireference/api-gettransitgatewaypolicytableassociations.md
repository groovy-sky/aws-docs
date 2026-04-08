# GetTransitGatewayPolicyTableAssociations

Gets a list of the transit gateway policy table associations.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters associated with the transit gateway policy table.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**TransitGatewayPolicyTableId**

The ID of the transit gateway policy table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**associations**

Returns details about the transit gateway policy table association.

Type: Array of [TransitGatewayPolicyTableAssociation](api-transitgatewaypolicytableassociation.md) objects

**nextToken**

The token for the next page of results.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/gettransitgatewaypolicytableassociations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetTransitGatewayMulticastDomainAssociations

GetTransitGatewayPolicyTableEntries
