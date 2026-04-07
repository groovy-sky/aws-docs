# CreateDelegateMacVolumeOwnershipTask

Delegates ownership of the Amazon EBS root volume for an Apple silicon
Mac instance to an administrative user.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring Idempotency](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the Amazon EC2 Mac instance.

Type: String

Required: Yes

**MacCredentials**

Specifies the following credentials:

- **Internal disk administrative user**

- **Username** \- Only the default administrative user
( `aws-managed-user`) is supported and it is used by default. You can't
specify a different administrative user.

- **Password** \- If you did not change the default
password for `aws-managed-user`, specify the default password, which is
_blank_. Otherwise, specify your password.

- **Amazon EBS root volume administrative user**

- **Username** \- If you did not change the default
administrative user, specify `ec2-user`. Otherwise, specify the username
for your administrative user.

- **Password** \- Specify the password for the
administrative user.

The credentials must be specified in the following JSON format:

`{
"internalDiskPassword":"internal-disk-admin_password",
"rootVolumeUsername":"root-volume-admin_username",
"rootVolumepassword":"root-volume-admin_password"
}`

Type: String

Required: Yes

**TagSpecification.N**

The tags to assign to the volume ownership delegation task.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**macModificationTask**

Information about the volume ownership delegation task.

Type: [MacModificationTask](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_MacModificationTask.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateDelegateMacVolumeOwnershipTask)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateDefaultVpc

CreateDhcpOptions
