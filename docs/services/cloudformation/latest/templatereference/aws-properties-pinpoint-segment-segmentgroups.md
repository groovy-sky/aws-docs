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

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-segment-groups.html) of [Groups](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-segment-groups.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Include`

Specifies how to handle multiple segment groups for the segment. For example, if the
segment includes three segment groups, whether the resulting segment includes endpoints
that match all, any, or none of the segment groups.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SegmentDimensions

SetDimension
