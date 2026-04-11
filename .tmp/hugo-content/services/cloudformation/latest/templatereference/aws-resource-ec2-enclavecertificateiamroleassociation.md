This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EnclaveCertificateIamRoleAssociation

Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate.
This enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave. For more
information, see [AWS Certificate Manager for Nitro Enclaves](../../../enclaves/latest/user/nitro-enclave-refapp.md) in the _AWS Nitro Enclaves_
_User Guide_.

When the IAM role is associated with the ACM certificate, the certificate, certificate chain, and encrypted
private key are placed in an Amazon S3 location that only the associated IAM role can access. The private key of the certificate
is encrypted with an AWS managed key that has an attached attestation-based key policy.

To enable the IAM role to access the Amazon S3 object, you must grant it permission to call `s3:GetObject`
on the Amazon S3 bucket returned by the command. To enable the IAM role to access the KMS key,
you must grant it permission to call `kms:Decrypt` on the KMS key returned by the command.
For more information, see [Grant the role permission to access the certificate and encryption key](../../../enclaves/latest/user/nitro-enclave-refapp.md#add-policy) in the
_AWS Nitro Enclaves User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::EnclaveCertificateIamRoleAssociation",
  "Properties" : {
      "CertificateArn" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::EnclaveCertificateIamRoleAssociation
Properties:
  CertificateArn: String
  RoleArn: String

```

## Properties

`CertificateArn`

The ARN of the ACM certificate with which to associate the IAM role.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$`

_Minimum_: `1`

_Maximum_: `1283`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM
certificate.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$`

_Minimum_: `1`

_Maximum_: `1283`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IAM role and ACM certificate association.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CertificateS3BucketName`

The name of the Amazon S3 bucket to which the certificate was uploaded.

`CertificateS3ObjectKey`

The Amazon S3 object key where the certificate, certificate chain, and encrypted private
key bundle are stored. The object key is formatted as follows:
`role_arn`/ `certificate_arn`.

`EncryptionKmsKeyId`

The ID of the AWS KMS key used to encrypt the private key of the
certificate.

## Examples

### Associating an IAM role with an ACM certificate

The following example associates IAM role
`arn:aws:iam::123456789012:role/my-acm-role` with ACM certificate
`arn:aws:acm:us-east-1:123456789012:certificate/123abcde-cdef-abcd-1234-123abEXAMPLE`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "myCertAssociation",
    "Resources": {
        "MyEnclaveCertificateIamRoleAssociation": {
            "Type": "AWS::EC2::EnclaveCertificateIamRoleAssociation",
            "Properties": {
        "CertificateArn": "arn:aws:acm:us-east-1:123456789012:certificate/123abcde-cdef-abcd-1234-123abEXAMPLE",
        "RoleArn": "arn:aws:iam::123456789012:role/my-acm-role"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myCertAssociation:
    Type: AWS::EC2::EnclaveCertificateIamRoleAssociation
    Properties:
      CertificateArn:
        "arn:aws:acm:us-east-1:123456789012:certificate/123abcde-cdef-abcd-1234-123abEXAMPLE"
      RoleArn:
        "arn:aws:iam::123456789012:role/my-acm-role"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::EIPAssociation

AWS::EC2::FlowLog

All content copied from https://docs.aws.amazon.com/.
