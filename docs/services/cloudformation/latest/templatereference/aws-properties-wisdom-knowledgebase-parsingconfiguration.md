This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase ParsingConfiguration

Settings for parsing document contents. By default, the service converts the contents of each document into text
before splitting it into chunks. To improve processing of PDF files with tables and images, you can configure the
data source to convert the pages of text into images and use a model to describe the contents of each page.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockFoundationModelConfiguration" : BedrockFoundationModelConfiguration,
  "ParsingStrategy" : String
}

```

### YAML

```yaml

  BedrockFoundationModelConfiguration:
    BedrockFoundationModelConfiguration
  ParsingStrategy: String

```

## Properties

`BedrockFoundationModelConfiguration`

Settings for a foundation model used to parse documents for a data source.

_Required_: No

_Type_: [BedrockFoundationModelConfiguration](aws-properties-wisdom-knowledgebase-bedrockfoundationmodelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsingStrategy`

The parsing strategy for the data source.

_Required_: Yes

_Type_: String

_Allowed values_: `BEDROCK_FOUNDATION_MODEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedSourceConfiguration

ParsingPrompt

All content copied from https://docs.aws.amazon.com/.
