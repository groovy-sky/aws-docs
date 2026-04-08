# DescribeClientVpnTargetNetworks

Describes the target networks associated with the specified Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationIds.N**

The IDs of the target network associations.

Type: Array of strings

Required: No

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. Filter names and values are case-sensitive.

- `association-id` \- The ID of the association.

- `target-network-id` \- The ID of the subnet specified as the target network.

- `vpc-id` \- The ID of the VPC in which the target network is located.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining results can be seen by sending another request with the nextToken value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**clientVpnTargetNetworks**

Information about the associated target networks.

Type: Array of [TargetNetwork](api-targetnetwork.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the target networks associated with a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeClientVpnTargetNetworks
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DescribeClientVpnTargetNetworksResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>1f92d56a-4494-4cbe-ad85-d9387EXAMPLE</requestId>
    <clientVpnTargetNetworks>
        <item>
            <associationId>cvpn-assoc-0822b0983cEXAMPLE</associationId>
            <clientVpnEndpointId>cvpn-endpoint-00c5d11fc4EXAMPLE</clientVpnEndpointId>
            <targetNetworkId>subnet-057fa0918fEXAMPLE</targetNetworkId>
            <securityGroups>
                <item>sg-123456EXAMPLE</item>
            </securityGroups>
            <status>
                <code>associated</code>
            </status>
            <vpcId>vpc-3db97056EXAMPLE</vpcId>
        </item>
    </clientVpnTargetNetworks>
</DescribeClientVpnTargetNetworksRespons>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeclientvpntargetnetworks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeclientvpntargetnetworks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeClientVpnRoutes

DescribeCoipPools
