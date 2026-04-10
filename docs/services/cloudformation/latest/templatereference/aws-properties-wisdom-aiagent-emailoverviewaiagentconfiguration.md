This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent EmailOverviewAIAgentConfiguration

The configuration for AI Agents of type `EMAIL_OVERVIEW`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailOverviewAIPromptId" : String,
  "Locale" : String
}

```

### YAML

```yaml

  EmailOverviewAIPromptId: String
  Locale: String

```

## Properties

`EmailOverviewAIPromptId`

The AI Prompt identifier for the Email Overview prompt used by the `EMAIL_OVERVIEW` AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Locale`

The locale to which specifies the language and region settings that determine the response language for [QueryAssistant](../../../../reference/connect/latest/apireference/api-amazon-q-connect-queryassistant.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmailGenerativeAnswerAIAgentConfiguration

EmailResponseAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
