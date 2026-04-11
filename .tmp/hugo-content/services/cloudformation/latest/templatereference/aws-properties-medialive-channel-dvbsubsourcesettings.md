This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel DvbSubSourceSettings

Information about the DVB Sub captions to extract from the input.

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

When using DVB-Sub with burn-in or SMPTE-TT, use this PID for the source content. It is unused for DVB-Sub
passthrough. All DVB-Sub content is passed through, regardless of selectors.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DvbSubDestinationSettings

DvbTdtSettings

All content copied from https://docs.aws.amazon.com/.
