This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail AIGuardrailSensitiveInformationPolicyConfig

Sensitive information policy configuration for a guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PiiEntitiesConfig" : [ GuardrailPiiEntityConfig, ... ],
  "RegexesConfig" : [ GuardrailRegexConfig, ... ]
}

```

### YAML

```yaml

  PiiEntitiesConfig:
    - GuardrailPiiEntityConfig
  RegexesConfig:
    - GuardrailRegexConfig

```

## Properties

`PiiEntitiesConfig`

List of entities.

_Required_: No

_Type_: Array of [GuardrailPiiEntityConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-guardrailpiientityconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexesConfig`

List of regex.

_Required_: No

_Type_: Array of [GuardrailRegexConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-guardrailregexconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AIGuardrailContextualGroundingPolicyConfig

AIGuardrailTopicPolicyConfig
