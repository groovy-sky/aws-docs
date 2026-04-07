This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Rotation RecurrenceSettings

Information about when an on-call rotation is in effect and how long the rotation period
lasts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DailySettings" : [ String, ... ],
  "MonthlySettings" : [ MonthlySetting, ... ],
  "NumberOfOnCalls" : Integer,
  "RecurrenceMultiplier" : Integer,
  "ShiftCoverages" : [ ShiftCoverage, ... ],
  "WeeklySettings" : [ WeeklySetting, ... ]
}

```

### YAML

```yaml

  DailySettings:
    - String
  MonthlySettings:
    - MonthlySetting
  NumberOfOnCalls: Integer
  RecurrenceMultiplier: Integer
  ShiftCoverages:
    - ShiftCoverage
  WeeklySettings:
    - WeeklySetting

```

## Properties

`DailySettings`

Information about on-call rotations that recur daily.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonthlySettings`

Information about on-call rotations that recur monthly.

_Required_: No

_Type_: Array of [MonthlySetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmcontacts-rotation-monthlysetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfOnCalls`

The number of contacts, or shift team members designated to be on call concurrently
during a shift. For example, in an on-call schedule that contains ten contacts, a value of
`2` designates that two of them are on call at any given time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecurrenceMultiplier`

The number of days, weeks, or months a single rotation lasts.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShiftCoverages`

Information about the days of the week included in on-call rotation coverage.

_Required_: No

_Type_: Array of [ShiftCoverage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmcontacts-rotation-shiftcoverage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeeklySettings`

Information about on-call rotations that recur weekly.

_Required_: No

_Type_: Array of [WeeklySetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmcontacts-rotation-weeklysetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MonthlySetting

ShiftCoverage
