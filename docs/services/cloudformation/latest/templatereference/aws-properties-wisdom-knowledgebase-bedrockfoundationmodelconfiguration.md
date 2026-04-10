This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase BedrockFoundationModelConfiguration

The configuration of the Bedrock foundation model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelArn" : String,
  "ParsingPrompt" : ParsingPrompt
}

```

### YAML

```yaml

  ModelArn: String
  ParsingPrompt:
    ParsingPrompt

```

## Properties

`ModelArn`

The model ARN of the Bedrock foundation model.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model\/anthropic.claude-3-haiku-20240307-v1:0$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsingPrompt`

The parsing prompt of the Bedrock foundation model configuration.

_Required_: No

_Type_: [ParsingPrompt](aws-properties-wisdom-knowledgebase-parsingprompt.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppIntegrationsConfiguration

ChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
