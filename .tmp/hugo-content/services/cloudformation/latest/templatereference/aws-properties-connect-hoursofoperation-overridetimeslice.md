This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation OverrideTimeSlice

The start time or end time for an hours of operation override.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hours" : Integer,
  "Minutes" : Integer
}

```

### YAML

```yaml

  Hours: Integer
  Minutes: Integer

```

## Properties

`Hours`

The hours.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minutes`

The minutes.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `59`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HoursOfOperationTimeSlice

RecurrenceConfig

All content copied from https://docs.aws.amazon.com/.
