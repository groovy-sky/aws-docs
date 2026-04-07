# DisassociateSubnetCidrBlock

Disassociates a CIDR block from a subnet. Currently, you can disassociate an IPv6 CIDR block only. You must detach or delete all gateways and resources that are associated with the CIDR block before you can disassociate it.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId**

The association ID for the CIDR block.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipv6CidrBlockAssociation**

Information about the IPv6 CIDR block association.

Type: [SubnetIpv6CidrBlockAssociation](api-subnetipv6cidrblockassociation.md) object

**requestId**

The ID of the request.

Type: String

**subnetId**

The ID of the subnet.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example disassociates the IPv6 CIDR block from the subnet.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisassociateSubnetCidrBlock
&AssociationId=subnet-cidr-assoc-3aa54053
&AUTHPARAMS
```

#### Sample Response

```

<DisassociateSubnetCidrBlockResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <subnetId>subnet-5f46ec3b</subnetId>
    <ipv6CidrBlockAssociation>
        <ipv6CidrBlock>2001:db8:1234:1a00::/64</ipv6CidrBlock>
        <ipv6CidrBlockState>
            <state>disassociating</state>
        </ipv6CidrBlockState>
        <associationId>subnet-cidr-assoc-3aa54053</associationId>
    </ipv6CidrBlockAssociation>
</DisassociateSubnetCidrBlockResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisassociateSubnetCidrBlock)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisassociateSubnetCidrBlock)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociatesubnetcidrblock.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisassociateSubnetCidrBlock)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociatesubnetcidrblock.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateSecurityGroupVpc

DisassociateTransitGatewayMulticastDomain
