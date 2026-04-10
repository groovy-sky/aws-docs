This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyStore KmsEncryptionSettings

A structure that contains the KMS encryption configuration for the policy store. The
encryption settings determine what customer-managed KMS key will be used to encrypt
all resources within the policy store, and any user-defined context key-value pairs to append
during encryption processes.

This data type is used as a field that is part of the [EncryptionSettings](../../../../reference/verifiedpermissions/latest/apireference/api-encryptionsettings.md)
type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionContext" : {Key: Value, ...},
  "Key" : String
}

```

### YAML

```yaml

  EncryptionContext:
    Key: Value
  Key: String

```

## Properties

`EncryptionContext`

User-defined, additional context to be added to encryption processes.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The customer-managed KMS key [Amazon Resource Name (ARN)](../../../../general/latest/gr/aws-arns-and-namespaces.md), alias or ID to be used for encryption processes.

Users can provide the full KMS key ARN, a KMS key alias, or a KMS key ID, but it will be
mapped to the full KMS key ARN after policy store creation, and referenced when
encrypting child resources.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:/_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionState

KmsEncryptionState

All content copied from https://docs.aws.amazon.com/.
