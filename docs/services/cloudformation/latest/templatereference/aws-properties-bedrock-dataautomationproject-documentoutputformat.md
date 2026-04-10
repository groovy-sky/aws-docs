This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject DocumentOutputFormat

A document output format.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalFileFormat" : DocumentOutputAdditionalFileFormat,
  "TextFormat" : DocumentOutputTextFormat
}

```

### YAML

```yaml

  AdditionalFileFormat:
    DocumentOutputAdditionalFileFormat
  TextFormat:
    DocumentOutputTextFormat

```

## Properties

`AdditionalFileFormat`

Output settings for additional file formats.

_Required_: Yes

_Type_: [DocumentOutputAdditionalFileFormat](aws-properties-bedrock-dataautomationproject-documentoutputadditionalfileformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextFormat`

An output text format.

_Required_: Yes

_Type_: [DocumentOutputTextFormat](aws-properties-bedrock-dataautomationproject-documentoutputtextformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentOutputAdditionalFileFormat

DocumentOutputTextFormat

All content copied from https://docs.aws.amazon.com/.
