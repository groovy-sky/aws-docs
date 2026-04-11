# AttachVpnGateway

Attaches an available virtual private gateway to a VPC. You can attach one virtual private
gateway to one VPC at a time.

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

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

**VpnGatewayId**

The ID of the virtual private gateway.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**attachment**

Information about the attachment.

Type: [VpcAttachment](api-vpcattachment.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example attaches the virtual private gateway with the ID
`vgw-8db04f81` to the VPC with the ID
`vpc-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AttachVpnGateway
&VpnGatewayId=vgw-8db04f81
&VpcId=vpc-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<AttachVpnGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <attachment>
      <vpcId>vpc-1a2b3c4d</vpcId>
      <state>attaching</state>
   </attachment>
</AttachVpnGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/attachvpngateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/attachvpngateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttachVolume

AuthorizeClientVpnIngress
