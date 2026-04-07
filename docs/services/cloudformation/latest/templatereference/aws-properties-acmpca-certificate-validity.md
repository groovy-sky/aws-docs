This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate Validity

Length of time for which the certificate issued by your private certificate authority
(CA), or by the private CA itself, is valid in days, months, or years. You can issue a
certificate by calling the `IssueCertificate` operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Type: String
  Value: Number

```

## Properties

`Type`

Specifies whether the `Value` parameter represents days, months, or
years.

_Required_: Yes

_Type_: String

_Allowed values_: `END_DATE | ABSOLUTE | DAYS | MONTHS | YEARS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

A long integer interpreted according to the value of `Type`, below.

_Required_: Yes

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Subject

AWS::ACMPCA::CertificateAuthority
