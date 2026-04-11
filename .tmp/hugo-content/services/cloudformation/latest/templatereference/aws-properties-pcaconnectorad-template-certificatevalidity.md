This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template CertificateValidity

Information describing the end of the validity period of the certificate. This parameter
sets the “Not After” date for the certificate. Certificate validity is the period of time
during which a certificate is valid. Validity can be expressed as an explicit date and time
when the certificate expires, or as a span of time after issuance, stated in days, months,
or years. For more information, see Validity in RFC 5280. This value is unaffected when
ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the
future, the certificate will expire 20 days from issuance time regardless of the
ValidityNotBefore value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RenewalPeriod" : ValidityPeriod,
  "ValidityPeriod" : ValidityPeriod
}

```

### YAML

```yaml

  RenewalPeriod:
    ValidityPeriod
  ValidityPeriod:
    ValidityPeriod

```

## Properties

`RenewalPeriod`

Renewal period is the period of time before certificate expiration when a new
certificate will be requested.

_Required_: Yes

_Type_: [ValidityPeriod](aws-properties-pcaconnectorad-template-validityperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidityPeriod`

Information describing the end of the validity period of the certificate. This parameter
sets the “Not After” date for the certificate. Certificate validity is the period of time
during which a certificate is valid. Validity can be expressed as an explicit date and time
when the certificate expires, or as a span of time after issuance, stated in days, months,
or years. For more information, see Validity in RFC 5280. This value is unaffected when
ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the
future, the certificate will expire 20 days from issuance time regardless of the
ValidityNotBefore value.

_Required_: Yes

_Type_: [ValidityPeriod](aws-properties-pcaconnectorad-template-validityperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationPolicy

EnrollmentFlagsV2

All content copied from https://docs.aws.amazon.com/.
