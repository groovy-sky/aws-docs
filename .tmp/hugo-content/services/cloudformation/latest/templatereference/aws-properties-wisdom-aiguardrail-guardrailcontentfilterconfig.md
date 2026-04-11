This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail GuardrailContentFilterConfig

Content filter configuration in content policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputStrength" : String,
  "OutputStrength" : String,
  "Type" : String
}

```

### YAML

```yaml

  InputStrength: String
  OutputStrength: String
  Type: String

```

## Properties

`InputStrength`

The strength of the input for the guardrail content filter.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputStrength`

The output strength of the guardrail content filter.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the guardrail content filter.

_Required_: Yes

_Type_: String

_Allowed values_: `SEXUAL | VIOLENCE | HATE | INSULTS | MISCONDUCT | PROMPT_ATTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AIGuardrailWordPolicyConfig

GuardrailContextualGroundingFilterConfig

All content copied from https://docs.aws.amazon.com/.
