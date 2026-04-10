This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition Range

The relative time period over which data is included in the aggregation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimestampFormat" : String,
  "TimestampSource" : String,
  "Unit" : String,
  "Value" : Integer,
  "ValueRange" : ValueRange
}

```

### YAML

```yaml

  TimestampFormat: String
  TimestampSource: String
  Unit: String
  Value: Integer
  ValueRange:
    ValueRange

```

## Properties

`TimestampFormat`

The format the timestamp field in your JSON object is specified. This value should be
one of EPOCHMILLI (for Unix epoch timestamps with second/millisecond level precision) or
ISO\_8601 (following ISO\_8601 format with second/millisecond level precision, with an
optional offset of Z or in the format HH:MM or HHMM.). E.g. if your object type is
MyType and source JSON is {"generatedAt": {"timestamp":
"2001-07-04T12:08:56.235-0700"}}, then TimestampFormat should be "ISO\_8601"

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimestampSource`

An expression specifying the field in your JSON object from which the date should be
parsed. The expression should follow the structure of \\"{ObjectTypeName.<Location of
timestamp field in JSON pointer format>}\\". E.g. if your object type is MyType and
source JSON is {"generatedAt": {"timestamp": "1737587945945"}}, then TimestampSource
should be "{MyType.generatedAt.timestamp}"

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Unit`

The unit of time.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The amount of time of the specified unit.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueRange`

A structure letting customers specify a relative time window over which over which
data is included in the Calculated Attribute. Use positive numbers to indicate that the
endpoint is in the past, and negative numbers to indicate it is in the future.
ValueRange overrides Value.

_Required_: No

_Type_: [ValueRange](aws-properties-customerprofiles-calculatedattributedefinition-valuerange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Conditions

Readiness

All content copied from https://docs.aws.amazon.com/.
