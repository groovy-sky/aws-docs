This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SecurityConfiguration S3Encryption

Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "S3EncryptionMode" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  S3EncryptionMode: String

```

## Properties

`KmsKeyArn`

The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[a-z0-9-]*:kms:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3EncryptionMode`

The encryption mode to use for Amazon S3 data.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | SSE-KMS | SSE-S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobBookmarksEncryption

AWS::Glue::Table

All content copied from https://docs.aws.amazon.com/.
