This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration DashPackage

Parameters for a packaging configuration that uses Dynamic Adaptive Streaming over HTTP (DASH) packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DashManifests" : [ DashManifest, ... ],
  "Encryption" : DashEncryption,
  "IncludeEncoderConfigurationInSegments" : Boolean,
  "IncludeIframeOnlyStream" : Boolean,
  "PeriodTriggers" : [ String, ... ],
  "SegmentDurationSeconds" : Integer,
  "SegmentTemplateFormat" : String
}

```

### YAML

```yaml

  DashManifests:
    - DashManifest
  Encryption:
    DashEncryption
  IncludeEncoderConfigurationInSegments: Boolean
  IncludeIframeOnlyStream: Boolean
  PeriodTriggers:
    - String
  SegmentDurationSeconds: Integer
  SegmentTemplateFormat: String

```

## Properties

`DashManifests`

A list of DASH manifest configurations that are available from this endpoint.

_Required_: Yes

_Type_: Array of [DashManifest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [DashEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-packagingconfiguration-dashencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeEncoderConfigurationInSegments`

When includeEncoderConfigurationInSegments is set to true, AWS Elemental MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment. This lets you use different SPS/PPS/VPS settings for your assets during content playback.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeIframeOnlyStream`

This applies only to stream sets with a single video track. When true, the stream set includes an
additional I-frame trick-play only stream, along with the other tracks. If false, this extra stream is not included.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodTriggers`

Controls whether AWS Elemental MediaPackage produces single-period or multi-period DASH manifests. For more information about periods, see [Multi-period DASH in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/multi-period.html).

Valid values:

- `ADS` \- AWS Elemental MediaPackage will produce multi-period DASH manifests. Periods are created based on the SCTE-35 ad markers present in the input manifest.

- _No value_ \- AWS Elemental MediaPackage will produce single-period DASH manifests. This is the default setting.

_Required_: No

_Type_: Array of String

_Allowed values_: `ADS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source segment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentTemplateFormat`

Determines the type of SegmentTemplate included in the Media Presentation Description (MPD). When set to `NUMBER_WITH_TIMELINE`, a full timeline is
presented in each SegmentTemplate, with $Number$ media URLs. When set to `TIME_WITH_TIMELINE`, a full timeline is presented in each
SegmentTemplate, with $Time$ media URLs. When set to `NUMBER_WITH_DURATION`, only a duration is included in each
SegmentTemplate, with $Number$ media URLs.

_Required_: No

_Type_: String

_Allowed values_: `NUMBER_WITH_TIMELINE | TIME_WITH_TIMELINE | NUMBER_WITH_DURATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DashManifest

EncryptionContractConfiguration
