This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig EncryptionConfig

The encryption configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionType" : String,
  "KeyId" : String
}

```

### YAML

```yaml

  EncryptionType: String
  KeyId: String

```

## Properties

`EncryptionType`

The type of encryption.

_Required_: Yes

_Type_: String

_Allowed values_: `KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyId`

The full ARN of the encryption key.

###### Note

Be sure to provide the full ARN of the encryption key, not just the ID.

Amazon Connect supports only KMS keys with the default key spec of [`SYMMETRIC_DEFAULT`](../../../kms/latest/developerguide/asymmetric-key-specs.md#key-spec-symmetric-default).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::InstanceStorageConfig

KinesisFirehoseConfig

All content copied from https://docs.aws.amazon.com/.
