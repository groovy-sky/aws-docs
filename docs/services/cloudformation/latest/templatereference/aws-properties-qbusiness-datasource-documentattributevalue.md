This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource DocumentAttributeValue

The value of a document attribute. You can only provide one value for a document
attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateValue" : String,
  "LongValue" : Number,
  "StringListValue" : [ String, ... ],
  "StringValue" : String
}

```

### YAML

```yaml

  DateValue: String
  LongValue: Number
  StringListValue:
    - String
  StringValue:
    String

```

## Properties

`DateValue`

A date expressed as an ISO 8601 string.

It's important for the time zone to be included in the ISO 8601 date-time format. For
example, 2012-03-25T12:30:10+01:00 is the ISO 8601 date-time format for March 25th 2012
at 12:30PM (plus 10 seconds) in Central European Time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LongValue`

A long integer value.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringListValue`

A list of strings.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

A string.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeTarget

DocumentEnrichmentConfiguration

All content copied from https://docs.aws.amazon.com/.
