This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail WordConfig

A word to configure for the guardrail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputAction" : String,
  "InputEnabled" : Boolean,
  "OutputAction" : String,
  "OutputEnabled" : Boolean,
  "Text" : String
}

```

### YAML

```yaml

  InputAction: String
  InputEnabled: Boolean
  OutputAction: String
  OutputEnabled: Boolean
  Text: String

```

## Properties

`InputAction`

Specifies the action to take when harmful content is detected in the input. Supported values include:

- `BLOCK` – Block the content and replace it with blocked
messaging.

- `NONE` – Take no action but return detection information in the trace
response.

_Required_: No

_Type_: String

_Allowed values_: `BLOCK | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputEnabled`

Specifies whether to enable guardrail evaluation on the intput. When disabled, you
aren't charged for the evaluation. The evaluation doesn't appear in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputAction`

Specifies the action to take when harmful content is detected in the output. Supported values include:

- `BLOCK` – Block the content and replace it with blocked
messaging.

- `NONE` – Take no action but return detection information in the trace
response.

_Required_: No

_Type_: String

_Allowed values_: `BLOCK | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputEnabled`

Specifies whether to enable guardrail evaluation on the output. When disabled, you
aren't charged for the evaluation. The evaluation doesn't appear in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

Text of the word configured for the guardrail to block.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicsTierConfig

WordPolicyConfig

All content copied from https://docs.aws.amazon.com/.
