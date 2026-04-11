This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template KeyUsagePropertyFlags

Specifies key usage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Decrypt" : Boolean,
  "KeyAgreement" : Boolean,
  "Sign" : Boolean
}

```

### YAML

```yaml

  Decrypt: Boolean
  KeyAgreement: Boolean
  Sign: Boolean

```

## Properties

`Decrypt`

Allows key for encryption and decryption.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyAgreement`

Allows key exchange without encryption.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sign`

Allow key use for digital signature.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyUsageProperty

PrivateKeyAttributesV2

All content copied from https://docs.aws.amazon.com/.
