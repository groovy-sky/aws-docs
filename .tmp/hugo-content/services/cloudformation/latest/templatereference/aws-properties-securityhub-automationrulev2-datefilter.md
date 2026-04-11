This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 DateFilter

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

_Type_: [DateRange](aws-properties-securityhub-automationrulev2-daterange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`End`

A timestamp that provides the end date for the date filter.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](../../../../reference/securityhub/1-0/apireference/welcome.md#timestamps).

_Required_: No

_Type_: String

_Pattern_: `^(\d\d\d\d)-([0][1-9]|[1][0-2])-([0][1-9]|[1-2](\d)|[3][0-1])[T](?:([0-1](\d)|[2][0-3]):[0-5](\d):[0-5](\d)|23:59:60)(?:\.(\d)+)?([Z]|[+-](\d\d)(:?(\d\d))?)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

A timestamp that provides the start date for the date filter.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](../../../../reference/securityhub/1-0/apireference/welcome.md#timestamps).

_Required_: No

_Type_: String

_Pattern_: `^(\d\d\d\d)-([0][1-9]|[1][0-2])-([0][1-9]|[1-2](\d)|[3][0-1])[T](?:([0-1](\d)|[2][0-3]):[0-5](\d):[0-5](\d)|23:59:60)(?:\.(\d)+)?([Z]|[+-](\d\d)(:?(\d\d))?)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Criteria

DateRange

All content copied from https://docs.aws.amazon.com/.
