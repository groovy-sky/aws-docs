This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent EmailResponseAIAgentConfiguration

The configuration for AI Agents of type `EMAIL_RESPONSE`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociationConfigurations" : [ AssociationConfiguration, ... ],
  "EmailQueryReformulationAIPromptId" : String,
  "EmailResponseAIPromptId" : String,
  "Locale" : String
}

```

### YAML

```yaml

  AssociationConfigurations:
    - AssociationConfiguration
  EmailQueryReformulationAIPromptId: String
  EmailResponseAIPromptId: String
  Locale: String

```

## Properties

`AssociationConfigurations`

The association configurations for overriding behavior on this AI Agent.

_Required_: No

_Type_: Array of [AssociationConfiguration](aws-properties-wisdom-aiagent-associationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailQueryReformulationAIPromptId`

The AI Prompt identifier for the Email Query Reformulation prompt used by the `EMAIL_RESPONSE` AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailResponseAIPromptId`

The AI Prompt identifier for the Email Response prompt used by the `EMAIL_RESPONSE` AI Agent.

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

EmailOverviewAIAgentConfiguration

KnowledgeBaseAssociationConfigurationData

All content copied from https://docs.aws.amazon.com/.
