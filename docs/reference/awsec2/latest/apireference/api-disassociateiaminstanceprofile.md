# DisassociateIamInstanceProfile

Disassociates an IAM instance profile from a running or stopped instance.

Use [DescribeIamInstanceProfileAssociations](api-describeiaminstanceprofileassociations.md) to get the association
ID.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId**

The ID of the IAM instance profile association.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**iamInstanceProfileAssociation**

Information about the IAM instance profile association.

Type: [IamInstanceProfileAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfileAssociation.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example disassociates the specified IAM instance profile
association.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisassociateIamInstanceProfile
&AssociationId=iip-assoc-08049da59357d598c
&AUTHPARAMS
```

#### Sample Response

```

<DisassociateIamInstanceProfileResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>4840f938-fc84-4791-8ae5-example</requestId>
    <iamInstanceProfileAssociation>
        <associationId>iip-assoc-08049da59357d598c</associationId>
        <iamInstanceProfile>
            <arn>arn:aws:iam::123456789012:instance-profile/AdminProfile</arn>
            <id>AIPAI5IVIHMFFYY2DKV5Y</id>
        </iamInstanceProfile>
        <instanceId>i-1234567890abcdef0</instanceId>
        <state>disassociating</state>
    </iamInstanceProfileAssociation>
</DisassociateIamInstanceProfileResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisassociateIamInstanceProfile)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DisassociateIamInstanceProfile)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DisassociateEnclaveCertificateIamRole

DisassociateInstanceEventWindow
