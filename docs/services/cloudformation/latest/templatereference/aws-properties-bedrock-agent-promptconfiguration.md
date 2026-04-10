This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent PromptConfiguration

Contains configurations to override a prompt template in one part of an agent sequence. For more information, see [Advanced prompts](../../../bedrock/latest/userguide/advanced-prompts.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalModelRequestFields" : Json,
  "BasePromptTemplate" : String,
  "FoundationModel" : String,
  "InferenceConfiguration" : InferenceConfiguration,
  "ParserMode" : String,
  "PromptCreationMode" : String,
  "PromptState" : String,
  "PromptType" : String
}

```

### YAML

```yaml

  AdditionalModelRequestFields: Json
  BasePromptTemplate: String
  FoundationModel: String
  InferenceConfiguration:
    InferenceConfiguration
  ParserMode: String
  PromptCreationMode: String
  PromptState: String
  PromptType: String

```

## Properties

`AdditionalModelRequestFields`

If the Converse or ConverseStream operations support the model,
`additionalModelRequestFields` contains additional inference parameters,
beyond the base set of inference parameters in the `inferenceConfiguration`
field.

For more information, see [Inference request parameters\
and response fields for foundation models](../../../bedrock/latest/userguide/model-parameters.md).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasePromptTemplate`

Defines the prompt template with which to replace the default prompt template. You can use placeholder variables in the base prompt template to customize the prompt. For more information, see [Prompt template placeholder variables](../../../bedrock/latest/userguide/prompt-placeholders.md). For more information, see [Configure the prompt templates](../../../bedrock/latest/userguide/advanced-prompts-configure.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FoundationModel`

The agent's foundation model.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:(([0-9]{12}:custom-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}(([:][a-z0-9-]{1,63}){0,2})?/[a-z0-9]{12})|(:foundation-model/([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|([0-9]{12}:(inference-profile|application-inference-profile)/[a-zA-Z0-9-:.]+))|(([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|(([0-9a-zA-Z][_-]?)+)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceConfiguration`

Contains inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the `promptType`. For more information, see [Inference parameters for foundation models](../../../bedrock/latest/userguide/model-parameters.md).

_Required_: No

_Type_: [InferenceConfiguration](aws-properties-bedrock-agent-inferenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParserMode`

Specifies whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the `promptType`. If you set the field as `OVERRIDDEN`, the `overrideLambda` field in the [PromptOverrideConfiguration](../../../../reference/bedrock/latest/apireference/api-agent-promptoverrideconfiguration.md) must be specified with the ARN of a Lambda function.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | OVERRIDDEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptCreationMode`

Specifies whether to override the default prompt template for this `promptType`. Set this value to `OVERRIDDEN` to use the prompt that you provide in the `basePromptTemplate`. If you leave it as `DEFAULT`, the agent uses a default prompt template.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | OVERRIDDEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptState`

Specifies whether to allow the inline agent to carry out the step specified in the `promptType`. If you set this value to `DISABLED`, the agent skips that step. The default state for each `promptType` is as follows.

- `PRE_PROCESSING` – `ENABLED`

- `ORCHESTRATION` – `ENABLED`

- `KNOWLEDGE_BASE_RESPONSE_GENERATION` – `ENABLED`

- `POST_PROCESSING` – `DISABLED`

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptType`

The step in the agent sequence that this prompt configuration applies to.

_Required_: No

_Type_: String

_Allowed values_: `PRE_PROCESSING | ORCHESTRATION | POST_PROCESSING | ROUTING_CLASSIFIER | MEMORY_SUMMARIZATION | KNOWLEDGE_BASE_RESPONSE_GENERATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterDetail

PromptOverrideConfiguration

All content copied from https://docs.aws.amazon.com/.
