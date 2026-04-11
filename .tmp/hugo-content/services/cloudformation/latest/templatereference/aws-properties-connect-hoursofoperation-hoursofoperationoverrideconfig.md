This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation HoursOfOperationOverrideConfig

Information about the hours of operation override config: day, start time, and end time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Day" : String,
  "EndTime" : OverrideTimeSlice,
  "StartTime" : OverrideTimeSlice
}

```

### YAML

```yaml

  Day: String
  EndTime:
    OverrideTimeSlice
  StartTime:
    OverrideTimeSlice

```

## Properties

`Day`

The day that the hours of operation override applies to.

_Required_: Yes

_Type_: String

_Allowed values_: `SUNDAY | MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTime`

The end time that your contact center closes if overrides are applied.

_Required_: Yes

_Type_: [OverrideTimeSlice](aws-properties-connect-hoursofoperation-overridetimeslice.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time when your contact center opens if overrides are applied.

_Required_: Yes

_Type_: [OverrideTimeSlice](aws-properties-connect-hoursofoperation-overridetimeslice.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HoursOfOperationOverride

HoursOfOperationsIdentifier

All content copied from https://docs.aws.amazon.com/.
