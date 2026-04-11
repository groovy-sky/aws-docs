# AssociatedRole

Information about the associated IAM roles.

## Contents

**associatedRoleArn**

The ARN of the associated IAM role.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**certificateS3BucketName**

The name of the Amazon S3 bucket in which the Amazon S3 object is stored.

Type: String

Required: No

**certificateS3ObjectKey**

The key of the Amazon S3 object where the certificate, certificate chain, and encrypted private key bundle
are stored. The object key is formatted as follows: `role_arn`/ `certificate_arn`.

Type: String

Required: No

**encryptionKmsKeyId**

The ID of the KMS key used to encrypt the private key.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associatedrole.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associatedrole.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associatedrole.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssignedPrivateIpAddress

AssociatedTargetNetwork
