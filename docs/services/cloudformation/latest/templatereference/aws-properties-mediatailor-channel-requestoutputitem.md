This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel RequestOutputItem

The output configuration for this channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DashPlaylistSettings" : DashPlaylistSettings,
  "HlsPlaylistSettings" : HlsPlaylistSettings,
  "ManifestName" : String,
  "SourceGroup" : String
}

```

### YAML

```yaml

  DashPlaylistSettings:
    DashPlaylistSettings
  HlsPlaylistSettings:
    HlsPlaylistSettings
  ManifestName: String
  SourceGroup: String

```

## Properties

`DashPlaylistSettings`

DASH manifest configuration parameters.

_Required_: No

_Type_: [DashPlaylistSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-dashplaylistsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsPlaylistSettings`

HLS playlist configuration parameters.

_Required_: No

_Type_: [HlsPlaylistSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-channel-hlsplaylistsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

The name of the manifest for the channel. The name appears in the `PlaybackUrl`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceGroup`

A string used to match which `HttpPackageConfiguration` is used for each `VodSource`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LogConfigurationForChannel

SlateSource
