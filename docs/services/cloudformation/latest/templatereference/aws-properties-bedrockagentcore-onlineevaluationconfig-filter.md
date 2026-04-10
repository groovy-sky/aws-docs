This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig Filter

A filter that applies conditions to agent traces during online evaluation to determine which traces should be evaluated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Operator" : String,
  "Value" : FilterValue
}

```

### YAML

```yaml

  Key: String
  Operator: String
  Value:
    FilterValue

```

## Properties

`Key`

The key or field name to filter on within the agent trace data.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The comparison operator to use for filtering. Valid values include `Equals`, `NotEquals`, `GreaterThan`, `LessThan`, `GreaterThanOrEqual`, `LessThanOrEqual`, `Contains`, and `NotContains`.

_Required_: Yes

_Type_: String

_Allowed values_: `Equals | NotEquals | GreaterThan | LessThan | GreaterThanOrEqual | LessThanOrEqual | Contains | NotContains`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to compare against using the specified operator.

_Required_: Yes

_Type_: [FilterValue](aws-properties-bedrockagentcore-onlineevaluationconfig-filtervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluatorReference

FilterValue

All content copied from https://docs.aws.amazon.com/.
