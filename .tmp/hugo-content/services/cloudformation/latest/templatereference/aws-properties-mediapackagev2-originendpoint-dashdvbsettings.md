This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashDvbSettings

For endpoints that use the DVB-DASH profile only. The font download and error reporting information that you want MediaPackage to pass through to the manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorMetrics" : [ DashDvbMetricsReporting, ... ],
  "FontDownload" : DashDvbFontDownload
}

```

### YAML

```yaml

  ErrorMetrics:
    - DashDvbMetricsReporting
  FontDownload:
    DashDvbFontDownload

```

## Properties

`ErrorMetrics`

Playback device error reporting settings.

_Required_: No

_Type_: Array of [DashDvbMetricsReporting](aws-properties-mediapackagev2-originendpoint-dashdvbmetricsreporting.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontDownload`

Subtitle font settings.

_Required_: No

_Type_: [DashDvbFontDownload](aws-properties-mediapackagev2-originendpoint-dashdvbfontdownload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashDvbMetricsReporting

DashManifestConfiguration

All content copied from https://docs.aws.amazon.com/.
