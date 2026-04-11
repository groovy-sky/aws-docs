This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template ValidityPeriod

Information describing the end of the validity period of the certificate. This parameter
sets the “Not After” date for the certificate. Certificate validity is the period of time
during which a certificate is valid. Validity can be expressed as an explicit date and time
when the certificate expires, or as a span of time after issuance, stated in hours, days,
months, or years. For more information, see Validity in RFC 5280. This value is unaffected
when ValidityNotBefore is also specified. For example, if Validity is set to 20 days in the
future, the certificate will expire 20 days from issuance time regardless of the
ValidityNotBefore value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Period" : Number,
  "PeriodType" : String
}

```

### YAML

```yaml

  Period: Number
  PeriodType: String

```

## Properties

`Period`

The numeric value for the validity period.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `8766000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodType`

The unit of time. You can select hours, days, weeks, months, and years.

_Required_: Yes

_Type_: String

_Allowed values_: `HOURS | DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateV4

AWS::PCAConnectorAD::TemplateGroupAccessControlEntry

All content copied from https://docs.aws.amazon.com/.
