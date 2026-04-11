This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashDvbFontDownload

For use with DVB-DASH profiles only. The settings for font downloads that you want AWS Elemental MediaPackage to pass through to the manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FontFamily" : String,
  "MimeType" : String,
  "Url" : String
}

```

### YAML

```yaml

  FontFamily: String
  MimeType: String
  Url: String

```

## Properties

`FontFamily`

The `fontFamily` name for subtitles, as described in [EBU-TT-D Subtitling Distribution Format](https://tech.ebu.ch/publications/tech3380).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MimeType`

The `mimeType` of the resource that's at the font download URL.

For information about font MIME types, see the [MPEG-DASH Profile for Transport of ISO BMFF Based DVB Services over IP Based Networks](https://dvb.org/wp-content/uploads/2021/06/A168r4_MPEG-DASH-Profile-for-Transport-of-ISO-BMFF-Based-DVB-Services_Draft-ts_103-285-v140_November_2021.pdf) document.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_/-]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL for downloading fonts for subtitles.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashBaseUrl

DashDvbMetricsReporting

All content copied from https://docs.aws.amazon.com/.
