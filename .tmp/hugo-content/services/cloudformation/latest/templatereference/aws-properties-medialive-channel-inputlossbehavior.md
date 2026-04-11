This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputLossBehavior

The configuration of channel behavior when the input is lost.

The parent of this entity is GlobalConfiguration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlackFrameMsec" : Integer,
  "InputLossImageColor" : String,
  "InputLossImageSlate" : InputLocation,
  "InputLossImageType" : String,
  "RepeatFrameMsec" : Integer
}

```

### YAML

```yaml

  BlackFrameMsec: Integer
  InputLossImageColor: String
  InputLossImageSlate:
    InputLocation
  InputLossImageType: String
  RepeatFrameMsec: Integer

```

## Properties

`BlackFrameMsec`

On input loss, the number of milliseconds to substitute black into the output before switching to the frame
specified by inputLossImageType. A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted
as infinite.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossImageColor`

When the input loss image type is "color," this field specifies the color to use. Value: 6 hex characters that
represent the values of RGB.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossImageSlate`

When the input loss image type is "slate," these fields specify the parameters for accessing the slate.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossImageType`

Indicates whether to substitute a solid color or a slate into the output after the input loss exceeds
blackFrameMsec.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepeatFrameMsec`

On input loss, the number of milliseconds to repeat the previous picture before substituting black into the
output. A value x, where 0 <= x <= 1,000,000 and a value of 1,000,000, is interpreted as infinite.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLocation

InputLossFailoverSettings

All content copied from https://docs.aws.amazon.com/.
