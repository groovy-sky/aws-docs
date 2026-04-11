This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment Groups

An array that defines the set of segment criteria to evaluate when handling segment
groups for the segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ SegmentDimensions, ... ],
  "SourceSegments" : [ SourceSegments, ... ],
  "SourceType" : String,
  "Type" : String
}

```

### YAML

```yaml

  Dimensions:
    - SegmentDimensions
  SourceSegments:
    - SourceSegments
  SourceType: String
  Type: String

```

## Properties

`Dimensions`

An array that defines the dimensions to include or exclude from the segment.

_Required_: No

_Type_: Array of [SegmentDimensions](aws-properties-pinpoint-segment-segmentdimensions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceSegments`

The base segment to build the segment on. A base segment, also called a
_source segment_, defines the initial population of endpoints for
a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base
segment by using the dimensions that you specify.

You can specify more than one dimensional segment or only one imported segment. If you
specify an imported segment, the segment size estimate that displays on the Amazon
Pinpoint console indicates the size of the imported segment without any filters applied
to it.

_Required_: No

_Type_: [Array](aws-properties-pinpoint-segment-sourcesegments.md) of [SourceSegments](aws-properties-pinpoint-segment-sourcesegments.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

Specifies how to handle multiple base segments for the segment. For example, if you
specify three base segments for the segment, whether the resulting segment is based on
all, any, or none of the base segments.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies how to handle multiple dimensions for the segment. For example, if you
specify three dimensions for the segment, whether the resulting segment includes
endpoints that match all, any, or none of the dimensions.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GPSPoint

Location

All content copied from https://docs.aws.amazon.com/.
