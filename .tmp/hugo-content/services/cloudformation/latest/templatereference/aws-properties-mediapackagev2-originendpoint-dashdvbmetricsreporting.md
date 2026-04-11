This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashDvbMetricsReporting

For use with DVB-DASH profiles only. The settings for error reporting from the playback device that you want AWS Elemental MediaPackage to pass through to the manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Probability" : Integer,
  "ReportingUrl" : String
}

```

### YAML

```yaml

  Probability: Integer
  ReportingUrl: String

```

## Properties

`Probability`

The number of playback devices per 1000 that will send error reports to the reporting URL. This represents the probability that a playback device will be a reporting player for this session.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportingUrl`

The URL where playback devices send error reports.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashDvbFontDownload

DashDvbSettings

All content copied from https://docs.aws.amazon.com/.
