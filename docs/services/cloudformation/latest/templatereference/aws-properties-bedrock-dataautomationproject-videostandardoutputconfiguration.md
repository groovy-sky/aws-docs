This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject VideoStandardOutputConfiguration

Output settings for processing video.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Extraction" : VideoStandardExtraction,
  "GenerativeField" : VideoStandardGenerativeField
}

```

### YAML

```yaml

  Extraction:
    VideoStandardExtraction
  GenerativeField:
    VideoStandardGenerativeField

```

## Properties

`Extraction`

Settings for populating data fields that describe the video.

_Required_: No

_Type_: [VideoStandardExtraction](aws-properties-bedrock-dataautomationproject-videostandardextraction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerativeField`

Whether to generate descriptions of the video.

_Required_: No

_Type_: [VideoStandardGenerativeField](aws-properties-bedrock-dataautomationproject-videostandardgenerativefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoStandardGenerativeField

AWS::Bedrock::DataSource

All content copied from https://docs.aws.amazon.com/.
