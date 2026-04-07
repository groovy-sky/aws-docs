# AssociateIamInstanceProfile

Associates an IAM instance profile with a running or stopped instance. You cannot
associate more than one IAM instance profile with an instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**IamInstanceProfile**

The IAM instance profile.

Type: [IamInstanceProfileSpecification](api-iaminstanceprofilespecification.md) object

Required: Yes

**InstanceId**

The ID of the instance.

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

This example associates the IAM instance profile with the specified
instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateIamInstanceProfile
&InstanceId=i-1234567890abcdef0
&IamInstanceProfile.Name=AdminProfile
&AUTHPARAMS
```

#### Sample Response

```

<AssociateIamInstanceProfileResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>e10deeaf-7cda-48e7-950b-example</requestId>
    <iamInstanceProfileAssociation>
        <associationId>iip-assoc-0750e3af14e2b40ad</associationId>
        <iamInstanceProfile>
            <arn>arn:aws:iam::123456789012:instance-profile/AdminProfile</arn>
            <id>AIPAJEDNCAA64SSD265D6</id>
        </iamInstanceProfile>
        <instanceId>i-1234567890abcdef0</instanceId>
        <state>associating</state>
    </iamInstanceProfileAssociation>
</AssociateIamInstanceProfileResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateIamInstanceProfile)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateIamInstanceProfile)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associateiaminstanceprofile.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateIamInstanceProfile)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associateiaminstanceprofile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateEnclaveCertificateIamRole

AssociateInstanceEventWindow
