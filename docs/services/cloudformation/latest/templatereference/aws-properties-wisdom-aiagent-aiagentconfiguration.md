This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent AIAgentConfiguration

A typed union that specifies the configuration based on the type of AI Agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerRecommendationAIAgentConfiguration" : AnswerRecommendationAIAgentConfiguration,
  "CaseSummarizationAIAgentConfiguration" : CaseSummarizationAIAgentConfiguration,
  "EmailGenerativeAnswerAIAgentConfiguration" : EmailGenerativeAnswerAIAgentConfiguration,
  "EmailOverviewAIAgentConfiguration" : EmailOverviewAIAgentConfiguration,
  "EmailResponseAIAgentConfiguration" : EmailResponseAIAgentConfiguration,
  "ManualSearchAIAgentConfiguration" : ManualSearchAIAgentConfiguration,
  "NoteTakingAIAgentConfiguration" : NoteTakingAIAgentConfiguration,
  "OrchestrationAIAgentConfiguration" : OrchestrationAIAgentConfiguration,
  "SelfServiceAIAgentConfiguration" : SelfServiceAIAgentConfiguration
}

```

### YAML

```yaml

  AnswerRecommendationAIAgentConfiguration:
    AnswerRecommendationAIAgentConfiguration
  CaseSummarizationAIAgentConfiguration:
    CaseSummarizationAIAgentConfiguration
  EmailGenerativeAnswerAIAgentConfiguration:
    EmailGenerativeAnswerAIAgentConfiguration
  EmailOverviewAIAgentConfiguration:
    EmailOverviewAIAgentConfiguration
  EmailResponseAIAgentConfiguration:
    EmailResponseAIAgentConfiguration
  ManualSearchAIAgentConfiguration:
    ManualSearchAIAgentConfiguration
  NoteTakingAIAgentConfiguration:
    NoteTakingAIAgentConfiguration
  OrchestrationAIAgentConfiguration:
    OrchestrationAIAgentConfiguration
  SelfServiceAIAgentConfiguration:
    SelfServiceAIAgentConfiguration

```

## Properties

`AnswerRecommendationAIAgentConfiguration`

The configuration for AI Agents of type `ANSWER_RECOMMENDATION`.

_Required_: No

_Type_: [AnswerRecommendationAIAgentConfiguration](aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaseSummarizationAIAgentConfiguration`

The configuration for AI Agents of type `CASE_SUMMARIZATION`.

_Required_: No

_Type_: [CaseSummarizationAIAgentConfiguration](aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailGenerativeAnswerAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_GENERATIVE_ANSWER`.

_Required_: No

_Type_: [EmailGenerativeAnswerAIAgentConfiguration](aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailOverviewAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_OVERVIEW`.

_Required_: No

_Type_: [EmailOverviewAIAgentConfiguration](aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailResponseAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_RESPONSE`.

_Required_: No

_Type_: [EmailResponseAIAgentConfiguration](aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManualSearchAIAgentConfiguration`

The configuration for AI Agents of type `MANUAL_SEARCH`.

_Required_: No

_Type_: [ManualSearchAIAgentConfiguration](aws-properties-wisdom-aiagent-manualsearchaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteTakingAIAgentConfiguration`

The configuration for AI Agents of type `NOTE_TAKING`.

_Required_: No

_Type_: [NoteTakingAIAgentConfiguration](aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationAIAgentConfiguration`

The configuration for AI Agents of type `ORCHESTRATION`.

_Required_: No

_Type_: [OrchestrationAIAgentConfiguration](aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServiceAIAgentConfiguration`

The self-service AI agent configuration.

_Required_: No

_Type_: [SelfServiceAIAgentConfiguration](aws-properties-wisdom-aiagent-selfserviceaiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Wisdom::AIAgent

AnswerRecommendationAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
