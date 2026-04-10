This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent CaseSummarizationAIAgentConfiguration

The configuration for AI Agents of type `CASE_SUMMARIZATION`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseSummarizationAIGuardrailId" : String,
  "CaseSummarizationAIPromptId" : String,
  "Locale" : String
}

```

### YAML

```yaml

  CaseSummarizationAIGuardrailId: String
  CaseSummarizationAIPromptId: String
  Locale: String

```

## Properties

`CaseSummarizationAIGuardrailId`

The AI Guardrail identifier used by the Case Summarization AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaseSummarizationAIPromptId`

The AI Prompt identifier used by the Case Summarization AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Locale`

The locale setting for the Case Summarization AI Agent.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociationConfigurationData

EmailGenerativeAnswerAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
