This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject DocumentStandardOutputConfiguration

Output settings for processing documents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Extraction" : DocumentStandardExtraction,
  "GenerativeField" : DocumentStandardGenerativeField,
  "OutputFormat" : DocumentOutputFormat
}

```

### YAML

```yaml

  Extraction:
    DocumentStandardExtraction
  GenerativeField:
    DocumentStandardGenerativeField
  OutputFormat:
    DocumentOutputFormat

```

## Properties

`Extraction`

Settings for populating data fields that describe the document.

_Required_: No

_Type_: [DocumentStandardExtraction](aws-properties-bedrock-dataautomationproject-documentstandardextraction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerativeField`

Whether to generate descriptions.

_Required_: No

_Type_: [DocumentStandardGenerativeField](aws-properties-bedrock-dataautomationproject-documentstandardgenerativefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFormat`

The output format to generate.

_Required_: No

_Type_: [DocumentOutputFormat](aws-properties-bedrock-dataautomationproject-documentoutputformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentStandardGenerativeField

ImageBoundingBox

All content copied from https://docs.aws.amazon.com/.
