This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject AudioStandardOutputConfiguration

Output settings for processing audio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Extraction" : AudioStandardExtraction,
  "GenerativeField" : AudioStandardGenerativeField
}

```

### YAML

```yaml

  Extraction:
    AudioStandardExtraction
  GenerativeField:
    AudioStandardGenerativeField

```

## Properties

`Extraction`

Settings for populating data fields that describe the audio.

_Required_: No

_Type_: [AudioStandardExtraction](aws-properties-bedrock-dataautomationproject-audiostandardextraction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerativeField`

Whether to generate descriptions of the data.

_Required_: No

_Type_: [AudioStandardGenerativeField](aws-properties-bedrock-dataautomationproject-audiostandardgenerativefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioStandardGenerativeField

BlueprintItem

All content copied from https://docs.aws.amazon.com/.
