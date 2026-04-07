This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashManifestConfiguration

The DASH manifest configuration associated with the origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseUrls" : [ DashBaseUrl, ... ],
  "Compactness" : String,
  "DrmSignaling" : String,
  "DvbSettings" : DashDvbSettings,
  "FilterConfiguration" : FilterConfiguration,
  "ManifestName" : String,
  "ManifestWindowSeconds" : Integer,
  "MinBufferTimeSeconds" : Integer,
  "MinUpdatePeriodSeconds" : Integer,
  "PeriodTriggers" : [ String, ... ],
  "Profiles" : [ String, ... ],
  "ProgramInformation" : DashProgramInformation,
  "ScteDash" : ScteDash,
  "SegmentTemplateFormat" : String,
  "SubtitleConfiguration" : DashSubtitleConfiguration,
  "SuggestedPresentationDelaySeconds" : Integer,
  "UtcTiming" : DashUtcTiming
}

```

### YAML

```yaml

  BaseUrls:
    - DashBaseUrl
  Compactness: String
  DrmSignaling: String
  DvbSettings:
    DashDvbSettings
  FilterConfiguration:
    FilterConfiguration
  ManifestName: String
  ManifestWindowSeconds: Integer
  MinBufferTimeSeconds: Integer
  MinUpdatePeriodSeconds: Integer
  PeriodTriggers:
    - String
  Profiles:
    - String
  ProgramInformation:
    DashProgramInformation
  ScteDash:
    ScteDash
  SegmentTemplateFormat: String
  SubtitleConfiguration:
    DashSubtitleConfiguration
  SuggestedPresentationDelaySeconds: Integer
  UtcTiming:
    DashUtcTiming

```

## Properties

`BaseUrls`

The base URLs to use for retrieving segments.

_Required_: No

_Type_: Array of [DashBaseUrl](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compactness`

The layout of the DASH manifest that MediaPackage produces. `STANDARD` indicates a default manifest, which is compacted. `NONE` indicates a full manifest.

For information about compactness, see [DASH manifest compactness](https://docs.aws.amazon.com/mediapackage/latest/userguide/compacted.html) in the _AWS Elemental MediaPackage v2 User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrmSignaling`

Determines how the DASH manifest signals the DRM content.

_Required_: No

_Type_: String

_Allowed values_: `INDIVIDUAL | REFERENCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbSettings`

For endpoints that use the DVB-DASH profile only. The font download and error reporting information that you want MediaPackage to pass through to the manifest.

_Required_: No

_Type_: [DashDvbSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-dashdvbsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.

_Required_: No

_Type_: [FilterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

A short string that's appended to the endpoint URL. The child manifest name creates a unique path to this endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

The total duration (in seconds) of the manifest's content.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBufferTimeSeconds`

Minimum amount of content (in seconds) that a player must keep available in the buffer.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinUpdatePeriodSeconds`

Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodTriggers`

A list of triggers that controls when AWS Elemental MediaPackage separates the MPEG-DASH manifest into multiple periods. Type `ADS` to indicate that AWS Elemental MediaPackage must create periods in the output manifest that correspond to SCTE-35 ad markers in the input source. Leave this value empty to indicate that the manifest is contained all in one period.
For more information about periods in the DASH manifest, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/userguide/multi-period.html).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profiles`

The profile that the output is compliant with.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramInformation`

Details about the content that you want MediaPackage to pass through in the manifest to the playback device.

_Required_: No

_Type_: [DashProgramInformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-dashprograminformation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScteDash`

The SCTE configuration.

_Required_: No

_Type_: [ScteDash](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-sctedash.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentTemplateFormat`

Determines the type of variable used in the `media` URL of the `SegmentTemplate` tag in the manifest. Also specifies if segment timeline information is included in `SegmentTimeline` or `SegmentTemplate`.

Value description:

- `NUMBER_WITH_TIMELINE` \- The `$Number$` variable is used in the `media` URL. The value of this variable is the sequential number of the segment. A full `SegmentTimeline` object is presented in each `SegmentTemplate`.

_Required_: No

_Type_: String

_Allowed values_: `NUMBER_WITH_TIMELINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubtitleConfiguration`

The configuration for DASH subtitles.

_Required_: No

_Type_: [DashSubtitleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-dashsubtitleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuggestedPresentationDelaySeconds`

The amount of time (in seconds) that the player should be from the end of the manifest.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UtcTiming`

Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).

_Required_: No

_Type_: [DashUtcTiming](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-dashutctiming.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DashDvbSettings

DashProgramInformation
