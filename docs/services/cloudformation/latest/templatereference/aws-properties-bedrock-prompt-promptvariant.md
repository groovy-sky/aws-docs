This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt PromptVariant

Contains details about a variant of the prompt.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalModelRequestFields" : Json,
  "GenAiResource" : PromptGenAiResource,
  "InferenceConfiguration" : PromptInferenceConfiguration,
  "Metadata" : [ PromptMetadataEntry, ... ],
  "ModelId" : String,
  "Name" : String,
  "TemplateConfiguration" : PromptTemplateConfiguration,
  "TemplateType" : String
}

```

### YAML

```yaml

  AdditionalModelRequestFields: Json
  GenAiResource:
    PromptGenAiResource
  InferenceConfiguration:
    PromptInferenceConfiguration
  Metadata:
    - PromptMetadataEntry
  ModelId: String
  Name: String
  TemplateConfiguration:
    PromptTemplateConfiguration
  TemplateType: String

```

## Properties

`AdditionalModelRequestFields`

Contains model-specific inference configurations that aren't in the `inferenceConfiguration` field. To see model-specific inference parameters, see [Inference request parameters and response fields for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenAiResource`

Specifies a generative AI resource with which to use the prompt.

_Required_: No

_Type_: [PromptGenAiResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-prompt-promptgenairesource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceConfiguration`

Contains inference configurations for the prompt variant.

_Required_: No

_Type_: [PromptInferenceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-prompt-promptinferenceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

An array of objects, each containing a key-value pair that defines a metadata tag and value to attach to a prompt variant.

_Required_: No

_Type_: Array of [PromptMetadataEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-prompt-promptmetadataentry.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) with which to run inference on the prompt.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws(-[^:]{1,12})?:(bedrock|sagemaker):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the prompt variant.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

Contains configurations for the prompt template.

_Required_: Yes

_Type_: [PromptTemplateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-prompt-prompttemplateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateType`

The type of prompt template to use.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT | CHAT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PromptTemplateConfiguration

SpecificToolChoice
