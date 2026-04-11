This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioSelector

Information about one audio to extract from the input.

The parent of this entity is InputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SelectorSettings" : AudioSelectorSettings
}

```

### YAML

```yaml

  Name: String
  SelectorSettings:
    AudioSelectorSettings

```

## Properties

`Name`

A name for this AudioSelector.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectorSettings`

Information about the specific audio to extract from the input.

_Required_: No

_Type_: [AudioSelectorSettings](aws-properties-medialive-channel-audioselectorsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioPidSelection

AudioSelectorSettings

All content copied from https://docs.aws.amazon.com/.
