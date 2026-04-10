This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashTtmlConfiguration

The settings for TTML subtitles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TtmlProfile" : String
}

```

### YAML

```yaml

  TtmlProfile: String

```

## Properties

`TtmlProfile`

The profile that MediaPackage uses when signaling subtitles in the manifest. `IMSC` is the default profile.
`EBU-TT-D` produces subtitles that are compliant with the EBU-TT-D TTML profile.
MediaPackage passes through subtitle styles to the manifest. For more information about EBU-TT-D subtitles, see [EBU-TT-D Subtitling Distribution Format](https://tech.ebu.ch/publications/tech3380).

_Required_: Yes

_Type_: String

_Allowed values_: `IMSC_1 | EBU_TT_D_101`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashSubtitleConfiguration

DashUtcTiming

All content copied from https://docs.aws.amazon.com/.
