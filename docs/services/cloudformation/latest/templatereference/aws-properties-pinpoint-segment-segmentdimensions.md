This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment SegmentDimensions

Specifies the dimension settings for a segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : Json,
  "Behavior" : Behavior,
  "Demographic" : Demographic,
  "Location" : Location,
  "Metrics" : Json,
  "UserAttributes" : Json
}

```

### YAML

```yaml

  Attributes: Json
  Behavior:
    Behavior
  Demographic:
    Demographic
  Location:
    Location
  Metrics: Json
  UserAttributes: Json

```

## Properties

`Attributes`

One or more custom attributes to use as criteria for the segment. For more information see [AttributeDimension](https://docs.aws.amazon.com/pinpoint/latest/apireference/apps-application-id-segments.html#apps-application-id-segments-model-attributedimension)

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Behavior`

The behavior-based criteria, such as how recently users have used your app,
for the segment.

_Required_: No

_Type_: [Behavior](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-segment-behavior.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Demographic`

The demographic-based criteria, such as device platform, for the
segment.

_Required_: No

_Type_: [Demographic](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-segment-demographic.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

The location-based criteria, such as region or GPS coordinates, for the
segment.

_Required_: No

_Type_: [Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-segment-location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

One or more custom metrics to use as criteria for the segment.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttributes`

One or more custom user attributes to use as criteria for the segment.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Recency

SegmentGroups
