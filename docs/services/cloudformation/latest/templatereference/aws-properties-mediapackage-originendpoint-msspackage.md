This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint MssPackage

Parameters for Microsoft Smooth Streaming packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : MssEncryption,
  "ManifestWindowSeconds" : Integer,
  "SegmentDurationSeconds" : Integer,
  "StreamSelection" : StreamSelection
}

```

### YAML

```yaml

  Encryption:
    MssEncryption
  ManifestWindowSeconds: Integer
  SegmentDurationSeconds: Integer
  StreamSelection:
    StreamSelection

```

## Properties

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [MssEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-mssencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

Time window (in seconds) contained in each manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Limitations for outputs from the endpoint, based on the video bitrate.

_Required_: No

_Type_: [StreamSelection](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-streamselection.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MssEncryption

SpekeKeyProvider
