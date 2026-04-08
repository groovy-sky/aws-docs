# AttachClassicLinkVpc

###### Note

This action is deprecated.

Links an EC2-Classic instance to a ClassicLink-enabled VPC through one or more of the VPC
security groups. You cannot link an EC2-Classic instance to more than one VPC at a time. You
can only link an instance that's in the `running` state. An instance is
automatically unlinked from a VPC when it's stopped - you can link it to the VPC again when
you restart it.

After you've linked an instance, you cannot change the VPC security groups that are associated with it. To change the security groups, you must first unlink the instance, and then link it again.

Linking your instance to a VPC is sometimes referred to as _attaching_ your instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the EC2-Classic instance.

Type: String

Required: Yes

**SecurityGroupId.N**

The IDs of the security groups. You cannot specify security groups from a different VPC.

Type: Array of strings

Required: Yes

**VpcId**

The ID of the ClassicLink-enabled VPC.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/attachclassiclinkvpc.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/attachclassiclinkvpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateVpcCidrBlock

AttachInternetGateway
