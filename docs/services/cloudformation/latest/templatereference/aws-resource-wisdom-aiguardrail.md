This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail

Creates an Amazon Q in Connect AI Guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIGuardrail",
  "Properties" : {
      "AssistantId" : String,
      "BlockedInputMessaging" : String,
      "BlockedOutputsMessaging" : String,
      "ContentPolicyConfig" : AIGuardrailContentPolicyConfig,
      "ContextualGroundingPolicyConfig" : AIGuardrailContextualGroundingPolicyConfig,
      "Description" : String,
      "Name" : String,
      "SensitiveInformationPolicyConfig" : AIGuardrailSensitiveInformationPolicyConfig,
      "Tags" : {Key: Value, ...},
      "TopicPolicyConfig" : AIGuardrailTopicPolicyConfig,
      "WordPolicyConfig" : AIGuardrailWordPolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIGuardrail
Properties:
  AssistantId: String
  BlockedInputMessaging: String
  BlockedOutputsMessaging: String
  ContentPolicyConfig:
    AIGuardrailContentPolicyConfig
  ContextualGroundingPolicyConfig:
    AIGuardrailContextualGroundingPolicyConfig
  Description: String
  Name: String
  SensitiveInformationPolicyConfig:
    AIGuardrailSensitiveInformationPolicyConfig
  Tags:
    Key: Value
  TopicPolicyConfig:
    AIGuardrailTopicPolicyConfig
  WordPolicyConfig:
    AIGuardrailWordPolicyConfig

```

## Properties

`AssistantId`

The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the
ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$|^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}){0,2}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BlockedInputMessaging`

The message to return when the AI Guardrail blocks a prompt.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedOutputsMessaging`

The message to return when the AI Guardrail blocks a model response.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentPolicyConfig`

Contains details about how to handle harmful content.

_Required_: No

_Type_: [AIGuardrailContentPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-aiguardrailcontentpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContextualGroundingPolicyConfig`

The policy configuration details for the AI Guardrail's contextual grounding policy.

_Required_: No

_Type_: [AIGuardrailContextualGroundingPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-aiguardrailcontextualgroundingpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the AI Guardrail.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AI Guardrail.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SensitiveInformationPolicyConfig`

Contains details about PII entities and regular expressions to configure for the AI Guardrail.

_Required_: No

_Type_: [AIGuardrailSensitiveInformationPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicPolicyConfig`

Contains details about topics that the AI Guardrail should identify and deny.

_Required_: No

_Type_: [AIGuardrailTopicPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-aiguardrailtopicpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordPolicyConfig`

Contains details about the word policy to configured for the AI Guardrail.

_Required_: No

_Type_: [AIGuardrailWordPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-aiguardrailwordpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AIGuardrailArn`

The Amazon Resource Name (ARN) of the AI Guardrail.

`AIGuardrailId`

The identifier of the Amazon Q in Connect AI Guardrail.

`AssistantArn`

The Amazon Resource Name (ARN) of the Amazon Q in Connect assistant.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::AIAgentVersion

AIGuardrailContentPolicyConfig
