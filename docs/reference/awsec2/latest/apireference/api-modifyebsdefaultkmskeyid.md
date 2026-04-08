# ModifyEbsDefaultKmsKeyId

Changes the default AWS KMS key for EBS encryption by default for your account in this Region.

AWS creates a unique AWS managed KMS key in each Region for use with encryption by default. If
you change the default KMS key to a symmetric customer managed KMS key, it is used instead of the AWS
managed KMS key. Amazon EBS does not support asymmetric KMS keys.

To reset the default KMS key to the AWS managed KMS key for EBS, use [ResetEbsDefaultKmsKeyId](api-resetebsdefaultkmskeyid.md).

If you delete or disable the customer managed KMS key that you specified for use with
encryption by default, your instances will fail to launch.

For more information, see [Amazon EBS encryption](../../../../services/ebs/latest/userguide/ebs-encryption.md)
in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**KmsKeyId**

The identifier of the AWS KMS key to use for Amazon EBS encryption.
If this parameter is not specified, your AWS KMS key for Amazon EBS is used. If `KmsKeyId` is
specified, the encrypted state must be `true`.

You can specify the KMS key using any of the following:

- Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.

- Key alias. For example, alias/ExampleAlias.

- Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.

- Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.

AWS authenticates the KMS key asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid,
the action can appear to complete, but eventually fails.

Amazon EBS does not support asymmetric KMS keys.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**kmsKeyId**

The Amazon Resource Name (ARN) of the default KMS key for encryption by default.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyebsdefaultkmskeyid.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyDefaultCreditSpecification

ModifyFleet
