This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate KeyUsage

Defines one or more purposes for which the key contained in the certificate can be
used. Default value for each option is false.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CRLSign" : Boolean,
  "DataEncipherment" : Boolean,
  "DecipherOnly" : Boolean,
  "DigitalSignature" : Boolean,
  "EncipherOnly" : Boolean,
  "KeyAgreement" : Boolean,
  "KeyCertSign" : Boolean,
  "KeyEncipherment" : Boolean,
  "NonRepudiation" : Boolean
}

```

### YAML

```yaml

  CRLSign: Boolean
  DataEncipherment: Boolean
  DecipherOnly: Boolean
  DigitalSignature: Boolean
  EncipherOnly: Boolean
  KeyAgreement: Boolean
  KeyCertSign: Boolean
  KeyEncipherment: Boolean
  NonRepudiation: Boolean

```

## Properties

`CRLSign`

Key can be used to sign CRLs.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataEncipherment`

Key can be used to decipher data.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DecipherOnly`

Key can be used only to decipher data.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DigitalSignature`

Key can be used for digital signing.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncipherOnly`

Key can be used only to encipher data.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyAgreement`

Key can be used in a key-agreement protocol.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyCertSign`

Key can be used to sign certificates.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyEncipherment`

Key can be used to encipher data.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NonRepudiation`

Key can be used for non-repudiation.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeneralName

OtherName

All content copied from https://docs.aws.amazon.com/.
