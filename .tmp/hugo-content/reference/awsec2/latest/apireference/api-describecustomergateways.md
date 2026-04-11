# DescribeCustomerGateways

Describes one or more of your VPN customer gateways.

For more information, see [AWS Site-to-Site VPN](../../../../services/vpn/latest/s2svpn/vpc-vpn.md) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CustomerGatewayId.N**

One or more customer gateway IDs.

Default: Describes all your customer gateways.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `bgp-asn` \- The customer gateway's Border Gateway Protocol (BGP)
Autonomous System Number (ASN).

- `customer-gateway-id` \- The ID of the customer gateway.

- `ip-address` \- The IP address of the customer gateway
device's external interface.

- `state` \- The state of the customer gateway ( `pending` \|
`available` \| `deleting` \|
`deleted`).

- `type` \- The type of customer gateway. Currently, the only
supported type is `ipsec.1`.

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**customerGatewaySet**

Information about one or more customer gateways.

Type: Array of [CustomerGateway](api-customergateway.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example request describes the specified customer gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeCustomerGateways
&CustomerGatewayId.1=cgw-b4dc3961
&AUTHPARAMS
```

#### Sample Response

```

<DescribeCustomerGatewaysResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
  <customerGatewaySet>
    <item>
       <customerGatewayId>cgw-b4dc3961</customerGatewayId>
       <state>available</state>
       <type>ipsec.1</type>
       <ipAddress>12.1.2.3</ipAddress>
       <bgpAsn>65534</bgpasn>
       <tagSet/>
    </item>
  </customerGatewaySet>
</DescribeCustomerGatewaysResponse>
```

### Example 2

This example request uses filters to describe any customer gateway you own
whose IP address is `12.1.2.3`, and whose state is either
`pending` or `available`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeCustomerGateways
&Filter.1.Name=ip-address
&Filter.1.Value.1=12.1.2.3
&Filter.2.Name=state
&Filter.2.Value.1=pending
&Filter.2.Value.2=available
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describecustomergateways.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecustomergateways.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeConversionTasks

DescribeDeclarativePoliciesReports
