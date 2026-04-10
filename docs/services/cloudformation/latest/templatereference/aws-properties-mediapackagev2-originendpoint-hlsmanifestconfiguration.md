This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint HlsManifestConfiguration

The HLS manifest configuration associated with the origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildManifestName" : String,
  "FilterConfiguration" : FilterConfiguration,
  "ManifestName" : String,
  "ManifestWindowSeconds" : Integer,
  "ProgramDateTimeIntervalSeconds" : Integer,
  "ScteHls" : ScteHls,
  "StartTag" : StartTag,
  "Url" : String,
  "UrlEncodeChildManifest" : Boolean
}

```

### YAML

```yaml

  ChildManifestName: String
  FilterConfiguration:
    FilterConfiguration
  ManifestName: String
  ManifestWindowSeconds: Integer
  ProgramDateTimeIntervalSeconds: Integer
  ScteHls:
    ScteHls
  StartTag:
    StartTag
  Url: String
  UrlEncodeChildManifest: Boolean

```

## Properties

`ChildManifestName`

The name of the child manifest associated with the HLS manifest configuration.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.

_Required_: No

_Type_: [FilterConfiguration](aws-properties-mediapackagev2-originendpoint-filterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

The name of the manifest associated with the HLS manifest configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

The duration of the manifest window, in seconds, for the HLS manifest configuration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramDateTimeIntervalSeconds`

The `EXT-X-PROGRAM-DATE-TIME` interval, in seconds, associated with the HLS manifest configuration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScteHls`

THE SCTE-35 HLS configuration associated with the HLS manifest configuration.

_Required_: No

_Type_: [ScteHls](aws-properties-mediapackagev2-originendpoint-sctehls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTag`

To insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset. When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.

_Required_: No

_Type_: [StartTag](aws-properties-mediapackagev2-originendpoint-starttag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL of the HLS manifest configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlEncodeChildManifest`

When enabled, MediaPackage URL-encodes the query string for API requests for HLS child manifests to comply with AWS Signature Version 4 (SigV4) signature signing protocol.
For more information, see [AWS Signature Version 4 for API requests](../../../iam/latest/userguide/reference-sigv.md) in _AWS Identity and Access Management User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForceEndpointErrorConfiguration

LowLatencyHlsManifestConfiguration

All content copied from https://docs.aws.amazon.com/.
