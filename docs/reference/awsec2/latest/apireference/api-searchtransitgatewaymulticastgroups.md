# SearchTransitGatewayMulticastGroups

Searches one or more transit gateway multicast groups and returns the group membership information.

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

- `group-ip-address` \- The IP address of the transit gateway multicast group.

- `is-group-member` \- The resource is a group member. Valid values are `true` \| `false`.

- `is-group-source` \- The resource is a group source. Valid values are `true` \| `false`.

- `member-type` \- The member type. Valid values are `igmp` \| `static`.

- `resource-id` \- The ID of the resource.

- `resource-type` \- The type of resource. Valid values are `vpc` \| `vpn` \| `direct-connect-gateway` \| `tgw-peering`.

- `source-type` \- The source type. Valid values are `igmp` \| `static`.

- `subnet-id` \- The ID of the subnet.

- `transit-gateway-attachment-id` \- The id of the transit gateway attachment.

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

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**multicastGroups**

Information about the transit gateway multicast group.

Type: Array of [TransitGatewayMulticastGroup](api-transitgatewaymulticastgroup.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example returns the group membership information for the specified
multicast domain.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=SearchTransitGatewayMulticastGroups
&TransitGatewayMulticastDomainId=tgw-mcast-domain-0c4905cef79d6e597
&AUTHPARAMS
```

#### Sample Response

```

<SearchTransitGatewayMulticastGroupsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>19af1e8f-f80b-479d-9cf4-c7d19EXAMPLE</requestId>
    <multicastGroups>
        <item>
            <groupIpAddress>224.0.1.0</groupIpAddress>
            <groupMember>true</groupMember>
            <groupSource>false</groupSource>
            <memberType>static</memberType>
            <networkInterfaceId>eni-07f290fc3cEXAMPLE</networkInterfaceId>
            <subnetId>subnet-000de86e3bEXAMPLE</subnetId>
            <transitGatewayAttachmentId>tgw-attach-028c1dd0f8EXAMPLE</transitGatewayAttachmentId>
        </item>
        <item>
            <groupIpAddress>224.0.1.0</groupIpAddress>
            <groupMember>false</groupMember>
            <groupSource>true</groupSource>
            <networkInterfaceId>eni-0e246d32695012e81</networkInterfaceId>
            <sourceType>static</sourceType>
            <subnetId>subnet-000de86e3bEXAMPLE</subnetId>
            <transitGatewayAttachmentId>tgw-attach-028c1dd0f8EXAMPLE</transitGatewayAttachmentId>
        </item>
    </multicastGroups>
</SearchTransitGatewayMulticastGroupsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/SearchTransitGatewayMulticastGroups)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/SearchTransitGatewayMulticastGroups)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/SearchTransitGatewayMulticastGroups)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/searchtransitgatewaymulticastgroups.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SearchLocalGatewayRoutes

SearchTransitGatewayRoutes
