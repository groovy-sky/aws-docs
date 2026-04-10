This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment Recency

Specifies how recently segment members were active.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Duration" : String,
  "RecencyType" : String
}

```

### YAML

```yaml

  Duration: String
  RecencyType: String

```

## Properties

`Duration`

The duration to use when determining which users have been active or inactive with
your app.

Possible values: `HR_24` \| `DAY_7` \| `DAY_14` \|
`DAY_30`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecencyType`

The type of recency dimension to use for the segment. Valid values are:
`ACTIVE` and `INACTIVE`. If the value is `ACTIVE`,
the segment includes users who have used your app within the specified duration are
included in the segment. If the value is `INACTIVE`, the segment includes
users who haven't used your app within the specified duration are included in the
segment.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Location

SegmentDimensions

All content copied from https://docs.aws.amazon.com/.
