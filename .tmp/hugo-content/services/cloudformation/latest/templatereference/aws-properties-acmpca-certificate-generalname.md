This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate GeneralName

Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280). Only one of
the following naming options should be provided. Providing more than one option results
in an `InvalidArgsException` error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DirectoryName" : Subject,
  "DnsName" : String,
  "EdiPartyName" : EdiPartyName,
  "IpAddress" : String,
  "OtherName" : OtherName,
  "RegisteredId" : String,
  "Rfc822Name" : String,
  "UniformResourceIdentifier" : String
}

```

### YAML

```yaml

  DirectoryName:
    Subject
  DnsName: String
  EdiPartyName:
    EdiPartyName
  IpAddress: String
  OtherName:
    OtherName
  RegisteredId: String
  Rfc822Name: String
  UniformResourceIdentifier: String

```

## Properties

`DirectoryName`

Contains information about the certificate subject. The certificate can be one issued
by your private certificate authority (CA) or it can be your private CA certificate. The
Subject field in the certificate identifies the entity that owns or controls the public
key in the certificate. The entity can be a user, computer, device, or service. The
Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative
distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN
must be unique for each entity, but your private CA can issue more than one certificate
with the same DN to the same entity.

_Required_: No

_Type_: [Subject](aws-properties-acmpca-certificate-subject.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DnsName`

Represents `GeneralName` as a DNS name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EdiPartyName`

Represents `GeneralName` as an `EdiPartyName` object.

_Required_: No

_Type_: [EdiPartyName](aws-properties-acmpca-certificate-edipartyname.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddress`

Represents `GeneralName` as an IPv4 or IPv6 address.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `39`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OtherName`

Represents `GeneralName` using an `OtherName` object.

_Required_: No

_Type_: [OtherName](aws-properties-acmpca-certificate-othername.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegisteredId`

Represents `GeneralName` as an object identifier (OID).

_Required_: No

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rfc822Name`

Represents `GeneralName` as an [RFC 822](https://datatracker.ietf.org/doc/html/rfc822) email
address.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UniformResourceIdentifier`

Represents `GeneralName` as a URI.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Extensions

KeyUsage

All content copied from https://docs.aws.amazon.com/.
