# DescribeSecurityGroupReferences

Describes the VPCs on the other side of a VPC peering or Transit Gateway connection that are referencing the security groups you've specified in this request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupId.N**

The IDs of the security groups in your account.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**securityGroupReferenceSet**

Information about the VPCs with the referencing security groups.

Type: Array of [SecurityGroupReference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroupReference.html) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the security group references for sg-11aa22bb.
The
response indicates that this security group is referenced by a security group in VPC vpc-1a2b3c4d.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSecurityGroupReferences
&GroupId.1=sg-11aa22bb
&AUTHPARAMS
```

#### Sample Response

```

<DescribeSecurityGroupReferencesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>19744c88-baa2-45df-905f-example</requestId>
    <securityGroupReferenceSet>
        <item>
            <referencingVpcId>vpc-1a2b3c4d</referencingVpcId>
            <vpcPeeringConnectionId>pcx-b04deed9</vpcPeeringConnectionId>
            <groupId>sg-11aa22bb</groupId>
        </item>
    </securityGroupReferenceSet>
</DescribeSecurityGroupReferencesResponse>
```

### Example 2

This example describes the security group references for sg-11aa22bb and sg-1111aaaa.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSecurityGroupReferences
&GroupId.1=sg-11aa22bb
&GroupId.2=sg-1111aaaa
&AUTHPARAMS
```

#### Sample Response

```

<DescribeSecurityGroupReferencesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d1835dca-61c1-459d-99cb-example</requestId>
    <securityGroupReferenceSet>
        <item>
            <referencingVpcId>vpc-81326ae4</referencingVpcId>
            <vpcPeeringConnectionId>pcx-b04deed9</vpcPeeringConnectionId>
            <groupId>sg-11aa22bb</groupId>
        </item>
        <item>
            <referencingVpcId>vpc-1a2b3c4d</referencingVpcId>
            <vpcPeeringConnectionId>pcx-aabbccdd</vpcPeeringConnectionId>
            <groupId>sg-1111aaaa</groupId>
        </item>
    </securityGroupReferenceSet>
</DescribeSecurityGroupReferencesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSecurityGroupReferences)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeSecurityGroupReferences)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeSecondarySubnets

DescribeSecurityGroupRules
