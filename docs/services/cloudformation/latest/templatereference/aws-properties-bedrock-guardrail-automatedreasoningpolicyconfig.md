This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail AutomatedReasoningPolicyConfig

Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfidenceThreshold" : Number,
  "Policies" : [ String, ... ]
}

```

### YAML

```yaml

  ConfidenceThreshold: Number
  Policies:
    - String

```

## Properties

`ConfidenceThreshold`

The minimum confidence level required for Automated Reasoning policy violations to trigger guardrail actions. Values range from 0.0 to 1.0.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

The list of Automated Reasoning policy ARNs that should be applied as part of this guardrail configuration.

_Required_: Yes

_Type_: Array of String

_Minimum_: `15 | 1`

_Maximum_: `2048 | 2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Bedrock::Guardrail

ContentFilterConfig
