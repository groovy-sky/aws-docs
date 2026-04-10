This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Signer::SigningProfile SignatureValidityPeriod

The validity period for the signing job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Type: String
  Value: Integer

```

## Properties

`Type`

The time unit for signature validity: DAYS \| MONTHS \| YEARS.

_Required_: No

_Type_: String

_Allowed values_: `DAYS | MONTHS | YEARS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The numerical value of the time unit for signature validity.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Signer::SigningProfile

Tag

All content copied from https://docs.aws.amazon.com/.
