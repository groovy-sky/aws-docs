This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment SegmentGroups

Specifies the set of segment criteria to evaluate when handling segment groups for the
segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Groups" : [ Groups, ... ],
  "Include" : String
}

```

### YAML

```yaml

  Groups:
    - Groups
  Include: String

```

## Properties

`Groups`

Specifies the set of segment criteria to evaluate when handling segment groups for the
segment.

_Required_: No

_Type_: [Array](aws-properties-pinpoint-segment-groups.md) of [Groups](aws-properties-pinpoint-segment-groups.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Include`

Specifies how to handle multiple segment groups for the segment. For example, if the
segment includes three segment groups, whether the resulting segment includes endpoints
that match all, any, or none of the segment groups.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SegmentDimensions

SetDimension

All content copied from https://docs.aws.amazon.com/.
