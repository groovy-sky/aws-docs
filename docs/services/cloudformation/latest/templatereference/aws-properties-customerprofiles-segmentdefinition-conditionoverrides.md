This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition ConditionOverrides

An object to override the original condition block of a calculated attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Range" : RangeOverride
}

```

### YAML

```yaml

  Range:
    RangeOverride

```

## Properties

`Range`

The relative time period over which data is included in the aggregation for this
override.

_Required_: No

_Type_: [RangeOverride](aws-properties-customerprofiles-segmentdefinition-rangeoverride.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculatedAttributeDimension

DateDimension

All content copied from https://docs.aws.amazon.com/.
