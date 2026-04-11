This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow GuardrailConfiguration

Configuration information for a guardrail that you use with the [Converse](../../../../reference/bedrock/latest/apireference/api-runtime-converse.md) operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GuardrailIdentifier" : String,
  "GuardrailVersion" : String
}

```

### YAML

```yaml

  GuardrailIdentifier: String
  GuardrailVersion: String

```

## Properties

`GuardrailIdentifier`

The identifier for the guardrail.

_Required_: No

_Type_: String

_Pattern_: `^(([a-z0-9]+)|(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail/[a-z0-9]+))$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardrailVersion`

The version of the guardrail.

_Required_: No

_Type_: String

_Pattern_: `^(([0-9]{1,8})|(DRAFT))$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowValidation

InlineCodeFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
