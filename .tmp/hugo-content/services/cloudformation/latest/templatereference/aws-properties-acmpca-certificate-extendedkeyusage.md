This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate ExtendedKeyUsage

Specifies additional purposes for which the certified public key may be used other
than basic purposes indicated in the `KeyUsage` extension.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExtendedKeyUsageObjectIdentifier" : String,
  "ExtendedKeyUsageType" : String
}

```

### YAML

```yaml

  ExtendedKeyUsageObjectIdentifier: String
  ExtendedKeyUsageType: String

```

## Properties

`ExtendedKeyUsageObjectIdentifier`

Specifies a custom `ExtendedKeyUsage` with an object identifier
(OID).

_Required_: No

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExtendedKeyUsageType`

Specifies a standard `ExtendedKeyUsage` as defined as in [RFC\
5280](https://datatracker.ietf.org/doc/html/rfc5280).

_Required_: No

_Type_: String

_Allowed values_: `SERVER_AUTH | CLIENT_AUTH | CODE_SIGNING | EMAIL_PROTECTION | TIME_STAMPING | OCSP_SIGNING | SMART_CARD_LOGIN | DOCUMENT_SIGNING | CERTIFICATE_TRANSPARENCY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EdiPartyName

Extensions

All content copied from https://docs.aws.amazon.com/.
