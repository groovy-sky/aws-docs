This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionSelector

Information about one caption to extract from the input.

The parent of this entity is InputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LanguageCode" : String,
  "Name" : String,
  "SelectorSettings" : CaptionSelectorSettings
}

```

### YAML

```yaml

  LanguageCode: String
  Name: String
  SelectorSettings:
    CaptionSelectorSettings

```

## Properties

`LanguageCode`

When specified, this field indicates the three-letter language code of the captions track to extract from the
source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name identifier for a captions selector. This name is used to associate this captions selector with one or
more captions descriptions. Names must be unique within a channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectorSettings`

Information about the specific audio to extract from the input.

_Required_: No

_Type_: [CaptionSelectorSettings](aws-properties-medialive-channel-captionselectorsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptionRectangle

CaptionSelectorSettings

All content copied from https://docs.aws.amazon.com/.
