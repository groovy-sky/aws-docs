# DescribeTransitGatewayAttachments

Describes one or more attachments between resources and transit gateways. By default, all attachments are described.
Alternatively, you can filter the results by attachment ID, attachment state, resource ID, or resource owner.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. The possible values are:

- `association.state` \- The state of the association ( `associating` \| `associated` \|
`disassociating`).

- `association.transit-gateway-route-table-id` \- The ID of the route table for the transit gateway.

- `resource-id` \- The ID of the resource.

- `resource-owner-id` \- The ID of the AWS account that owns the resource.

- `resource-type` \- The resource type. Valid values are `vpc`
\| `vpn` \| `direct-connect-gateway` \| `peering`
\| `connect`.

- `state` \- The state of the attachment. Valid values are `available` \| `deleted` \| `deleting` \| `failed` \| `failing` \| `initiatingRequest` \| `modifying` \| `pendingAcceptance` \| `pending` \| `rollingBack` \| `rejected` \| `rejecting`.

- `transit-gateway-attachment-id` \- The ID of the attachment.

- `transit-gateway-id` \- The ID of the transit gateway.

- `transit-gateway-owner-id` \- The ID of the AWS account that owns the transit gateway.

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

**TransitGatewayAttachmentIds.N**

The IDs of the attachments.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**transitGatewayAttachments**

Information about the attachments.

Type: Array of [TransitGatewayAttachment](api-transitgatewayattachment.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describetransitgatewayattachments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describetransitgatewayattachments.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeTrafficMirrorTargets

DescribeTransitGatewayConnectPeers
