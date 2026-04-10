This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent OrchestrationAIAgentConfiguration

The configuration for AI Agents of type `ORCHESTRATION`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectInstanceArn" : String,
  "Locale" : String,
  "OrchestrationAIGuardrailId" : String,
  "OrchestrationAIPromptId" : String,
  "ToolConfigurations" : [ ToolConfiguration, ... ]
}

```

### YAML

```yaml

  ConnectInstanceArn: String
  Locale: String
  OrchestrationAIGuardrailId: String
  OrchestrationAIPromptId: String
  ToolConfigurations:
    - ToolConfiguration

```

## Properties

`ConnectInstanceArn`

The Amazon Resource Name (ARN) of the Amazon Connect instance used by the Orchestration AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z-]+?:[a-z-]+?:[a-z0-9-]*?:([0-9]{12})?:[a-zA-Z0-9-:/]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Locale`

The locale setting for the Orchestration AI Agent.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationAIGuardrailId`

The AI Guardrail identifier used by the Orchestration AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationAIPromptId`

The AI Prompt identifier used by the Orchestration AI Agent.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(:[A-Z0-9_$]+){0,1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolConfigurations`

The tool configurations used by the Orchestration AI Agent.

_Required_: No

_Type_: Array of [ToolConfiguration](aws-properties-wisdom-aiagent-toolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoteTakingAIAgentConfiguration

OrCondition

All content copied from https://docs.aws.amazon.com/.
