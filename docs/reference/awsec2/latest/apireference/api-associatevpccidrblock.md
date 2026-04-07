# AssociateVpcCidrBlock

Associates a CIDR block with your VPC. You can associate a secondary IPv4 CIDR block,
an Amazon-provided IPv6 CIDR block, or an IPv6 CIDR block from an IPv6 address pool that
you provisioned through bring your own IP addresses ( [BYOIP](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)).

You must specify one of the following in the request: an IPv4 CIDR block, an IPv6
pool, or an Amazon-provided IPv6 CIDR block.

For more information about associating CIDR blocks with your VPC and applicable
restrictions, see [IP addressing for your VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-ip-addressing.html)
in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AmazonProvidedIpv6CidrBlock**

Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You
cannot specify the range of IPv6 addresses or the size of the CIDR block.

Type: Boolean

Required: No

**CidrBlock**

An IPv4 CIDR block to associate with the VPC.

Type: String

Required: No

**Ipv4IpamPoolId**

Associate a CIDR allocated from an IPv4 IPAM pool to a VPC. For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

Type: String

Required: No

**Ipv4NetmaskLength**

The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

Type: Integer

Required: No

**Ipv6CidrBlock**

An IPv6 CIDR block from the IPv6 address pool. You must also specify `Ipv6Pool` in the request.

To let Amazon choose the IPv6 CIDR block for you, omit this parameter.

Type: String

Required: No

**Ipv6CidrBlockNetworkBorderGroup**

The name of the location from which we advertise the IPV6 CIDR block. Use this parameter
to limit the CIDR block to this location.

You must set `AmazonProvidedIpv6CidrBlock` to `true` to use this parameter.

You can have one IPv6 CIDR block association per network border group.

Type: String

Required: No

**Ipv6IpamPoolId**

Associates a CIDR allocated from an IPv6 IPAM pool to a VPC. For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

Type: String

Required: No

**Ipv6NetmaskLength**

The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

Type: Integer

Required: No

**Ipv6Pool**

The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.

Type: String

Required: No

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**cidrBlockAssociation**

Information about the IPv4 CIDR block association.

Type: [VpcCidrBlockAssociation](api-vpccidrblockassociation.md) object

**ipv6CidrBlockAssociation**

Information about the IPv6 CIDR block association.

Type: [VpcIpv6CidrBlockAssociation](api-vpcipv6cidrblockassociation.md) object

**requestId**

The ID of the request.

Type: String

**vpcId**

The ID of the VPC.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example associates an IPv6 CIDR block with VPC `vpc-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateVpcCidrBlock
&VpcId=vpc-1a2b3c4d
&AmazonProvidedIpv6CidrBlock=true
&AUTHPARAMS
```

#### Sample Response

```

<AssociateVpcCidrBlock xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <ipv6CidrBlockAssociation>
        <ipv6CidrBlockState>
            <state>associating</state>
        </ipv6CidrBlockState>
        <associationId>vpc-cidr-assoc-e2a5408b</associationId>
    </ipv6CidrBlockAssociation>
   <vpcId>vpc-1a2b3c4d</vpcId>
</AssociateVpcCidrBlock>
```

### Example 2

This example associates the IPv4 CIDR block `10.2.0.0/16` with
VPC `vpc-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateVpcCidrBlock
&VpcId=vpc-1a2b3c4d
&CidrBlock=10.2.0.0/16
&AUTHPARAMS
```

#### Sample Response

```

<AssociateVpcCidrBlockResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>33af6c54-1139-4d50-b4f7-15a8example</requestId>
    <vpcId>vpc-1a2b3c4d</vpcId>
    <cidrBlockAssociation>
        <associationId>vpc-cidr-assoc-0280ab6b</associationId>
        <cidrBlock>10.2.0.0/16</cidrBlock>
        <cidrBlockState>
            <state>associating</state>
        </cidrBlockState>
    </cidrBlockAssociation>
</AssociateVpcCidrBlockResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateVpcCidrBlock)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AssociateVpcCidrBlock)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssociateTrunkInterface

AttachClassicLinkVpc
