This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint FilterConfiguration

Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClipStartTime" : String,
  "DrmSettings" : String,
  "End" : String,
  "ManifestFilter" : String,
  "Start" : String,
  "TimeDelaySeconds" : Integer
}

```

### YAML

```yaml

  ClipStartTime: String
  DrmSettings: String
  End: String
  ManifestFilter: String
  Start: String
  TimeDelaySeconds: Integer

```

## Properties

`ClipStartTime`

Optionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrmSettings`

Optionally specify one or more DRM settings for all of your manifest egress requests. When you include a DRM setting, note that you cannot use an identical DRM setting query parameter for this manifest's endpoint URL.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`End`

Optionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestFilter`

Optionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

Optionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeDelaySeconds`

Optionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1209600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionMethod

ForceEndpointErrorConfiguration

All content copied from https://docs.aws.amazon.com/.
