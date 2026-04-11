This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition ExtraLengthValueProfileDimension

Object that segments on various Customer profile's fields that are larger than
normal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionType" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  DimensionType: String
  Values:
    - String

```

## Properties

`DimensionType`

The action to segment with.

_Required_: Yes

_Type_: String

_Allowed values_: `INCLUSIVE | EXCLUSIVE | CONTAINS | BEGINS_WITH | ENDS_WITH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Values`

The values to apply the DimensionType on.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1000 | 50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

Group

All content copied from https://docs.aws.amazon.com/.
