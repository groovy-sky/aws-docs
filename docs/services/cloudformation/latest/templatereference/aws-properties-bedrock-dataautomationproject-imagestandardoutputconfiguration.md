This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject ImageStandardOutputConfiguration

Output settings for processing images.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Extraction" : ImageStandardExtraction,
  "GenerativeField" : ImageStandardGenerativeField
}

```

### YAML

```yaml

  Extraction:
    ImageStandardExtraction
  GenerativeField:
    ImageStandardGenerativeField

```

## Properties

`Extraction`

Settings for populating data fields that describe the image.

_Required_: No

_Type_: [ImageStandardExtraction](aws-properties-bedrock-dataautomationproject-imagestandardextraction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerativeField`

Whether to generate descriptions of the data.

_Required_: No

_Type_: [ImageStandardGenerativeField](aws-properties-bedrock-dataautomationproject-imagestandardgenerativefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageStandardGenerativeField

ModalityProcessingConfiguration

All content copied from https://docs.aws.amazon.com/.
