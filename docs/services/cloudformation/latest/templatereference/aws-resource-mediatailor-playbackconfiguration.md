This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration

Adds a new playback configuration to AWS Elemental MediaTailor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::PlaybackConfiguration",
  "Properties" : {
      "AdConditioningConfiguration" : AdConditioningConfiguration,
      "AdDecisionServerConfiguration" : AdDecisionServerConfiguration,
      "AdDecisionServerUrl" : String,
      "AvailSuppression" : AvailSuppression,
      "Bumper" : Bumper,
      "CdnConfiguration" : CdnConfiguration,
      "ConfigurationAliases" : {Key: Value, ...},
      "DashConfiguration" : DashConfiguration,
      "HlsConfiguration" : HlsConfiguration,
      "InsertionMode" : String,
      "LivePreRollConfiguration" : LivePreRollConfiguration,
      "LogConfiguration" : LogConfiguration,
      "ManifestProcessingRules" : ManifestProcessingRules,
      "Name" : String,
      "PersonalizationThresholdSeconds" : Integer,
      "SlateAdUrl" : String,
      "Tags" : [ Tag, ... ],
      "TranscodeProfileName" : String,
      "VideoContentSourceUrl" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::PlaybackConfiguration
Properties:
  AdConditioningConfiguration:
    AdConditioningConfiguration
  AdDecisionServerConfiguration:
    AdDecisionServerConfiguration
  AdDecisionServerUrl: String
  AvailSuppression:
    AvailSuppression
  Bumper:
    Bumper
  CdnConfiguration:
    CdnConfiguration
  ConfigurationAliases:
    Key: Value
  DashConfiguration:
    DashConfiguration
  HlsConfiguration:
    HlsConfiguration
  InsertionMode: String
  LivePreRollConfiguration:
    LivePreRollConfiguration
  LogConfiguration:
    LogConfiguration
  ManifestProcessingRules:
    ManifestProcessingRules
  Name: String
  PersonalizationThresholdSeconds: Integer
  SlateAdUrl: String
  Tags:
    - Tag
  TranscodeProfileName: String
  VideoContentSourceUrl: String

```

## Properties

`AdConditioningConfiguration`

The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns, and what priority MediaTailor uses when inserting ads.

_Required_: No

_Type_: [AdConditioningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-adconditioningconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdDecisionServerConfiguration`

Property description not available.

_Required_: No

_Type_: [AdDecisionServerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-addecisionserverconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdDecisionServerUrl`

The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailSuppression`

The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).

_Required_: No

_Type_: [AvailSuppression](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-availsuppression.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bumper`

The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).

_Required_: No

_Type_: [Bumper](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-bumper.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CdnConfiguration`

The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.

_Required_: No

_Type_: [CdnConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationAliases`

The player parameters and aliases used as dynamic variables during session initialization. For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html).

_Required_: No

_Type_: Object

_Pattern_: `player_params\.\w+\Z`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DashConfiguration`

The configuration for a DASH source.

_Required_: No

_Type_: [DashConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-dashconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsConfiguration`

The configuration for HLS content.

_Required_: No

_Type_: [HlsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-hlsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsertionMode`

The setting that controls whether players can use stitched or guided ad insertion. The default, `STITCHED_ONLY`, forces all player sessions to use stitched (server-side) ad insertion. Choosing `PLAYER_SELECT` allows players to select either stitched or guided ad insertion at session-initialization time. The default for players that do not specify an insertion mode is stitched.

_Required_: No

_Type_: String

_Allowed values_: `STITCHED_ONLY | PLAYER_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LivePreRollConfiguration`

The configuration for pre-roll ad insertion.

_Required_: No

_Type_: [LivePreRollConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-liveprerollconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

Defines where AWS Elemental MediaTailor sends logs for the playback configuration.

_Required_: No

_Type_: [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestProcessingRules`

The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.

_Required_: No

_Type_: [ManifestProcessingRules](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-manifestprocessingrules.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The identifier for the playback configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PersonalizationThresholdSeconds`

Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to _ad replacement_ in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in AWS Elemental MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlateAdUrl`

The URL for a video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID playback configurations. For VPAID, the slate is required because MediaTailor provides it in the slots designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the playback configuration. Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TranscodeProfileName`

The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoContentSourceUrl`

The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DashConfiguration.ManifestEndpointPrefix`

The URL generated by MediaTailor to initiate a playback session. The session uses server-side reporting. This setting is ignored in PUT operations.

`HlsConfiguration.ManifestEndpointPrefix`

The URL that is used to initiate a playback session for devices that support Apple HLS. The session uses server-side reporting.

`PlaybackConfigurationArn`

The Amazon Resource Name (ARN) for the playback configuration.

`PlaybackEndpointPrefix`

The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.

`SessionInitializationEndpointPrefix`

The URL that the player uses to initialize a session that uses client-side reporting.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AdConditioningConfiguration
