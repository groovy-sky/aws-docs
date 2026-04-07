This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail

Creates a guardrail to detect and filter harmful content in your generative AI
application.

Amazon Bedrock Guardrails provides the following safeguards (also known as policies) to detect and filter harmful
content:

- **Content filters** \- Detect and filter harmful
text or image content in input prompts or model responses. Filtering is done based
on detection of certain predefined harmful content categories: Hate, Insults,
Sexual, Violence, Misconduct and Prompt Attack. You also can adjust the filter
strength for each of these categories.

- **Denied topics** \- Define a set of topics that
are undesirable in the context of your application. The filter will help block them if
detected in user queries or model responses.

- **Word filters** \- Configure filters to help
block undesirable words, phrases, and profanity (exact match). Such words can include offensive terms,
competitor names, etc.

- **Sensitive information filters** \- Configure filters to help block or
mask sensitive information, such as personally identifiable information (PII), or custom
regex in user inputs and model responses. Blocking or masking is done based on probabilistic detection of
sensitive information in standard formats in entities such as SSN number, Date of Birth, address, etc. This also allows configuring regular expression based
detection of patterns for identifiers.

- **Contextual grounding check** \- Help detect and filter
hallucinations in model responses based on grounding in a source and
relevance to the user query.

For more information, see [How Amazon Bedrock Guardrails works](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-how.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::Guardrail",
  "Properties" : {
      "AutomatedReasoningPolicyConfig" : AutomatedReasoningPolicyConfig,
      "BlockedInputMessaging" : String,
      "BlockedOutputsMessaging" : String,
      "ContentPolicyConfig" : ContentPolicyConfig,
      "ContextualGroundingPolicyConfig" : ContextualGroundingPolicyConfig,
      "CrossRegionConfig" : GuardrailCrossRegionConfig,
      "Description" : String,
      "KmsKeyArn" : String,
      "Name" : String,
      "SensitiveInformationPolicyConfig" : SensitiveInformationPolicyConfig,
      "Tags" : [ Tag, ... ],
      "TopicPolicyConfig" : TopicPolicyConfig,
      "WordPolicyConfig" : WordPolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::Guardrail
Properties:
  AutomatedReasoningPolicyConfig:
    AutomatedReasoningPolicyConfig
  BlockedInputMessaging: String
  BlockedOutputsMessaging: String
  ContentPolicyConfig:
    ContentPolicyConfig
  ContextualGroundingPolicyConfig:
    ContextualGroundingPolicyConfig
  CrossRegionConfig:
    GuardrailCrossRegionConfig
  Description: String
  KmsKeyArn: String
  Name: String
  SensitiveInformationPolicyConfig:
    SensitiveInformationPolicyConfig
  Tags:
    - Tag
  TopicPolicyConfig:
    TopicPolicyConfig
  WordPolicyConfig:
    WordPolicyConfig

```

## Properties

`AutomatedReasoningPolicyConfig`

Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.

_Required_: No

_Type_: [AutomatedReasoningPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-automatedreasoningpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedInputMessaging`

The message to return when the guardrail blocks a prompt.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedOutputsMessaging`

The message to return when the guardrail blocks a model response.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentPolicyConfig`

The content filter policies to configure for the guardrail.

_Required_: No

_Type_: [ContentPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-contentpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContextualGroundingPolicyConfig`

Property description not available.

_Required_: No

_Type_: [ContextualGroundingPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-contextualgroundingpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossRegionConfig`

The system-defined guardrail profile that you're using with your guardrail. Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed. Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.

For more information, see the [Amazon Bedrock User Guide](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region.html).

_Required_: No

_Type_: [GuardrailCrossRegionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-guardrailcrossregionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the guardrail.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the AWS KMS key that you use to encrypt the guardrail.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the guardrail.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-_]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitiveInformationPolicyConfig`

The sensitive information policy to configure for the guardrail.

_Required_: No

_Type_: [SensitiveInformationPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that you want to attach to the guardrail.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicPolicyConfig`

The topic policies to configure for the guardrail.

_Required_: No

_Type_: [TopicPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-topicpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordPolicyConfig`

The word policy you configure for the guardrail.

_Required_: No

_Type_: [WordPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-wordpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

Timestamp.

`FailureRecommendations`

Appears if the `status` of the guardrail is `FAILED`. A list of recommendations to carry out before retrying the request.

`GuardrailArn`

The ARN of the guardrail.

`GuardrailId`

The unique identifier of the guardrail.

`Status`

The status of the guardrail.

`StatusReasons`

Appears if the `status` is `FAILED`. A list of reasons for why the guardrail failed to be created, updated, versioned, or deleted.

`UpdatedAt`

Timestamp.

`Version`

The version of the guardrail that was created.
This value will always be `DRAFT`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VectorSearchRerankingConfiguration

AutomatedReasoningPolicyConfig
