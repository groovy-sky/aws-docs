This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig EvaluatorReference

A reference to an evaluator used in the online evaluation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EvaluatorId" : String
}

```

### YAML

```yaml

  EvaluatorId: String

```

## Properties

`EvaluatorId`

The unique identifier of the evaluator. Can reference builtin evaluators (e.g., `Builtin.Helpfulness`) or custom evaluators.

_Required_: Yes

_Type_: String

_Pattern_: `^(Builtin\.[a-zA-Z0-9_-]+|[a-zA-Z][a-zA-Z0-9-_]{0,99}-[a-zA-Z0-9]{10})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceConfig

Filter

All content copied from https://docs.aws.amazon.com/.
