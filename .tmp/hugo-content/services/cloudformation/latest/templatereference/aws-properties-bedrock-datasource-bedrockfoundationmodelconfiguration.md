This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource BedrockFoundationModelConfiguration

Settings for a foundation model used to parse documents for a data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelArn" : String,
  "ParsingModality" : String,
  "ParsingPrompt" : ParsingPrompt
}

```

### YAML

```yaml

  ModelArn: String
  ParsingModality: String
  ParsingPrompt:
    ParsingPrompt

```

## Properties

`ModelArn`

The ARN of the foundation model to use for parsing.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-cn|-us-gov|-eusc|-iso(-[b-f])?)?:(bedrock):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParsingModality`

Specifies whether to enable parsing of multimodal data, including both text and/or images.

_Required_: No

_Type_: String

_Allowed values_: `MULTIMODAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParsingPrompt`

Instructions for interpreting the contents of a document.

_Required_: No

_Type_: [ParsingPrompt](aws-properties-bedrock-datasource-parsingprompt.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockDataAutomationConfiguration

BedrockFoundationModelContextEnrichmentConfiguration

All content copied from https://docs.aws.amazon.com/.
