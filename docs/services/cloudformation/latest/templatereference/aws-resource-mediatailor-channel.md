This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel

The configuration parameters for a channel. For information about MediaTailor channels, see [Working with channels](https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html) in the _MediaTailor User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::Channel",
  "Properties" : {
      "Audiences" : [ String, ... ],
      "ChannelName" : String,
      "FillerSlate" : SlateSource,
      "LogConfiguration" : LogConfigurationForChannel,
      "Outputs" : [ RequestOutputItem, ... ],
      "PlaybackMode" : String,
      "Tags" : [ Tag, ... ],
      "Tier" : String,
      "TimeShiftConfiguration" : TimeShiftConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::Channel
Properties:
  Audiences:
    - String
  ChannelName: String
  FillerSlate:
    SlateSource
  LogConfiguration:
    LogConfigurationForChannel
  Outputs:
    - RequestOutputItem
  PlaybackMode: String
  Tags:
    - Tag
  Tier: String
  TimeShiftConfiguration:
    TimeShiftConfiguration

```

## Properties

`Audiences`

The list of audiences defined in channel.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelName`

The name of the channel.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FillerSlate`

The slate used to fill gaps between programs in the schedule. You must configure filler slate if your channel uses the `LINEAR` `PlaybackMode`. MediaTailor doesn't support filler slate for channels using the `LOOP` `PlaybackMode`.

_Required_: No

_Type_: [SlateSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-slatesource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The log configuration.

_Required_: No

_Type_: [LogConfigurationForChannel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-logconfigurationforchannel.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

The channel's output properties.

_Required_: Yes

_Type_: Array of [RequestOutputItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-requestoutputitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlaybackMode`

The type of playback mode for this channel.

`LINEAR` \- Programs play back-to-back only once.

`LOOP` \- Programs play back-to-back in an endless loop. When the last program in the schedule plays, playback loops back to the first program in the schedule.

_Required_: Yes

_Type_: String

_Allowed values_: `LOOP | LINEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the channel. Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The tier for this channel. STANDARD tier channels can contain live programs.

_Required_: No

_Type_: String

_Allowed values_: `BASIC | STANDARD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeShiftConfiguration`

The configuration for time-shifted viewing.

_Required_: No

_Type_: [TimeShiftConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-timeshiftconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Elemental MediaTailor

DashPlaylistSettings
