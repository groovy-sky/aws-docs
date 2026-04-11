This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::DirectoryBucket ServerSideEncryptionByDefault

Describes the default server-side encryption to apply to new objects in the bucket. If a
PUT Object request doesn't specify any server-side encryption, this default encryption will
be applied. For more information, see [PutBucketEncryption](../../../s3/latest/api/restbucketputencryption.md) in
the _Amazon S3 API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSMasterKeyID" : String,
  "SSEAlgorithm" : String
}

```

### YAML

```yaml

  KMSMasterKeyID: String
  SSEAlgorithm: String

```

## Properties

`KMSMasterKeyID`

AWS Key Management Service (KMS) customer managed key ID to use for the
default encryption. This parameter is allowed only if `SSEAlgorithm` is set to
`aws:kms`.

You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the
KMS key. You can’t use the key alias of the KMS key.

- Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`

- Key ARN:
`arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`

If you are using encryption with cross-account or AWS service operations,
you must use a fully qualified KMS key ARN. For more information, see [Using encryption for cross-account operations](../../../s3/latest/userguide/s3-express-bucket-encryption.md#s3-express-bucket-encryption-update-bucket-policy).

###### Note

Your SSE-KMS configuration can only support 1 [customer managed\
key](../../../kms/latest/developerguide/concepts.md#customer-cmk) per directory bucket for the lifetime of the bucket. [AWS managed key](../../../kms/latest/developerguide/concepts.md#aws-managed-cmk) ( `aws/s3`) isn't supported. Also,
after you specify a customer managed key for SSE-KMS and upload objects with this
configuration, you can't override the customer managed key for your SSE-KMS
configuration. To use a new customer manager key for your data, we recommend copying
your existing objects to a new directory bucket with a new customer managed key.

###### Important

Amazon S3 only supports symmetric encryption KMS keys. For more information, see
[Asymmetric keys in AWS KMS](../../../kms/latest/developerguide/symmetric-asymmetric.md) in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSEAlgorithm`

Server-side encryption algorithm to use for the default encryption.

###### Note

For directory buckets, there are only two supported values for server-side encryption: `AES256` and `aws:kms`.

_Required_: Yes

_Type_: String

_Allowed values_: `aws:kms | AES256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

ServerSideEncryptionRule

All content copied from https://docs.aws.amazon.com/.
