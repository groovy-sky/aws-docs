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

_Type_: [AnswerRecommendationAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-answerrecommendationaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaseSummarizationAIAgentConfiguration`

The configuration for AI Agents of type `CASE_SUMMARIZATION`.

_Required_: No

_Type_: [CaseSummarizationAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailGenerativeAnswerAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_GENERATIVE_ANSWER`.

_Required_: No

_Type_: [EmailGenerativeAnswerAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-emailgenerativeansweraiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailOverviewAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_OVERVIEW`.

_Required_: No

_Type_: [EmailOverviewAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailResponseAIAgentConfiguration`

The configuration for AI Agents of type `EMAIL_RESPONSE`.

_Required_: No

_Type_: [EmailResponseAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-emailresponseaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManualSearchAIAgentConfiguration`

The configuration for AI Agents of type `MANUAL_SEARCH`.

_Required_: No

_Type_: [ManualSearchAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-manualsearchaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteTakingAIAgentConfiguration`

The configuration for AI Agents of type `NOTE_TAKING`.

_Required_: No

_Type_: [NoteTakingAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationAIAgentConfiguration`

The configuration for AI Agents of type `ORCHESTRATION`.

_Required_: No

_Type_: [OrchestrationAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServiceAIAgentConfiguration`

The self-service AI agent configuration.

_Required_: No

_Type_: [SelfServiceAIAgentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiagent-selfserviceaiagentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::AIAgent

AnswerRecommendationAIAgentConfiguration
