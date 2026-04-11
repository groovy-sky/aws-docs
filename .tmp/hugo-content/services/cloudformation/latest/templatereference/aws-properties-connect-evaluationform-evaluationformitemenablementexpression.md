This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormItemEnablementExpression

An expression that defines a basic building block of conditional enablement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparator" : String,
  "Source" : EvaluationFormItemEnablementSource,
  "Values" : [ EvaluationFormItemEnablementSourceValue, ... ]
}

```

### YAML

```yaml

  Comparator: String
  Source:
    EvaluationFormItemEnablementSource
  Values:
    - EvaluationFormItemEnablementSourceValue

```

## Properties

`Comparator`

A comparator to be used against list of values.

_Required_: Yes

_Type_: String

_Allowed values_: `IN | NOT_IN | ALL_IN | EXACT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

A source item of enablement expression.

_Required_: Yes

_Type_: [EvaluationFormItemEnablementSource](aws-properties-connect-evaluationform-evaluationformitemenablementsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

A list of values from source item.

_Required_: Yes

_Type_: Array of [EvaluationFormItemEnablementSourceValue](aws-properties-connect-evaluationform-evaluationformitemenablementsourcevalue.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormItemEnablementConfiguration

EvaluationFormItemEnablementSource

All content copied from https://docs.aws.amazon.com/.
