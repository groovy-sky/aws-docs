# CreateVpnGateway

Creates a virtual private gateway. A virtual private gateway is the endpoint on the
VPC side of your VPN connection. You can create a virtual private gateway before
creating the VPC itself.

For more information, see [AWS Site-to-Site VPN](../../../../services/vpn/latest/s2svpn/vpc-vpn.md) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AmazonSideAsn**

A private Autonomous System Number (ASN) for the Amazon side of a BGP session. If
you're using a 16-bit ASN, it must be in the 64512 to 65534 range. If you're using a
32-bit ASN, it must be in the 4200000000 to 4294967294 range.

Default: 64512

Type: Long

Required: No

**AvailabilityZone**

The Availability Zone for the virtual private gateway.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to apply to the virtual private gateway.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Type**

The type of VPN connection this virtual private gateway supports.

Type: String

Valid Values: `ipsec.1`

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpnGateway**

Information about the virtual private gateway.

Type: [VpnGateway](api-vpngateway.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example creates a virtual private gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnGateway
&Type=ipsec.1
&AUTHPARAMS
```

#### Sample Response

```

<CreateVpnGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
  <vpnGateway>
        <vpnGatewayId>vgw-fe4aa197</vpnGatewayId>
        <state>available</state>
        <type>ipsec.1</type>
        <amazonSideAsn>64512</amazonSideAsn>
        <attachments/>
    </vpnGateway>
</CreateVpnGatewayResponse>
```

### Example 2

This example creates a virtual private gateway and specifies a private ASN of
65001 for the Amazon side of the gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnGateway
&Type=ipsec.1
&AmazonSideAsn=65001
&AUTHPARAMS
```

#### Sample Response

```

<CreateVpnGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>fe90b404-d4e5-4153-8677-31dexample</requestId>
    <vpnGateway>
        <vpnGatewayId>vgw-f74aa19e</vpnGatewayId>
        <state>available</state>
        <type>ipsec.1</type>
        <amazonSideAsn>65001</amazonSideAsn>
        <attachments/>
    </vpnGateway>
</CreateVpnGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createvpngateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createvpngateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVpnConnectionRoute

DeleteCapacityManagerDataExport
