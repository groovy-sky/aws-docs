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

_Type_: [DashPlaylistSettings](aws-properties-mediatailor-channel-dashplaylistsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsPlaylistSettings`

HLS playlist configuration parameters.

_Required_: No

_Type_: [HlsPlaylistSettings](aws-properties-mediatailor-channel-hlsplaylistsettings.md)

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfigurationForChannel

SlateSource

All content copied from https://docs.aws.amazon.com/.
