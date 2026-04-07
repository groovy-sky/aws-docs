# DisassociateEnclaveCertificateIamRole

Disassociates an IAM role from an AWS Certificate Manager (ACM) certificate. Disassociating an IAM role
from an ACM certificate removes the Amazon S3 object that contains the certificate, certificate chain, and
encrypted private key from the Amazon S3 bucket. It also revokes the IAM role's permission to use the
KMS key used to encrypt the private key. This effectively revokes the role's permission
to use the certificate.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CertificateArn**

The ARN of the ACM certificate from which to disassociate the IAM role.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RoleArn**

The ARN of the IAM role to disassociate.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisassociateEnclaveCertificateIamRole)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisassociateEnclaveCertificateIamRole)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisassociateEnclaveCertificateIamRole)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociateenclavecertificateiamrole.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateClientVpnTargetNetwork

DisassociateIamInstanceProfile
