This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel TemporalFilterSettings

Settings for the temporal filter to apply to the video.

The parents of this entity are H264FilterSettings, H265FilterSettings, and Mpeg2FilterSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PostFilterSharpening" : String,
  "Strength" : String
}

```

### YAML

```yaml

  PostFilterSharpening: String
  Strength: String

```

## Properties

`PostFilterSharpening`

If you enable this filter, the results are the following:
\- If the source content is noisy (it contains excessive digital artifacts), the filter cleans up the source.
\- If the source content is already clean, the filter tends to decrease the bitrate, especially when the rate control mode is QVBR.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Strength`

Choose a filter strength. We recommend a strength of 1 or 2. A higher strength might take out good information, resulting in an image that is overly soft.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TeletextSourceSettings

ThumbnailConfiguration

All content copied from https://docs.aws.amazon.com/.
