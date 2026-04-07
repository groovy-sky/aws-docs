# AssociateEnclaveCertificateIamRole

Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate.
This enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave. For more
information, see [AWS Certificate Manager for Nitro Enclaves](../../../../services/enclaves/latest/user/nitro-enclave-refapp.md) in the _AWS Nitro Enclaves_
_User Guide_.

When the IAM role is associated with the ACM certificate, the certificate, certificate chain, and encrypted
private key are placed in an Amazon S3 location that only the associated IAM role can access. The private key of the certificate
is encrypted with an AWS managed key that has an attached attestation-based key policy.

To enable the IAM role to access the Amazon S3 object, you must grant it permission to call `s3:GetObject`
on the Amazon S3 bucket returned by the command. To enable the IAM role to access the KMS key,
you must grant it permission to call `kms:Decrypt` on the KMS key returned by the command.
For more information, see [Grant the role permission to access the certificate and encryption key](../../../../services/enclaves/latest/user/nitro-enclave-refapp.md#add-policy) in the
_AWS Nitro Enclaves User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CertificateArn**

The ARN of the ACM certificate with which to associate the IAM role.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RoleArn**

The ARN of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM
certificate.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**certificateS3BucketName**

The name of the Amazon S3 bucket to which the certificate was uploaded.

Type: String

**certificateS3ObjectKey**

The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored. The
object key is formatted as follows: `role_arn`/ `certificate_arn`.

Type: String

**encryptionKmsKeyId**

The ID of the KMS key used to encrypt the private key of the certificate.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateEnclaveCertificateIamRole)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateEnclaveCertificateIamRole)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associateenclavecertificateiamrole.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateEnclaveCertificateIamRole)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associateenclavecertificateiamrole.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateDhcpOptions

AssociateIamInstanceProfile
