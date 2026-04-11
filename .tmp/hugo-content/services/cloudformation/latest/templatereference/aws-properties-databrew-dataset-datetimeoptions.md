This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset DatetimeOptions

Represents additional options for correct interpretation of datetime parameters used
in the Amazon S3 path of a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Format" : String,
  "LocaleCode" : String,
  "TimezoneOffset" : String
}

```

### YAML

```yaml

  Format: String
  LocaleCode: String
  TimezoneOffset: String

```

## Properties

`Format`

Required option, that defines the datetime format used for a date parameter in the
Amazon S3 path. Should use only supported datetime specifiers and
separation characters, all litera a-z or A-Z character should be escaped with single
quotes. E.g. "MM.dd.yyyy-'at'-HH:mm".

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleCode`

Optional value for a non-US locale code, needed for correct interpretation of some
date formats.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_\.#@\-]+$`

_Minimum_: `2`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimezoneOffset`

Optional value for a timezone offset of the datetime parameter value in the Amazon S3 path. Shouldn't be used if Format for this parameter includes timezone
fields. If no offset specified, UTC is assumed.

_Required_: No

_Type_: String

_Pattern_: `^(Z|[-+](\d|\d{2}|\d{2}:?\d{2}))$`

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetParameter

ExcelOptions

All content copied from https://docs.aws.amazon.com/.
