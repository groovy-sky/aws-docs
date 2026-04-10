This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template KeyUsageFlags

The key usage flags represent the purpose (e.g., encipherment, signature) of the key
contained in the certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataEncipherment" : Boolean,
  "DigitalSignature" : Boolean,
  "KeyAgreement" : Boolean,
  "KeyEncipherment" : Boolean,
  "NonRepudiation" : Boolean
}

```

### YAML

```yaml

  DataEncipherment: Boolean
  DigitalSignature: Boolean
  KeyAgreement: Boolean
  KeyEncipherment: Boolean
  NonRepudiation: Boolean

```

## Properties

`DataEncipherment`

DataEncipherment is asserted when the subject public key is used for directly
enciphering raw user data without the use of an intermediate symmetric cipher.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DigitalSignature`

The digitalSignature is asserted when the subject public key is used for verifying
digital signatures.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyAgreement`

KeyAgreement is asserted when the subject public key is used for key agreement.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyEncipherment`

KeyEncipherment is asserted when the subject public key is used for enciphering private
or secret keys, i.e., for key transport.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NonRepudiation`

NonRepudiation is asserted when the subject public key is used to verify digital
signatures.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyUsage

KeyUsageProperty

All content copied from https://docs.aws.amazon.com/.
