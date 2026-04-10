This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate Extensions

Contains X.509 extension information for a certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificatePolicies" : [ PolicyInformation, ... ],
  "CustomExtensions" : [ CustomExtension, ... ],
  "ExtendedKeyUsage" : [ ExtendedKeyUsage, ... ],
  "KeyUsage" : KeyUsage,
  "SubjectAlternativeNames" : [ GeneralName, ... ]
}

```

### YAML

```yaml

  CertificatePolicies:
    - PolicyInformation
  CustomExtensions:
    - CustomExtension
  ExtendedKeyUsage:
    - ExtendedKeyUsage
  KeyUsage:
    KeyUsage
  SubjectAlternativeNames:
    - GeneralName

```

## Properties

`CertificatePolicies`

Contains a sequence of one or more policy information terms, each of which consists of
an object identifier (OID) and optional qualifiers. For more information, see NIST's
definition of [Object\
Identifier (OID)](https://csrc.nist.gov/glossary/term/Object_Identifier).

In an end-entity certificate, these terms indicate the policy under which the
certificate was issued and the purposes for which it may be used. In a CA certificate,
these terms limit the set of policies for certification paths that include this
certificate.

_Required_: No

_Type_: Array of [PolicyInformation](aws-properties-acmpca-certificate-policyinformation.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomExtensions`

Contains a sequence of one or more X.509 extensions, each of which consists of an
object identifier (OID), a base64-encoded value, and the critical flag. For more
information, see the [Global OID reference\
database.](https://oidref.com/2.5.29)

_Required_: No

_Type_: Array of [CustomExtension](aws-properties-acmpca-certificate-customextension.md)

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExtendedKeyUsage`

Specifies additional purposes for which the certified public key may be used other
than basic purposes indicated in the `KeyUsage` extension.

_Required_: No

_Type_: [Array](aws-properties-acmpca-certificate-extendedkeyusage.md) of [ExtendedKeyUsage](aws-properties-acmpca-certificate-extendedkeyusage.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyUsage`

Defines one or more purposes for which the key contained in the certificate can be
used. Default value for each option is false.

_Required_: No

_Type_: [KeyUsage](aws-properties-acmpca-certificate-keyusage.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubjectAlternativeNames`

The subject alternative name extension allows identities to be bound to the subject of
the certificate. These identities may be included in addition to or in place of the
identity in the subject field of the certificate.

_Required_: No

_Type_: Array of [GeneralName](aws-properties-acmpca-certificate-generalname.md)

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtendedKeyUsage

GeneralName

All content copied from https://docs.aws.amazon.com/.
