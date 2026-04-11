This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent AnswerRecommendationAIAgentConfiguration

The configuration for AI Agents of type `ANSWER_RECOMMENDATION`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerGenerationAIGuardrailId" : String,
  "AnswerGenerationAIPromptId" : String,
  "AssociationConfigurations" : [ AssociationConfiguration, ... ],
  "IntentLabelingGenerationAIPromptId" : String,
  "Locale" : String,
  "QueryReformulationAIPromptId" : String
}

```

### YAML

```yaml

  AnswerGenerationAIGuardrailId: String
  AnswerGenerationAIPromptId: String
  AssociationConfigurations:
    - AssociationConfiguration
  IntentLabelingGenerationAIPromptId: String
  Locale: String
  QueryReformulationAIPromptId: String

```

## Properties

`AnswerGenerationAIGuardrailId`

The ID of the answer generation AI guardrail.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnswerGenerationAIPromptId`

The AI Prompt identifier for the Answer Generation prompt used by the
`ANSWER_RECOMMENDATION` AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociationConfigurations`

The association configurations for overriding behavior on this AI Agent.

_Required_: No

_Type_: Array of [AssociationConfiguration](aws-properties-wisdom-aiagent-associationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntentLabelingGenerationAIPromptId`

The AI Prompt identifier for the Intent Labeling prompt used by the
`ANSWER_RECOMMENDATION` AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Locale`

The locale to which specifies the language and region settings that determine the
response language for [QueryAssistant](../../../../reference/connect/latest/apireference/api-amazon-q-connect-queryassistant.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryReformulationAIPromptId`

The AI Prompt identifier for the Query Reformulation prompt used by the
`ANSWER_RECOMMENDATION` AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AIAgentConfiguration

AssociationConfiguration

All content copied from https://docs.aws.amazon.com/.
