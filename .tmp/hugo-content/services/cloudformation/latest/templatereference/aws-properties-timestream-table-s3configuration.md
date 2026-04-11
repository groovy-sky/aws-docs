This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table S3Configuration

The configuration that specifies an S3 location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "EncryptionOption" : String,
  "KmsKeyId" : String,
  "ObjectKeyPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  EncryptionOption: String
  KmsKeyId: String
  ObjectKeyPrefix: String

```

## Properties

`BucketName`

The bucket name of the customer S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionOption`

The encryption option for the customer S3 location. Options are S3 server-side encryption with an S3 managed key
or AWS managed key.

_Required_: Yes

_Type_: String

_Allowed values_: `SSE_S3 | SSE_KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The AWS KMS key ID for the customer S3 location when encrypting with an AWS managed
key.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectKeyPrefix`

The object key preview for the customer S3 location.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9|!\-_*'\(\)]([a-zA-Z0-9]|[!\-_*'\(\)\/.])+`

_Minimum_: `1`

_Maximum_: `928`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionProperties

Schema

All content copied from https://docs.aws.amazon.com/.
