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

_Type_: [Range](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-range.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The threshold for the calculated attribute.

_Required_: No

_Type_: [Threshold](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-threshold.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributeItem

Range
