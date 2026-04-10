This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot BedrockAgentIntentKnowledgeBaseConfiguration

The `BedrockAgentIntentKnowledgeBaseConfiguration` property type specifies Property description not available. for an [AWS::Lex::Bot](aws-resource-lex-bot.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockKnowledgeBaseArn" : String,
  "BedrockModelConfiguration" : BedrockModelSpecification
}

```

### YAML

```yaml

  BedrockKnowledgeBaseArn: String
  BedrockModelConfiguration:
    BedrockModelSpecification

```

## Properties

`BedrockKnowledgeBaseArn`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BedrockModelConfiguration`

Property description not available.

_Required_: Yes

_Type_: [BedrockModelSpecification](aws-properties-lex-bot-bedrockmodelspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockAgentIntentConfiguration

BedrockGuardrailConfiguration

All content copied from https://docs.aws.amazon.com/.
