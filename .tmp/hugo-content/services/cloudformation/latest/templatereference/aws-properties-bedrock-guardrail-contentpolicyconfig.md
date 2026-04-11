This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail ContentPolicyConfig

Contains details about how to handle harmful content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentFiltersTierConfig" : ContentFiltersTierConfig,
  "FiltersConfig" : [ ContentFilterConfig, ... ]
}

```

### YAML

```yaml

  ContentFiltersTierConfig:
    ContentFiltersTierConfig
  FiltersConfig:
    - ContentFilterConfig

```

## Properties

`ContentFiltersTierConfig`

The tier that your guardrail uses for content filters. Consider using a tier that balances performance, accuracy, and compatibility with your existing generative AI workflows.

_Required_: No

_Type_: [ContentFiltersTierConfig](aws-properties-bedrock-guardrail-contentfilterstierconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FiltersConfig`

Contains the type of the content filter and how strongly it should apply to prompts and model responses.

_Required_: Yes

_Type_: Array of [ContentFilterConfig](aws-properties-bedrock-guardrail-contentfilterconfig.md)

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentFiltersTierConfig

ContextualGroundingFilterConfig

All content copied from https://docs.aws.amazon.com/.
