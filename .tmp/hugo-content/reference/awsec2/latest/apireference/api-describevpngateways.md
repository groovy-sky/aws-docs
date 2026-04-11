# DescribeVpnGateways

Describes one or more of your virtual private gateways.

For more information, see [AWS Site-to-Site VPN](../../../../services/vpn/latest/s2svpn/vpc-vpn.md) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `amazon-side-asn` \- The Autonomous System Number (ASN) for the
Amazon side of the gateway.

- `attachment.state` \- The current state of the attachment between
the gateway and the VPC ( `attaching` \| `attached` \|
`detaching` \| `detached`).

- `attachment.vpc-id` \- The ID of an attached VPC.

- `availability-zone` \- The Availability Zone for the virtual private
gateway (if applicable).

- `state` \- The state of the virtual private gateway
( `pending` \| `available` \| `deleting` \|
`deleted`).

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `type` \- The type of virtual private gateway. Currently the only
supported type is `ipsec.1`.

- `vpn-gateway-id` \- The ID of the virtual private gateway.

Type: Array of [Filter](api-filter.md) objects

Required: No

**VpnGatewayId.N**

One or more virtual private gateway IDs.

Default: Describes all your virtual private gateways.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpnGatewaySet**

Information about one or more virtual private gateways.

Type: Array of [VpnGateway](api-vpngateway.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the specified virtual private gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpnGateways
&VpnGatewayId.1=vgw-8db04f81
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpnGatewaysResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
  <vpnGatewaySet>
        <item>
            <vpnGatewayId>vgw-8db04f81</vpnGatewayId>
            <state>available</state>
            <type>ipsec.1</type>
            <attachments>
                <item>
                    <vpcId>vpc-4c090c2a</vpcId>
                    <state>attached</state>
                </item>
            </attachments>
            <amazonSideAsn>65001</amazonSideAsn>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>NYOffice</value>
                </item>
            </tagSet>
        </item>
    </vpnGatewaySet>
</DescribeVpnGatewaysResponse>
```

### Example 2

This example uses filters to describe any virtual private gateway you own
whose state is either `pending` or `available`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpnGateways
&Filter.1.Name=state
&Filter.1.Value.1=pending
&Filter.1.Value.2=available
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describevpngateways.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpngateways.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpnConnections

DetachClassicLinkVpc
