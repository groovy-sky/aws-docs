This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail AIGuardrailContextualGroundingPolicyConfig

Contextual grounding policy config for a guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FiltersConfig" : [ GuardrailContextualGroundingFilterConfig, ... ]
}

```

### YAML

```yaml

  FiltersConfig:
    - GuardrailContextualGroundingFilterConfig

```

## Properties

`FiltersConfig`

List of contextual grounding filter configs.

_Required_: Yes

_Type_: Array of [GuardrailContextualGroundingFilterConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-aiguardrail-guardrailcontextualgroundingfilterconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AIGuardrailContentPolicyConfig

AIGuardrailSensitiveInformationPolicyConfig
