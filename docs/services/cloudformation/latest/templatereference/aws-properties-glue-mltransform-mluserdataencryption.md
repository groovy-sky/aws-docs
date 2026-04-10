This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform MLUserDataEncryption

The encryption-at-rest settings of the transform that apply to accessing user data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "MLUserDataEncryptionMode" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  MLUserDataEncryptionMode: String

```

## Properties

`KmsKeyId`

The ID for the customer-provided KMS key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MLUserDataEncryptionMode`

The encryption mode applied to user data. Valid values are:

- DISABLED: encryption is disabled.

- SSEKMS: use of server-side encryption with AWS Key Management Service (SSE-KMS) for user data
stored in Amazon S3.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputRecordTables

TransformEncryption

All content copied from https://docs.aws.amazon.com/.
