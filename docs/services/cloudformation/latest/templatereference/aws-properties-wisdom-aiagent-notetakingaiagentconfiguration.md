This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent NoteTakingAIAgentConfiguration

The configuration for AI Agents of type `NOTE_TAKING`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Locale" : String,
  "NoteTakingAIGuardrailId" : String,
  "NoteTakingAIPromptId" : String
}

```

### YAML

```yaml

  Locale: String
  NoteTakingAIGuardrailId: String
  NoteTakingAIPromptId: String

```

## Properties

`Locale`

The locale setting for language-specific case summarization generation (for example, en\_US, es\_ES).

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteTakingAIGuardrailId`

The AI Guardrail identifier used by the Note Taking AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteTakingAIPromptId`

The AI Prompt identifier used by the Note Taking AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManualSearchAIAgentConfiguration

OrchestrationAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
