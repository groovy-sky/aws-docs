This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionRectangle

Settings to configure the caption rectangle for an output captions that will be created using this Teletext
source captions.

The parent of this entity is TeletextSourceSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Height" : Number,
  "LeftOffset" : Number,
  "TopOffset" : Number,
  "Width" : Number
}

```

### YAML

```yaml

  Height: Number
  LeftOffset: Number
  TopOffset: Number
  Width: Number

```

## Properties

`Height`

See the description in leftOffset.

For height, specify the entire height of the rectangle as a percentage of the underlying frame height. For
example, \\"80\\" means the rectangle height is 80% of the underlying frame height. The topOffset and rectangleHeight
must add up to 100% or less. This field corresponds to tts:extent - Y in the TTML standard.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LeftOffset`

Applies only if you plan to convert these source captions to EBU-TT-D or TTML in an output. (Make sure to leave
the default if you don't have either of these formats in the output.) You can define a display rectangle for the
captions that is smaller than the underlying video frame. You define the rectangle by specifying the position of the
left edge, top edge, bottom edge, and right edge of the rectangle, all within the underlying video frame. The units
for the measurements are percentages. If you specify a value for one of these fields, you must specify a value for
all of them.

For leftOffset, specify the position of the left edge of the rectangle, as a percentage of the underlying frame
width, and relative to the left edge of the frame. For example, \\"10\\" means the measurement is 10% of the underlying
frame width. The rectangle left edge starts at that position from the left edge of the frame. This field corresponds
to tts:origin - X in the TTML standard.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopOffset`

See the description in leftOffset.

For topOffset, specify the position of the top edge of the rectangle, as a percentage of the underlying frame
height, and relative to the top edge of the frame. For example, \\"10\\" means the measurement is 10% of the underlying
frame height. The rectangle top edge starts at that position from the top edge of the frame. This field corresponds
to tts:origin - Y in the TTML standard.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

See the description in leftOffset.

For width, specify the entire width of the rectangle as a percentage of the underlying frame width. For
example, \\"80\\" means the rectangle width is 80% of the underlying frame width. The leftOffset and rectangleWidth
must add up to 100% or less. This field corresponds to tts:extent - X in the TTML standard.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptionLanguageMapping

CaptionSelector

All content copied from https://docs.aws.amazon.com/.
