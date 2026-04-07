# AssociateSubnetCidrBlock

Associates a CIDR block with your subnet. You can only associate a single IPv6 CIDR
block with your subnet.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Ipv6CidrBlock**

The IPv6 CIDR block for your subnet.

Type: String

Required: No

**Ipv6IpamPoolId**

An IPv6 IPAM pool ID.

Type: String

Required: No

**Ipv6NetmaskLength**

An IPv6 netmask length.

Type: Integer

Required: No

**SubnetId**

The ID of your subnet.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipv6CidrBlockAssociation**

Information about the IPv6 association.

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

This example associates IPv6 CIDR block
`2001:db8:1234:1a00::/64` with subnet
`subnet-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateSubnetCidrBlock
&SubnetId=subnet-1a2b3c4d
&Ipv6CidrBlock=2001:db8:1234:1a00::/64
&AUTHPARAMS
```

#### Sample Response

```

<AssociateSubnetCidrBlock xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <subnetId>vpc-1a2b3c4d</subnetId>
   <ipv6CidrBlockAssociation>
        <ipv6CidrBlock>2001:db8:1234:1a00::/64</ipv6CidrBlock>
        <ipv6CidrBlockState>
            <state>associating</state>
        </ipv6CidrBlockState>
        <associationId>subnet-cidr-assoc-3aa54053</associationId>
   </ipv6CidrBlockAssociation>
</AssociateSubnetCidrBlock>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateSubnetCidrBlock)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AssociateSubnetCidrBlock)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssociateSecurityGroupVpc

AssociateTransitGatewayMulticastDomain
