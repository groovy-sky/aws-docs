# DescribeClientVpnRoutes

Describes the routes for the specified Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

- `destination-cidr` \- The CIDR of the route destination.

- `origin` \- How the route was associated with the Client VPN endpoint ( `associate` \| `add-route`).

- `target-subnet` \- The ID of the subnet through which traffic is routed.

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

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**routes**

Information about the Client VPN endpoint routes.

Type: Array of [ClientVpnRoute](api-clientvpnroute.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the routes for a specific Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeClientVpnRoutes
&ClientVpnEndpointId.1=cvpn-endpoint-EXAMPLEc8db8d3536
&AUTHPARAMS
```

#### Sample Response

```

<DescribeClientVpnRoutesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>08fb643f-4d8f-443b-b853-EXAMPLE9cc8c</requestId>
    <routes>
        <item>
            <destinationCidr>10.0.0.0/16</destinationCidr>
            <clientVpnEndpointId>cvpn-endpoint-EXAMPLEc8db8d3536</clientVpnEndpointId>
            <origin>associate</origin>
            <status>
                <code>active</code>
            </status>
            <targetSubnet>subnet-EXAMPLE18f440ab91</targetSubnet>
            <type>Nat</type>
        </item>
        <item>
            <destinationCidr>10.0.1.128/28</destinationCidr>
            <clientVpnEndpointId>cvpn-endpoint-EXAMPLEc8db8d3536</clientVpnEndpointId>
            <origin>add-route</origin>
            <status>
                <code>active</code>
            </status>
            <targetSubnet>EXAMPLE18f440ab91</targetSubnet>
            <type>Nat</type>
        </item>
    </routes>
</DescribeClientVpnRoutesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeclientvpnroutes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeclientvpnroutes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeClientVpnEndpoints

DescribeClientVpnTargetNetworks
