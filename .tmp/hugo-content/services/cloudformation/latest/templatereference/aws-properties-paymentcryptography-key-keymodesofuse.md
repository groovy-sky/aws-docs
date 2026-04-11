This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PaymentCryptography::Key KeyModesOfUse

The list of cryptographic operations that you can perform using the key. The modes of
use are deﬁned in section A.5.3 of the TR-31 spec.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Decrypt" : Boolean,
  "DeriveKey" : Boolean,
  "Encrypt" : Boolean,
  "Generate" : Boolean,
  "NoRestrictions" : Boolean,
  "Sign" : Boolean,
  "Unwrap" : Boolean,
  "Verify" : Boolean,
  "Wrap" : Boolean
}

```

### YAML

```yaml

  Decrypt: Boolean
  DeriveKey: Boolean
  Encrypt: Boolean
  Generate: Boolean
  NoRestrictions: Boolean
  Sign: Boolean
  Unwrap: Boolean
  Verify: Boolean
  Wrap: Boolean

```

## Properties

`Decrypt`

Speciﬁes whether an AWS Payment Cryptography key can be used to decrypt data.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeriveKey`

Speciﬁes whether an AWS Payment Cryptography key can be used to derive new keys.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypt`

Speciﬁes whether an AWS Payment Cryptography key can be used to encrypt data.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Generate`

Speciﬁes whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN
verification keys.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoRestrictions`

Speciﬁes whether an AWS Payment Cryptography key has no special restrictions other than the
restrictions implied by `KeyUsage`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sign`

Speciﬁes whether an AWS Payment Cryptography key can be used for signing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unwrap`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Verify`

Speciﬁes whether an AWS Payment Cryptography key can be used to verify signatures.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Wrap`

Speciﬁes whether an AWS Payment Cryptography key can be used to wrap other keys.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyAttributes

ReplicationStatusType

All content copied from https://docs.aws.amazon.com/.
