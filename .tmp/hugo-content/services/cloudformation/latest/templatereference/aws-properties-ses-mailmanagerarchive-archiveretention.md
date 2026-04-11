This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerArchive ArchiveRetention

The retention policy for an email archive that specifies how long emails are kept
before being automatically deleted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetentionPeriod" : String
}

```

### YAML

```yaml

  RetentionPeriod: String

```

## Properties

`RetentionPeriod`

The enum value sets the period for retaining emails in an archive.

_Required_: Yes

_Type_: String

_Allowed values_: `THREE_MONTHS | SIX_MONTHS | NINE_MONTHS | ONE_YEAR | EIGHTEEN_MONTHS | TWO_YEARS | THIRTY_MONTHS | THREE_YEARS | FOUR_YEARS | FIVE_YEARS | SIX_YEARS | SEVEN_YEARS | EIGHT_YEARS | NINE_YEARS | TEN_YEARS | PERMANENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MailManagerArchive

Tag

All content copied from https://docs.aws.amazon.com/.
