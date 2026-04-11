This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject AudioExtractionCategoryTypeConfiguration

Allows configuration of extractions for different types of data, such as transcript and content moderation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Transcript" : TranscriptConfiguration
}

```

### YAML

```yaml

  Transcript:
    TranscriptConfiguration

```

## Properties

`Transcript`

This element allows you to configure different extractions for your transcript data, such as speaker and channel labeling.

_Required_: No

_Type_: [TranscriptConfiguration](aws-properties-bedrock-dataautomationproject-transcriptconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioExtractionCategory

AudioLanguageConfiguration

All content copied from https://docs.aws.amazon.com/.
