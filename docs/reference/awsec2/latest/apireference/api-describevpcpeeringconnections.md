# DescribeVpcPeeringConnections

Describes your VPC peering connections. The default is to describe all your VPC peering connections.
Alternatively, you can specify specific VPC peering connection IDs or filter the results to
include only the VPC peering connections that match specific criteria.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `accepter-vpc-info.cidr-block` \- The IPv4 CIDR block of the accepter
VPC.

- `accepter-vpc-info.owner-id` \- The ID of the AWS account that owns the
accepter VPC.

- `accepter-vpc-info.vpc-id` \- The ID of the accepter VPC.

- `expiration-time` \- The expiration date and time for the VPC peering
connection.

- `requester-vpc-info.cidr-block` \- The IPv4 CIDR block of the
requester's VPC.

- `requester-vpc-info.owner-id` \- The ID of the AWS account that owns the
requester VPC.

- `requester-vpc-info.vpc-id` \- The ID of the requester VPC.

- `status-code` \- The status of the VPC peering connection
( `pending-acceptance` \| `failed` \|
`expired` \| `provisioning` \| `active` \|
`deleting` \| `deleted` \|
`rejected`).

- `status-message` \- A message that provides more information about the status
of the VPC peering connection, if applicable.

- `tag` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `vpc-peering-connection-id` \- The ID of the VPC peering connection.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**VpcPeeringConnectionId.N**

The IDs of the VPC peering connections.

Default: Describes all your VPC peering connections.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**vpcPeeringConnectionSet**

Information about the VPC peering connections.

Type: Array of [VpcPeeringConnection](api-vpcpeeringconnection.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes all of your VPC peering connections.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcPeeringConnections
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpcPeeringConnectionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
 <vpcPeeringConnectionSet>
 <item>
  <vpcPeeringConnectionId>pcx-111aaa22</vpcPeeringConnectionId>
    <requesterVpcInfo>
     <ownerId>777788889999</ownerId>
     <vpcId>vpc-1a2b3c4d</vpcId>
     <cidrBlock>172.31.0.0/16</cidrBlock>
    </requesterVpcInfo>
    <accepterVpcInfo>
     <ownerId>123456789012</ownerId>
     <vpcId>vpc-aa22cc33</vpcId>
     <cidrBlock>10.0.0.0/16</cidrBlock>
     <peeringOptions>
        <allowEgressFromLocalClassicLinkToRemoteVpc>false</allowEgressFromLocalClassicLinkToRemoteVpc>
        <allowEgressFromLocalVpcToRemoteClassicLink>true</allowEgressFromLocalVpcToRemoteClassicLink>
        <allowDnsResolutionFromRemoteVpc>false</allowDnsResolutionFromRemoteVpc>
     </peeringOptions>
    </accepterVpcInfo>
     <status>
      <code>active</code>
      <message>Active</message>
     </status>
     <tagSet/>
     </item>
   </vpcPeeringConnectionSet>
</DescribeVpcPeeringConnectionsResponse>
```

### Example 2

This example describes all of your VPC peering connections that are in the
`pending-acceptance` state.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcPeeringConnections
&Filter.1.Name=status-code
&Filter.1.Value=pending-acceptance
&AUTHPARAMS
```

### Example 3

This example describes all of your VPC peering connections that have the tag
`Name=Finance` or `Name=Accounts`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcPeeringConnections
&Filter.1.Name=tag:Name
&Filter.1.Value.1=Finance
&Filter.1.Value.2=Accounts
&AUTHPARAMS
```

### Example 4

This example describes all of the VPC peering connections for your specified
VPC, `vpc-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcPeeringConnections
&Filter.1.Name=requester-vpc-info.vpc-id
&Filter.1.Value=vpc-1a2b3c4d
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeVpcPeeringConnections)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeVpcPeeringConnections)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpcpeeringconnections.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeVpcPeeringConnections)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpcpeeringconnections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpcEndpointServices

DescribeVpcs
