This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint DashPackage

Parameters for DASH packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdsOnDeliveryRestrictions" : String,
  "AdTriggers" : [ String, ... ],
  "Encryption" : DashEncryption,
  "IncludeIframeOnlyStream" : Boolean,
  "ManifestLayout" : String,
  "ManifestWindowSeconds" : Integer,
  "MinBufferTimeSeconds" : Integer,
  "MinUpdatePeriodSeconds" : Integer,
  "PeriodTriggers" : [ String, ... ],
  "Profile" : String,
  "SegmentDurationSeconds" : Integer,
  "SegmentTemplateFormat" : String,
  "StreamSelection" : StreamSelection,
  "SuggestedPresentationDelaySeconds" : Integer,
  "UtcTiming" : String,
  "UtcTimingUri" : String
}

```

### YAML

```yaml

  AdsOnDeliveryRestrictions: String
  AdTriggers:
    - String
  Encryption:
    DashEncryption
  IncludeIframeOnlyStream: Boolean
  ManifestLayout: String
  ManifestWindowSeconds: Integer
  MinBufferTimeSeconds: Integer
  MinUpdatePeriodSeconds: Integer
  PeriodTriggers:
    - String
  Profile: String
  SegmentDurationSeconds: Integer
  SegmentTemplateFormat: String
  StreamSelection:
    StreamSelection
  SuggestedPresentationDelaySeconds: Integer
  UtcTiming: String
  UtcTimingUri: String

```

## Properties

`AdsOnDeliveryRestrictions`

The flags on SCTE-35 segmentation descriptors that have to be present for AWS Elemental MediaPackage to insert ad markers in the output manifest. For information about SCTE-35 in AWS Elemental MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](../../../mediapackage/latest/ug/scte.md).

_Required_: No

_Type_: String

_Allowed values_: `NONE | RESTRICTED | UNRESTRICTED | BOTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdTriggers`

Specifies the SCTE-35 message types that AWS Elemental MediaPackage treats as ad markers in the output manifest.

Valid values:

- `BREAK`

- `DISTRIBUTOR_ADVERTISEMENT`

- `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`.

- `DISTRIBUTOR_PLACEMENT_OPPORTUNITY`.

- `PROVIDER_ADVERTISEMENT`.

- `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY`.

- `PROVIDER_PLACEMENT_OPPORTUNITY`.

- `SPLICE_INSERT`.

_Required_: No

_Type_: Array of String

_Allowed values_: `SPLICE_INSERT | BREAK | PROVIDER_ADVERTISEMENT | DISTRIBUTOR_ADVERTISEMENT | PROVIDER_PLACEMENT_OPPORTUNITY | DISTRIBUTOR_PLACEMENT_OPPORTUNITY | PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY | DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [DashEncryption](aws-properties-mediapackage-originendpoint-dashencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeIframeOnlyStream`

This applies only to stream sets with a single video track. When true, the stream set includes an
additional I-frame trick-play only stream, along with the other tracks. If false, this extra stream is not included.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestLayout`

Determines the position of some tags in the manifest.

Valid values:

- `FULL` \- Elements like `SegmentTemplate` and `ContentProtection` are included in each `Representation`.

- `COMPACT` \- Duplicate elements are combined and presented at the `AdaptationSet` level.

_Required_: No

_Type_: String

_Allowed values_: `FULL | COMPACT | DRM_TOP_LEVEL_COMPACT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

Time window (in seconds) contained in each manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBufferTimeSeconds`

Minimum amount of content (measured in seconds) that a player must keep available in the buffer.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinUpdatePeriodSeconds`

Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodTriggers`

Controls whether AWS Elemental MediaPackage produces single-period or multi-period DASH manifests. For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](../../../mediapackage/latest/ug/multi-period.md).

Valid values:

- `ADS` \- AWS Elemental MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.

- _No value_ \- AWS Elemental MediaPackage will produce single-period DASH manifests. This is the default setting.

_Required_: No

_Type_: Array of String

_Allowed values_: `ADS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profile`

The DASH profile for the output.

Valid values:

- `NONE` \- The output doesn't use a DASH profile.

- `HBBTV_1_5` \- The output is compliant with HbbTV v1.5.

- `DVB_DASH_2014` \- The output is compliant with DVB-DASH 2014.

_Required_: No

_Type_: String

_Allowed values_: `NONE | HBBTV_1_5 | HYBRIDCAST | DVB_DASH_2014`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentTemplateFormat`

Determines the type of variable used in the `media` URL of the `SegmentTemplate` tag in the manifest. Also specifies if segment timeline information is included in `SegmentTimeline` or `SegmentTemplate`.

Valid values:

- `NUMBER_WITH_TIMELINE` \- The `$Number$` variable is used in the `media` URL. The value of this variable is the sequential number of the segment. A full `SegmentTimeline` object is presented in each `SegmentTemplate`.

- `NUMBER_WITH_DURATION` \- The `$Number$` variable is used in the `media` URL and a `duration` attribute is added to
the segment template. The `SegmentTimeline` object is removed from the representation.

- `TIME_WITH_TIMELINE` \- The `$Time$` variable is used in the `media` URL. The value of this variable is the timestamp of when the segment starts. A full `SegmentTimeline` object is presented in each `SegmentTemplate`.

_Required_: No

_Type_: String

_Allowed values_: `NUMBER_WITH_TIMELINE | TIME_WITH_TIMELINE | NUMBER_WITH_DURATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Limitations for outputs from the endpoint, based on the video bitrate.

_Required_: No

_Type_: [StreamSelection](aws-properties-mediapackage-originendpoint-streamselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuggestedPresentationDelaySeconds`

Amount of time (in seconds) that the player should be from the live point at the end of the manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UtcTiming`

Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).

_Required_: No

_Type_: String

_Allowed values_: `HTTP-XSDATE | HTTP-ISO | HTTP-HEAD | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UtcTimingUri`

Specifies the value attribute of the UTC timing field when utcTiming is set to HTTP-ISO or HTTP-HEAD.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashEncryption

EncryptionContractConfiguration

All content copied from https://docs.aws.amazon.com/.
