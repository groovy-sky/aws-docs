This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor BoundingBox

Identifies the bounding box around the label, face, text, or personal protective equipment.
The `left` (x-coordinate) and `top` (y-coordinate) are coordinates representing the top and
left sides of the bounding box. Note that the upper-left corner of the image is the origin
(0,0).

The `top` and `left` values returned are ratios of the overall
image size. For example, if the input image is 700x200 pixels, and the top-left coordinate of
the bounding box is 350x50 pixels, the API returns a `left` value of 0.5 (350/700)
and a `top` value of 0.25 (50/200).

The `width` and `height` values represent the dimensions of the
bounding box as a ratio of the overall image dimension. For example, if the input image is
700x200 pixels, and the bounding box width is 70 pixels, the width returned is 0.1. For more information, see
[BoundingBox](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_BoundingBox).

###### Note

The bounding box coordinates can have negative values. For example, if Amazon Rekognition is
able to detect a face that is at the image edge and is only partially visible, the service
can return coordinates that are outside the image bounds and, depending on the image edge,
you might get negative values or values greater than 1 for the `left` or
`top` values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Height" : Number,
  "Left" : Number,
  "Top" : Number,
  "Width" : Number
}

```

### YAML

```yaml

  Height: Number
  Left: Number
  Top: Number
  Width: Number

```

## Properties

`Height`

Height of the bounding box as a ratio of the overall image height.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Left`

Left coordinate of the bounding box as a ratio of overall image width.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Top`

Top coordinate of the bounding box as a ratio of overall image height.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Width`

Width of the bounding box as a ratio of the overall image width.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Rekognition::StreamProcessor

ConnectedHomeSettings
