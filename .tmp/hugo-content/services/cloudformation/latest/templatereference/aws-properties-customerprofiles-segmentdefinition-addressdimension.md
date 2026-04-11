This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition AddressDimension

Object that segments on Customer Profile's address object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "City" : ProfileDimension,
  "Country" : ProfileDimension,
  "County" : ProfileDimension,
  "PostalCode" : ProfileDimension,
  "Province" : ProfileDimension,
  "State" : ProfileDimension
}

```

### YAML

```yaml

  City:
    ProfileDimension
  Country:
    ProfileDimension
  County:
    ProfileDimension
  PostalCode:
    ProfileDimension
  Province:
    ProfileDimension
  State:
    ProfileDimension

```

## Properties

`City`

The city belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Country`

The country belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`County`

The county belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PostalCode`

The postal code belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Province`

The province belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`State`

The state belonging to the address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CustomerProfiles::SegmentDefinition

AttributeDimension

All content copied from https://docs.aws.amazon.com/.
