This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Scte27SourceSettings

Information about the SCTE-27 captions to extract from the input.

The parent of this entity is CaptionSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OcrLanguage" : String,
  "Pid" : Integer
}

```

### YAML

```yaml

  OcrLanguage: String
  Pid: Integer

```

## Properties

`OcrLanguage`

If you will configure a WebVTT caption description that references this caption selector, use this field to
provide the language to consider when translating the image-based source to text.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pid`

The PID field is used in conjunction with the captions selector languageCode field as follows: Specify PID and
Language: Extracts captions from that PID; the language is "informational." Specify PID and omit Language: Extracts
the specified PID. Omit PID and specify Language: Extracts the specified language, whichever PID that happens to be.
Omit PID and omit Language: Valid only if source is DVB-Sub that is being passed through; all languages are passed
through.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scte20SourceSettings

Scte35SpliceInsert

All content copied from https://docs.aws.amazon.com/.
