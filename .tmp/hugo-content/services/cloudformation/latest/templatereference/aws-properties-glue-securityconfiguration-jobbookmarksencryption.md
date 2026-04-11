This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SecurityConfiguration JobBookmarksEncryption

Specifies how job bookmark data should be encrypted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobBookmarksEncryptionMode" : String,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  JobBookmarksEncryptionMode: String
  KmsKeyArn: String

```

## Properties

`JobBookmarksEncryptionMode`

The encryption mode to use for job bookmarks data.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | CSE-KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[a-z0-9-]*:kms:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

S3Encryption

All content copied from https://docs.aws.amazon.com/.
