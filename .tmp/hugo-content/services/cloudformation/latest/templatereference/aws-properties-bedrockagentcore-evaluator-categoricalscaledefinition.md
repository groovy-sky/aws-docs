This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator CategoricalScaleDefinition

The definition of a categorical rating scale option that provides a named category with its description for evaluation scoring.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : String,
  "Label" : String
}

```

### YAML

```yaml

  Definition: String
  Label: String

```

## Properties

`Definition`

The description that explains what this categorical rating represents and when it should be used.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label or name of this categorical rating option.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockEvaluatorModelConfig

CodeBasedEvaluatorConfig

All content copied from https://docs.aws.amazon.com/.
