# CreateCustomerGateway

Provides information to AWS about your customer gateway device. The
customer gateway device is the appliance at your end of the VPN connection. You
must provide the IP address of the customer gateway device’s external
interface. The IP address must be static and can be behind a device performing network
address translation (NAT).

For devices that use Border Gateway Protocol (BGP), you can also provide the device's
BGP Autonomous System Number (ASN). You can use an existing ASN assigned to your network.
If you don't have an ASN already, you can use a private ASN. For more information, see
[Customer gateway \
options for your Site-to-Site VPN connection](../../../../services/vpn/latest/s2svpn/cgw-options.md) in the _AWS Site-to-Site VPN User Guide_.

To create more than one customer gateway with the same VPN type, IP address, and
BGP ASN, specify a unique device name for each customer gateway. An identical request
returns information about the existing customer gateway; it doesn't create a new customer
gateway.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**BgpAsn**

For customer gateway devices that support BGP, specify the device's ASN. You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647`, you must use `BgpAsnExtended`.

Default: 65000

Valid values: `1` to `2,147,483,647`

Type: Integer

Required: No

**BgpAsnExtended**

For customer gateway devices that support BGP, specify the device's ASN. You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647`, you must use `BgpAsnExtended`.

Valid values: `2,147,483,648` to `4,294,967,295`

Type: Long

Required: No

**CertificateArn**

The Amazon Resource Name (ARN) for the customer gateway certificate.

Type: String

Required: No

**DeviceName**

A name for the customer gateway device.

Length Constraints: Up to 255 characters.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**IpAddress**

The IP address for the customer gateway device's outside interface. The address must be
static. If `OutsideIpAddressType` in your VPN connection options is set to
`PrivateIpv4`, you can use an RFC6598 or RFC1918 private IPv4 address. If
`OutsideIpAddressType` is set to `Ipv6`, you can use an IPv6 address.

Type: String

Required: No

**PublicIp**

_This member has been deprecated._ The Internet-routable IP address for the customer gateway's outside interface. The
address must be static.

Type: String

Required: No

**TagSpecification.N**

The tags to apply to the customer gateway.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Type**

The type of VPN connection that this customer gateway supports
( `ipsec.1`).

Type: String

Valid Values: `ipsec.1`

Required: Yes

## Response Elements

The following elements are returned by the service.

**customerGateway**

Information about the customer gateway.

Type: [CustomerGateway](api-customergateway.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a customer gateway with the name `my-device`,
the IP address `12.1.2.3`, and BGP ASN `65534`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateCustomerGateway
&Type=ipsec.1
&IpAddress=12.1.2.3
&BgpAsn=65534
&DeviceName=my-device
&AUTHPARAMS
```

#### Sample Response

```

<CreateCustomerGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <customerGateway>
      <customerGatewayId>cgw-11223344556677abc</customerGatewayId>
      <state>pending</state>
      <type>ipsec.1</type>
      <ipAddress>12.1.2.3</ipAddress>
      <bgpAsn>65534</bgpAsn>
      <deviceName>my-device</deviceName>
      <tagSet/>
   </customerGateway>
</CreateCustomerGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateCustomerGateway)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateCustomerGateway)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createcustomergateway.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateCustomerGateway)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createcustomergateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateCoipPool

CreateDefaultSubnet
