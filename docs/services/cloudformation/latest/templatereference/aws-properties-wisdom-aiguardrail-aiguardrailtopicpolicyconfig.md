This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail AIGuardrailTopicPolicyConfig

Topic policy configuration for a guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicsConfig" : [ GuardrailTopicConfig, ... ]
}

```

### YAML

```yaml

  TopicsConfig:
    - GuardrailTopicConfig

```

## Properties

`TopicsConfig`

List of topic configs in topic policy.

_Required_: Yes

_Type_: Array of [GuardrailTopicConfig](aws-properties-wisdom-aiguardrail-guardrailtopicconfig.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AIGuardrailSensitiveInformationPolicyConfig

AIGuardrailWordPolicyConfig

All content copied from https://docs.aws.amazon.com/.
