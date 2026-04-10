This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Insight DateFilter

A date filter for querying findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateRange" : DateRange,
  "End" : String,
  "Start" : String
}

```

### YAML

```yaml

  DateRange:
    DateRange
  End: String
  Start: String

```

## Properties

`DateRange`

A date range for the date filter.

_Required_: No

_Type_: [DateRange](aws-properties-securityhub-insight-daterange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`End`

A timestamp that provides the end date for the date filter.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](../../../../reference/securityhub/1-0/apireference/welcome.md#timestamps).

_Required_: No

_Type_: String

_Pattern_: `^([\+-]?\d{4}(?!\d{2}))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([tT]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

A timestamp that provides the start date for the date filter.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](../../../../reference/securityhub/1-0/apireference/welcome.md#timestamps).

_Required_: No

_Type_: String

_Pattern_: `^([\+-]?\d{4}(?!\d{2}))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([tT]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BooleanFilter

DateRange

All content copied from https://docs.aws.amazon.com/.
