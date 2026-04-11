This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail ContextualGroundingPolicyConfig

The policy configuration details for the guardrails contextual grounding policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FiltersConfig" : [ ContextualGroundingFilterConfig, ... ]
}

```

### YAML

```yaml

  FiltersConfig:
    - ContextualGroundingFilterConfig

```

## Properties

`FiltersConfig`

Property description not available.

_Required_: Yes

_Type_: Array of [ContextualGroundingFilterConfig](aws-properties-bedrock-guardrail-contextualgroundingfilterconfig.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContextualGroundingFilterConfig

GuardrailCrossRegionConfig

All content copied from https://docs.aws.amazon.com/.
