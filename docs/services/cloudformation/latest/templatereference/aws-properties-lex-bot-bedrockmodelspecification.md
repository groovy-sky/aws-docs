This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot BedrockModelSpecification

Contains information about the Amazon Bedrock model used to interpret the prompt used in descriptive bot building.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockGuardrailConfiguration" : BedrockGuardrailConfiguration,
  "BedrockModelCustomPrompt" : String,
  "BedrockTraceStatus" : String,
  "ModelArn" : String
}

```

### YAML

```yaml

  BedrockGuardrailConfiguration:
    BedrockGuardrailConfiguration
  BedrockModelCustomPrompt: String
  BedrockTraceStatus: String
  ModelArn: String

```

## Properties

`BedrockGuardrailConfiguration`

Property description not available.

_Required_: No

_Type_: [BedrockGuardrailConfiguration](aws-properties-lex-bot-bedrockguardrailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BedrockModelCustomPrompt`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BedrockTraceStatus`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelArn`

The ARN of the foundation model used in descriptive bot building.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockKnowledgeStoreConfiguration

BKBExactResponseFields

All content copied from https://docs.aws.amazon.com/.
