This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIPrompt

Creates an Amazon Q in Connect AI Prompt.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIPrompt",
  "Properties" : {
      "ApiFormat" : String,
      "AssistantId" : String,
      "Description" : String,
      "ModelId" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "TemplateConfiguration" : AIPromptTemplateConfiguration,
      "TemplateType" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIPrompt
Properties:
  ApiFormat: String
  AssistantId: String
  Description: String
  ModelId: String
  Name: String
  Tags:
    Key: Value
  TemplateConfiguration:
    AIPromptTemplateConfiguration
  TemplateType: String
  Type: String

```

## Properties

`ApiFormat`

The API format used for this AI Prompt.

_Required_: Yes

_Type_: String

_Allowed values_: `ANTHROPIC_CLAUDE_MESSAGES | ANTHROPIC_CLAUDE_TEXT_COMPLETIONS | MESSAGES | TEXT_COMPLETIONS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssistantId`

The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the
ARN.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$|^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}){0,2}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the AI Prompt.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The identifier of the model used for this AI Prompt. The following model Ids are supported:

- `anthropic.claude-3-haiku--v1:0`

- `apac.amazon.nova-lite-v1:0`

- `apac.amazon.nova-micro-v1:0`

- `apac.amazon.nova-pro-v1:0`

- `apac.anthropic.claude-3-5-sonnet--v2:0`

- `apac.anthropic.claude-3-haiku-20240307-v1:0`

- `eu.amazon.nova-lite-v1:0`

- `eu.amazon.nova-micro-v1:0`

- `eu.amazon.nova-pro-v1:0`

- `eu.anthropic.claude-3-7-sonnet-20250219-v1:0`

- `eu.anthropic.claude-3-haiku-20240307-v1:0`

- `us.amazon.nova-lite-v1:0`

- `us.amazon.nova-micro-v1:0`

- `us.amazon.nova-pro-v1:0`

- `us.anthropic.claude-3-5-haiku-20241022-v1:0`

- `us.anthropic.claude-3-7-sonnet-20250219-v1:0`

- `us.anthropic.claude-3-haiku-20240307-v1:0`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AI Prompt

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateConfiguration`

The configuration of the prompt template for this AI Prompt.

_Required_: Yes

_Type_: [AIPromptTemplateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiprompt-aiprompttemplateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateType`

The type of the prompt template for this AI Prompt.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of this AI Prompt.

_Required_: Yes

_Type_: String

_Allowed values_: `ANSWER_GENERATION | INTENT_LABELING_GENERATION | QUERY_REFORMULATION | SELF_SERVICE_PRE_PROCESSING | SELF_SERVICE_ANSWER_GENERATION | EMAIL_RESPONSE | EMAIL_OVERVIEW | EMAIL_GENERATIVE_ANSWER | EMAIL_QUERY_REFORMULATION | ORCHESTRATION | NOTE_TAKING | CASE_SUMMARIZATION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`AIPromptArn`

The Amazon Resource Name (ARN) of the AI Prompt.

`AIPromptId`

The identifier of the Amazon Q in Connect AI prompt.

`AssistantArn`

The Amazon Resource Name (ARN) of the Amazon Q in Connect assistant.

`ModifiedTimeSeconds`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::AIGuardrailVersion

AIPromptTemplateConfiguration
