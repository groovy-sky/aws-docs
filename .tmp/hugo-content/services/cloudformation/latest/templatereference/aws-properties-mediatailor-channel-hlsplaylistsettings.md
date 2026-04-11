This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel HlsPlaylistSettings

HLS playlist configuration parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkupType" : [ String, ... ],
  "ManifestWindowSeconds" : Number
}

```

### YAML

```yaml

  AdMarkupType:
    - String
  ManifestWindowSeconds: Number

```

## Properties

`AdMarkupType`

Determines the type of SCTE 35 tags to use in ad markup. Specify `DATERANGE` to use `DATERANGE` tags (for live or VOD content). Specify `SCTE35_ENHANCED` to use `EXT-X-CUE-OUT` and `EXT-X-CUE-IN` tags (for VOD content only).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

The total duration (in seconds) of each manifest. Minimum value: `30` seconds. Maximum value: `3600` seconds.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashPlaylistSettings

LogConfigurationForChannel

All content copied from https://docs.aws.amazon.com/.
