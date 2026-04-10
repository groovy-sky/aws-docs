This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject AudioLanguageConfiguration

This allows you to set the input and output language of your audio. The input language can be set to any of the languages supported by Bedrock Data Automation.
The output can either be set to english or whatever the dominant language is of the audio, determined by the language spoken for the most seconds.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GenerativeOutputLanguage" : String,
  "IdentifyMultipleLanguages" : Boolean,
  "InputLanguages" : [ String, ... ]
}

```

### YAML

```yaml

  GenerativeOutputLanguage: String
  IdentifyMultipleLanguages: Boolean
  InputLanguages:
    - String

```

## Properties

`GenerativeOutputLanguage`

The output language of your processing results. This can either be set to `EN` (English) or `DEFAULT` which will output
the results in the dominant language of the audio. The dominant language is determined as the language in the audio, spoken the longest in the
input audio.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | EN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentifyMultipleLanguages`

The toggle determining if you want to detect multiple languages from your audio.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLanguages`

The input language of your audio. This can be set to any of the currently supported languages via the language codes.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioExtractionCategoryTypeConfiguration

AudioOverrideConfiguration

All content copied from https://docs.aws.amazon.com/.
