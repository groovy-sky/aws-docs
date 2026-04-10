This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioNormalizationSettings

The settings for normalizing video.

The parent of this entity is AudioDescription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Algorithm" : String,
  "AlgorithmControl" : String,
  "TargetLkfs" : Number
}

```

### YAML

```yaml

  Algorithm: String
  AlgorithmControl: String
  TargetLkfs: Number

```

## Properties

`Algorithm`

The audio normalization algorithm to use. itu17701 conforms to the CALM Act specification. itu17702 conforms to
the EBU R-128 specification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlgorithmControl`

When set to correctAudio, the output audio is corrected using the chosen algorithm. If set to measureOnly, the
audio is measured but not adjusted.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetLkfs`

The Target LKFS(loudness) to adjust volume to. If no value is entered, a default value is used according to the
chosen algorithm. The CALM Act (1770-1) recommends a target of -24 LKFS. The EBU R-128 specification (1770-2)
recommends a target of -23 LKFS.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioLanguageSelection

AudioOnlyHlsSettings

All content copied from https://docs.aws.amazon.com/.
