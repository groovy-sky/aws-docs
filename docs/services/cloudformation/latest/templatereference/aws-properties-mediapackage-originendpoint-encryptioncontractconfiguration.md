This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint EncryptionContractConfiguration

Use `encryptionContractConfiguration` to configure one or more content
encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract
defines the content keys used to encrypt the audio and video tracks in your stream.
To configure the encryption contract, specify which audio and video encryption
presets to use. For more information about these presets, see [SPEKE Version 2.0 Presets](../../../mediapackage/latest/ug/drm-content-speke-v2-presets.md).

Note the following considerations when using
`encryptionContractConfiguration`:

- You can use `encryptionContractConfiguration` for DASH endpoints
that use SPEKE Version 2.0. SPEKE Version 2.0 relies on the CPIX Version 2.3 specification.

- You cannot combine an `UNENCRYPTED` preset with
`UNENCRYPTED` or `SHARED` presets across
`presetSpeke20Audio` and `presetSpeke20Video`.

- When you use a `SHARED` preset, you must use it for both
`presetSpeke20Audio` and `presetSpeke20Video`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PresetSpeke20Audio" : String,
  "PresetSpeke20Video" : String
}

```

### YAML

```yaml

  PresetSpeke20Audio: String
  PresetSpeke20Video: String

```

## Properties

`PresetSpeke20Audio`

A collection of audio encryption presets.

Value description:

- `PRESET-AUDIO-1` \- Use one content key to encrypt all of the
audio tracks in your stream.

- `PRESET-AUDIO-2` \- Use one content key to encrypt all of the
stereo audio tracks and one content key to encrypt all of the multichannel
audio tracks.

- `PRESET-AUDIO-3` \- Use one content key to encrypt all of the
stereo audio tracks, one content key to encrypt all of the multichannel audio
tracks with 3 to 6 channels, and one content key to encrypt all of the
multichannel audio tracks with more than 6 channels.

- `SHARED` \- Use the same content key for all of the audio and
video tracks in your stream.

- `UNENCRYPTED` \- Don't encrypt any of the audio tracks in your
stream.

_Required_: Yes

_Type_: String

_Allowed values_: `PRESET-AUDIO-1 | PRESET-AUDIO-2 | PRESET-AUDIO-3 | SHARED | UNENCRYPTED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PresetSpeke20Video`

A collection of video encryption presets.

Value description:

- `PRESET-VIDEO-1` \- Use one content key to encrypt all of the video
tracks in your stream.

- `PRESET-VIDEO-2` \- Use one content key to encrypt all of the SD
video tracks and one content key for all HD and higher resolutions video
tracks.

- `PRESET-VIDEO-3` \- Use one content key to encrypt all of the SD
video tracks, one content key for HD video tracks and one content key for all
UHD video tracks.

- `PRESET-VIDEO-4` \- Use one content key to encrypt all of the SD
video tracks, one content key for HD video tracks, one content key for all UHD1
video tracks and one content key for all UHD2 video tracks.

- `PRESET-VIDEO-5` \- Use one content key to encrypt all of the SD
video tracks, one content key for HD1 video tracks, one content key for HD2
video tracks, one content key for all UHD1 video tracks and one content key for
all UHD2 video tracks.

- `PRESET-VIDEO-6` \- Use one content key to encrypt all of the SD
video tracks, one content key for HD1 video tracks, one content key for HD2
video tracks and one content key for all UHD video tracks.

- `PRESET-VIDEO-7` \- Use one content key to encrypt all of the SD+HD1
video tracks, one content key for HD2 video tracks and one content key for all
UHD video tracks.

- `PRESET-VIDEO-8` \- Use one content key to encrypt all of the SD+HD1
video tracks, one content key for HD2 video tracks, one content key for all
UHD1 video tracks and one content key for all UHD2 video tracks.

- `SHARED` \- Use the same content key for all of the video and audio
tracks in your stream.

- `UNENCRYPTED` \- Don't encrypt any of the video tracks in your
stream.

_Required_: Yes

_Type_: String

_Allowed values_: `PRESET-VIDEO-1 | PRESET-VIDEO-2 | PRESET-VIDEO-3 | PRESET-VIDEO-4 | PRESET-VIDEO-5 | PRESET-VIDEO-6 | PRESET-VIDEO-7 | PRESET-VIDEO-8 | SHARED | UNENCRYPTED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashPackage

HlsEncryption

All content copied from https://docs.aws.amazon.com/.
