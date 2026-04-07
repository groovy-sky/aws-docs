This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail TopicPolicyConfig

Contains details about topics that the guardrail should identify and deny.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicsConfig" : [ TopicConfig, ... ],
  "TopicsTierConfig" : TopicsTierConfig
}

```

### YAML

```yaml

  TopicsConfig:
    - TopicConfig
  TopicsTierConfig:
    TopicsTierConfig

```

## Properties

`TopicsConfig`

A list of policies related to topics that the guardrail should deny.

_Required_: Yes

_Type_: Array of [TopicConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-topicconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicsTierConfig`

The tier that your guardrail uses for denied topic filters.

_Required_: No

_Type_: [TopicsTierConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-guardrail-topicstierconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TopicConfig

TopicsTierConfig
