This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::Index EncryptionConfiguration

The encryption configuration for a vector bucket or index. By default, if you don't specify, all
new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed
keys (SSE-S3), specifically `AES256`. You can optionally override bucket level encryption settings, and set a specific encryption configuration for a vector index at the time of index creation.

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

AWS Key Management Service (KMS) customer managed key ID to use for the encryption
configuration. This parameter is allowed if and only if `sseType` is set to
`aws:kms`.

To specify the KMS key, you must use the format of the KMS key
Amazon Resource Name (ARN).

For example, specify Key ARN in the following format:
`arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[-a-z0-9]*:kms:[-a-z0-9]*:[0-9]{12}:key/.+)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseType`

The server-side encryption type to use for the encryption configuration of the vector
bucket. By default, if you don't specify, all new vectors in Amazon S3 vector buckets use
server-side encryption with Amazon S3 managed keys (SSE-S3), specifically
`AES256`.

_Required_: No

_Type_: String

_Allowed values_: `AES256 | aws:kms`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3Vectors::Index

MetadataConfiguration
