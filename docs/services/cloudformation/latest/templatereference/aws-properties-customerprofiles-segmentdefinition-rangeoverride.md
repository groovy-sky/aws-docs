This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition RangeOverride

Overrides the original range on a calculated attribute definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "End" : Integer,
  "Start" : Integer,
  "Unit" : String
}

```

### YAML

```yaml

  End: Integer
  Start: Integer
  Unit: String

```

## Properties

`End`

The end time of when to include objects.

_Required_: No

_Type_: Integer

_Minimum_: `-2147483648`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Start`

The start time of when to include objects.

_Required_: Yes

_Type_: Integer

_Minimum_: `-2147483648`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Unit`

The unit for start and end.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProfileTypeDimension

SegmentGroup

All content copied from https://docs.aws.amazon.com/.
