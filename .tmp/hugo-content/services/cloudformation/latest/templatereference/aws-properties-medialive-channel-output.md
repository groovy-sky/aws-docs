This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Output

The output settings.

The parent of this entity is OutputGroup.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioDescriptionNames" : [ String, ... ],
  "CaptionDescriptionNames" : [ String, ... ],
  "OutputName" : String,
  "OutputSettings" : OutputSettings,
  "VideoDescriptionName" : String
}

```

### YAML

```yaml

  AudioDescriptionNames:
    - String
  CaptionDescriptionNames:
    - String
  OutputName: String
  OutputSettings:
    OutputSettings
  VideoDescriptionName: String

```

## Properties

`AudioDescriptionNames`

The names of the audio descriptions that are used as audio sources for this output.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionDescriptionNames`

The names of the caption descriptions that are used as captions sources for this output.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputName`

The name that is used to identify an output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSettings`

The output type-specific settings.

_Required_: No

_Type_: [OutputSettings](aws-properties-medialive-channel-outputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoDescriptionName`

The name of the VideoDescription that is used as the source for this output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NielsenWatermarksSettings

OutputDestination

All content copied from https://docs.aws.amazon.com/.
