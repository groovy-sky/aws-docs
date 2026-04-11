This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FailoverConditionSettings

Settings for one failover condition.

The parent of this entity is FailoverCondition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioSilenceSettings" : AudioSilenceFailoverSettings,
  "InputLossSettings" : InputLossFailoverSettings,
  "VideoBlackSettings" : VideoBlackFailoverSettings
}

```

### YAML

```yaml

  AudioSilenceSettings:
    AudioSilenceFailoverSettings
  InputLossSettings:
    InputLossFailoverSettings
  VideoBlackSettings:
    VideoBlackFailoverSettings

```

## Properties

`AudioSilenceSettings`

MediaLive will perform a failover if the specified audio selector is silent for the specified period.

_Required_: No

_Type_: [AudioSilenceFailoverSettings](aws-properties-medialive-channel-audiosilencefailoversettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossSettings`

MediaLive will perform a failover if content is not detected in this input for the specified period.

_Required_: No

_Type_: [InputLossFailoverSettings](aws-properties-medialive-channel-inputlossfailoversettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoBlackSettings`

MediaLive will perform a failover if content is considered black for the specified period.

_Required_: No

_Type_: [VideoBlackFailoverSettings](aws-properties-medialive-channel-videoblackfailoversettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailoverCondition

FeatureActivations

All content copied from https://docs.aws.amazon.com/.
