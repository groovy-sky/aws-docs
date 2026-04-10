This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::RecordingConfiguration RenditionConfiguration

The RenditionConfiguration property type describes which renditions should be recorded for a stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Renditions" : [ String, ... ],
  "RenditionSelection" : String
}

```

### YAML

```yaml

  Renditions:
    - String
  RenditionSelection: String

```

## Properties

`Renditions`

A list of which renditions are recorded for a stream, if `renditionSelection`
is `CUSTOM`; otherwise, this field is irrelevant. The selected renditions are
recorded if they are available during the stream. If a selected rendition is unavailable, the
best available rendition is recorded. For details on the resolution dimensions of each
rendition, see [Auto-Record to Amazon S3](../../../ivs/latest/lowlatencyuserguide/record-to-s3.md).

_Required_: No

_Type_: Array of String

_Allowed values_: `FULL_HD | HD | SD | LOWEST_RESOLUTION`

_Minimum_: `0`

_Maximum_: `4`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RenditionSelection`

The set of renditions are recorded for a stream. For `BASIC`
channels, the `CUSTOM` value has no effect. If `CUSTOM` is specified, a
set of renditions can be specified in the `renditions` field. Default:
`ALL`.

_Required_: No

_Type_: String

_Allowed values_: `ALL | NONE | CUSTOM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationConfiguration

S3DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
