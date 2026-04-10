This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition SegmentGroup

Contains all groups of the segment definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Groups" : [ Group, ... ],
  "Include" : String
}

```

### YAML

```yaml

  Groups:
    - Group
  Include: String

```

## Properties

`Groups`

Holds the list of groups within the segment definition.

_Required_: No

_Type_: Array of [Group](aws-properties-customerprofiles-segmentdefinition-group.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Include`

Defines whether to include or exclude the profiles that fit the segment
criteria.

_Required_: No

_Type_: String

_Allowed values_: `ALL | ANY | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RangeOverride

SegmentSort

All content copied from https://docs.aws.amazon.com/.
