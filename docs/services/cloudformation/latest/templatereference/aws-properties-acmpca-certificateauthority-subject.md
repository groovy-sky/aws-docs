This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority Subject

ASN1 subject for the certificate authority.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommonName" : String,
  "Country" : String,
  "CustomAttributes" : [ CustomAttribute, ... ],
  "DistinguishedNameQualifier" : String,
  "GenerationQualifier" : String,
  "GivenName" : String,
  "Initials" : String,
  "Locality" : String,
  "Organization" : String,
  "OrganizationalUnit" : String,
  "Pseudonym" : String,
  "SerialNumber" : String,
  "State" : String,
  "Surname" : String,
  "Title" : String
}

```

### YAML

```yaml

  CommonName: String
  Country: String
  CustomAttributes:
    - CustomAttribute
  DistinguishedNameQualifier: String
  GenerationQualifier: String
  GivenName: String
  Initials: String
  Locality: String
  Organization: String
  OrganizationalUnit: String
  Pseudonym: String
  SerialNumber: String
  State: String
  Surname: String
  Title: String

```

## Properties

`CommonName`

Fully qualified domain name (FQDN) associated with the certificate subject.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Country`

Two-digit code that specifies the country in which the certificate subject
located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomAttributes`

Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of
which consists of an object identifier (OID) and a value. For more information, see
NIST’s definition of [Object Identifier (OID)](https://csrc.nist.gov/glossary/term/Object_Identifier).

###### Note

Custom attributes cannot be used in combination with standard attributes.

_Required_: No

_Type_: Array of [CustomAttribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-acmpca-certificateauthority-customattribute.html)

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DistinguishedNameQualifier`

Disambiguating information for the certificate subject.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GenerationQualifier`

Typically a qualifier appended to the name of an individual. Examples include Jr. for
junior, Sr. for senior, and III for third.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GivenName`

First name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Initials`

Concatenation that typically contains the first letter of the GivenName, the first
letter of the middle name if one exists, and the first letter of the SurName.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Locality`

The locality (such as a city or town) in which the certificate subject is
located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Organization`

Legal name of the organization with which the certificate subject is
affiliated.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationalUnit`

A subdivision or unit of the organization (such as sales or finance) with which the
certificate subject is affiliated.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Pseudonym`

Typically a shortened version of a longer GivenName. For example, Jonathan is often
shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SerialNumber`

The certificate serial number.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`State`

State in which the subject of the certificate is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Surname`

Family name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Title`

A personal title such as Mr.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RevocationConfiguration

Tag
