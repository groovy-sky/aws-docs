# EnableEbsEncryptionByDefault

Enables EBS encryption by default for your account in the current Region.

After you enable encryption by default, the EBS volumes that you create are
always encrypted, either using the default KMS key or the KMS key that you specified
when you created each volume. For more information, see [Amazon EBS encryption](../../../../services/ebs/latest/userguide/ebs-encryption.md) in the
_Amazon EBS User Guide_.

You can specify the default KMS key for encryption by default using
[ModifyEbsDefaultKmsKeyId](api-modifyebsdefaultkmskeyid.md) or [ResetEbsDefaultKmsKeyId](api-resetebsdefaultkmskeyid.md).

Enabling encryption by default has no effect on the encryption status of your
existing volumes.

After you enable encryption by default, you can no longer launch instances
using instance types that do not support encryption. For more information, see [Supported\
instance types](../../../../services/ebs/latest/userguide/ebs-encryption-requirements.md#ebs-encryption_supported_instances).

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**ebsEncryptionByDefault**

The updated status of encryption by default.

Type: Boolean

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/enableebsencryptionbydefault.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enableebsencryptionbydefault.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableCapacityManager

EnableFastLaunch
