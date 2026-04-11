This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition Group

Contains dimensions that determine what to segment on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ Dimension, ... ],
  "SourceSegments" : [ SourceSegment, ... ],
  "SourceType" : String,
  "Type" : String
}

```

### YAML

```yaml

  Dimensions:
    - Dimension
  SourceSegments:
    - SourceSegment
  SourceType: String
  Type: String

```

## Properties

`Dimensions`

Defines the attributes to segment on.

_Required_: No

_Type_: Array of [Dimension](aws-properties-customerprofiles-segmentdefinition-dimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceSegments`

Defines the starting source of data.

_Required_: No

_Type_: Array of [SourceSegment](aws-properties-customerprofiles-segmentdefinition-sourcesegment.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceType`

Defines how to interact with the source data.

_Required_: No

_Type_: String

_Allowed values_: `ALL | ANY | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

Defines how to interact with the profiles found in the current filtering.

_Required_: No

_Type_: String

_Allowed values_: `ALL | ANY | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtraLengthValueProfileDimension

ProfileAttributes

All content copied from https://docs.aws.amazon.com/.
