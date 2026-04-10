This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition Conditions

The conditions including range, object count, and threshold for the calculated
attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectCount" : Integer,
  "Range" : Range,
  "Threshold" : Threshold
}

```

### YAML

```yaml

  ObjectCount: Integer
  Range:
    Range
  Threshold:
    Threshold

```

## Properties

`ObjectCount`

The number of profile objects used for the calculated attribute.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Range`

The relative time period over which data is included in the aggregation.

_Required_: No

_Type_: [Range](aws-properties-customerprofiles-calculatedattributedefinition-range.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The threshold for the calculated attribute.

_Required_: No

_Type_: [Threshold](aws-properties-customerprofiles-calculatedattributedefinition-threshold.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeItem

Range

All content copied from https://docs.aws.amazon.com/.
