# DescribeTransitGatewayPeeringAttachments

Describes your transit gateway peering attachments.

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

- `transit-gateway-attachment-id` \- The ID of the transit gateway attachment.

- `local-owner-id` \- The ID of your AWS account.

- `remote-owner-id` \- The ID of the AWS account in the remote Region that owns the transit gateway.

- `state` \- The state of the peering attachment. Valid values are `available` \| `deleted` \| `deleting` \| `failed` \| `failing` \| `initiatingRequest` \| `modifying` \| `pendingAcceptance` \| `pending` \| `rollingBack` \| `rejected` \| `rejecting`).

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources that have a tag with a specific key, regardless of the tag value.

- `transit-gateway-id` \- The ID of the transit gateway.

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

One or more IDs of the transit gateway peering attachments.

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

**transitGatewayPeeringAttachments**

The transit gateway peering attachments.

Type: Array of [TransitGatewayPeeringAttachment](api-transitgatewaypeeringattachment.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes your peering attachment.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTransitGatewayPeeringAttachments
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTransitGatewayPeeringAttachmentsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>f2ad2616-b1bc-42ab-8533-bd50example</requestId>
    <transitGatewayPeeringAttachments>
        <item>
            <accepterTgwInfo>
                <ownerId>111111111111</ownerId>
                <region>us-west-2</region>
                <transitGatewayId>tgw-123456789012abc12</transitGatewayId>
            </accepterTgwInfo>
            <creationTime>2019-11-11T11:25:31.000Z</creationTime>
            <requesterTgwInfo>
                <ownerId>123456789012</ownerId>
                <region>us-east-1</region>
                <transitGatewayId>tgw-abc123abc123abc12</transitGatewayId>
            </requesterTgwInfo>
            <state>pendingAcceptance</state>
            <tagSet/>
            <transitGatewayAttachmentId>tgw-attach-1122334455aabbcc1</transitGatewayAttachmentId>
        </item>
    </transitGatewayPeeringAttachments>
</DescribeTransitGatewayPeeringAttachmentsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describetransitgatewaypeeringattachments.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeTransitGatewayMulticastDomains

DescribeTransitGatewayPolicyTables
