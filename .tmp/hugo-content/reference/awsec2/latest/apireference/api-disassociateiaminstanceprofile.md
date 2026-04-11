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

Type: [IamInstanceProfileAssociation](api-iaminstanceprofileassociation.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disassociateiaminstanceprofile.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociateiaminstanceprofile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateEnclaveCertificateIamRole

DisassociateInstanceEventWindow
