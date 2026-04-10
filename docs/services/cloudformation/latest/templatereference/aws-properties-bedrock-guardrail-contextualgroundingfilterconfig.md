This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail ContextualGroundingFilterConfig

The filter configuration details for the guardrails contextual grounding filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Enabled" : Boolean,
  "Threshold" : Number,
  "Type" : String
}

```

### YAML

```yaml

  Action: String
  Enabled: Boolean
  Threshold: Number
  Type: String

```

## Properties

`Action`

Specifies the action to take when content fails the contextual grounding evaluation. Supported values include:

- `BLOCK` – Block the content and replace it with blocked
messaging.

- `NONE` – Take no action but return detection information in the trace
response.

_Required_: No

_Type_: String

_Allowed values_: `BLOCK | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether to enable contextual grounding evaluation. When disabled, you aren't charged for the evaluation. The evaluation doesn't appear in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The threshold details for the guardrails contextual grounding filter.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The filter details for the guardrails contextual grounding filter.

_Required_: Yes

_Type_: String

_Allowed values_: `GROUNDING | RELEVANCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentPolicyConfig

ContextualGroundingPolicyConfig

All content copied from https://docs.aws.amazon.com/.
