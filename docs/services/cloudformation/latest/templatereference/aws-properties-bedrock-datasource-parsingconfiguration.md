This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource ParsingConfiguration

Settings for parsing document contents. If you exclude this field, the default parser converts the contents of each
document into text before splitting it into chunks. Specify the parsing strategy to use in the `parsingStrategy` field and include the relevant configuration, or omit it to use the Amazon Bedrock default parser. For more information, see [Parsing options for your data source](../../../bedrock/latest/userguide/kb-advanced-parsing.md).

###### Note

If you specify `BEDROCK_DATA_AUTOMATION` or `BEDROCK_FOUNDATION_MODEL` and it fails to parse a file, the Amazon Bedrock default parser will be used instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockDataAutomationConfiguration" : BedrockDataAutomationConfiguration,
  "BedrockFoundationModelConfiguration" : BedrockFoundationModelConfiguration,
  "ParsingStrategy" : String
}

```

### YAML

```yaml

  BedrockDataAutomationConfiguration:
    BedrockDataAutomationConfiguration
  BedrockFoundationModelConfiguration:
    BedrockFoundationModelConfiguration
  ParsingStrategy: String

```

## Properties

`BedrockDataAutomationConfiguration`

If you specify `BEDROCK_DATA_AUTOMATION` as the parsing strategy for ingesting your data source, use this object to modify configurations for using the Amazon Bedrock Data Automation parser.

_Required_: No

_Type_: [BedrockDataAutomationConfiguration](aws-properties-bedrock-datasource-bedrockdataautomationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BedrockFoundationModelConfiguration`

If you specify `BEDROCK_FOUNDATION_MODEL` as the parsing strategy for ingesting your data source, use this object to modify configurations for using a foundation model to parse documents.

_Required_: No

_Type_: [BedrockFoundationModelConfiguration](aws-properties-bedrock-datasource-bedrockfoundationmodelconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParsingStrategy`

The parsing strategy for the data source.

_Required_: Yes

_Type_: String

_Allowed values_: `BEDROCK_FOUNDATION_MODEL | BEDROCK_DATA_AUTOMATION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntermediateStorage

ParsingPrompt

All content copied from https://docs.aws.amazon.com/.
