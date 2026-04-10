This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent SelfServiceAIAgentConfiguration

The configuration of the self-service AI agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociationConfigurations" : [ AssociationConfiguration, ... ],
  "SelfServiceAIGuardrailId" : String,
  "SelfServiceAnswerGenerationAIPromptId" : String,
  "SelfServicePreProcessingAIPromptId" : String
}

```

### YAML

```yaml

  AssociationConfigurations:
    - AssociationConfiguration
  SelfServiceAIGuardrailId: String
  SelfServiceAnswerGenerationAIPromptId: String
  SelfServicePreProcessingAIPromptId: String

```

## Properties

`AssociationConfigurations`

The association configuration of the self-service AI agent.

_Required_: No

_Type_: Array of [AssociationConfiguration](aws-properties-wisdom-aiagent-associationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServiceAIGuardrailId`

The ID of the self-service AI guardrail.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServiceAnswerGenerationAIPromptId`

The ID of the self-service answer generation AI prompt.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServicePreProcessingAIPromptId`

The ID of the self-service preprocessing AI prompt.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrCondition

TagCondition

All content copied from https://docs.aws.amazon.com/.
