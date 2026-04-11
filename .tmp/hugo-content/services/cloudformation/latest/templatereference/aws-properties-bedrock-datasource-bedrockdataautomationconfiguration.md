This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource BedrockDataAutomationConfiguration

Contains configurations for using Amazon Bedrock Data Automation as the parser for ingesting your data sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParsingModality" : String
}

```

### YAML

```yaml

  ParsingModality: String

```

## Properties

`ParsingModality`

Specifies whether to enable parsing of multimodal data, including both text and/or images.

_Required_: No

_Type_: String

_Allowed values_: `MULTIMODAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::DataSource

BedrockFoundationModelConfiguration

All content copied from https://docs.aws.amazon.com/.
