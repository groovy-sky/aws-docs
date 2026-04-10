This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::VectorBucket EncryptionConfiguration

Specifies the encryption configuration for the vector bucket. By default, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "SseType" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  SseType: String

```

## Properties

`KmsKeyArn`

AWS Key Management Service (KMS) customer managed key ARN to use for the encryption configuration. This parameter is required if and only if `SseType` is set to `aws:kms`.

You must specify the full ARN of the KMS key. Key IDs or key aliases aren't supported.

###### Important

Amazon S3 Vectors only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in AWS KMS](../../../kms/latest/developerguide/symmetric-asymmetric.md) in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[-a-z0-9]*:kms:[-a-z0-9]*:[0-9]{12}:key/.+)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseType`

The server-side encryption type to use for the encryption configuration of the vector bucket. Valid values are `AES256` for Amazon S3 managed keys and `aws:kms` for AWS KMS keys.

_Required_: No

_Type_: String

_Allowed values_: `AES256 | aws:kms`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3Vectors::VectorBucket

Tag

All content copied from https://docs.aws.amazon.com/.
