This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition Dimension

Defines the attribute to segment on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CalculatedAttributes" : {Key: Value, ...},
  "ProfileAttributes" : ProfileAttributes
}

```

### YAML

```yaml

  CalculatedAttributes:
    Key: Value
  ProfileAttributes:
    ProfileAttributes

```

## Properties

`CalculatedAttributes`

Object that holds the calculated attributes to segment on.

_Required_: No

_Type_: Object of [CalculatedAttributeDimension](aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileAttributes`

Object that holds the profile attributes to segment on.

_Required_: No

_Type_: [ProfileAttributes](aws-properties-customerprofiles-segmentdefinition-profileattributes.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateDimension

ExtraLengthValueProfileDimension

All content copied from https://docs.aws.amazon.com/.
