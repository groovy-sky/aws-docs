This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail TopicsTierConfig

The tier that your guardrail uses for denied topic filters. Consider using a tier that balances performance, accuracy, and compatibility with your existing generative AI workflows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TierName" : String
}

```

### YAML

```yaml

  TierName: String

```

## Properties

`TierName`

The tier that your guardrail uses for denied topic filters. Valid values include:

- `CLASSIC` tier – Provides established guardrails functionality supporting English, French, and Spanish languages.

- `STANDARD` tier – Provides a more robust solution than the `CLASSIC` tier and has more comprehensive language support. This tier requires that your guardrail use [cross-Region inference](../../../bedrock/latest/userguide/guardrails-cross-region.md).

_Required_: Yes

_Type_: String

_Allowed values_: `CLASSIC | STANDARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicPolicyConfig

WordConfig

All content copied from https://docs.aws.amazon.com/.
