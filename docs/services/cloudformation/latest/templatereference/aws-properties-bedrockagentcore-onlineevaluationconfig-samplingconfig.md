This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig SamplingConfig

The sampling configuration that determines what percentage of agent traces to evaluate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SamplingPercentage" : Number
}

```

### YAML

```yaml

  SamplingPercentage: Number

```

## Properties

`SamplingPercentage`

The percentage of agent traces to sample for evaluation, ranging from 0.01% to 100%.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

SessionConfig

All content copied from https://docs.aws.amazon.com/.
