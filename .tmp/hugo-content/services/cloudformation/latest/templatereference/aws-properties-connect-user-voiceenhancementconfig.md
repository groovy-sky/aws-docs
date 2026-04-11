This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User VoiceEnhancementConfig

Configuration settings for voice enhancement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Channel" : String,
  "VoiceEnhancementMode" : String
}

```

### YAML

```yaml

  Channel: String
  VoiceEnhancementMode: String

```

## Properties

`Channel`

The channel for this voice enhancement configuration. **Only `VOICE` is supported for this data type.**

_Required_: Yes

_Type_: String

_Allowed values_: `VOICE | CHAT | TASK | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceEnhancementMode`

The voice enhancement mode.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | VOICE_ISOLATION | NOISE_SUPPRESSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserProficiency

AWS::Connect::UserHierarchyGroup

All content copied from https://docs.aws.amazon.com/.
