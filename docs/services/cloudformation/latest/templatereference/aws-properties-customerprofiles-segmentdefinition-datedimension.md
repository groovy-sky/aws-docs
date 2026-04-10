This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition DateDimension

Object that segments on various Customer Profile's date fields.

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

The action to segment on.

_Required_: Yes

_Type_: String

_Allowed values_: `BEFORE | AFTER | BETWEEN | NOT_BETWEEN | ON`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Values`

The values to apply the DimensionType on.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionOverrides

Dimension

All content copied from https://docs.aws.amazon.com/.
