This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard Function

Function details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "Facet" : String,
  "Function" : String
}

```

### YAML

```yaml

  Condition: String
  Facet: String
  Function: String

```

## Properties

`Condition`

An optional description of any conditions of your objective function metric.

_Required_: No

_Type_: String

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Facet`

The metric of the model's objective function. For example, _loss_ or
_rmse_. The following list shows examples of the values that you can specify for the
metric:

- `ACCURACY`

- `AUC`

- `LOSS`

- `MAE`

- `RMSE`

_Required_: No

_Type_: String

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Function`

The optimization direction of the model's objective function. You must specify one of the following
values:

- `Maximize`

- `Minimize`

_Required_: No

_Type_: [String](aws-properties-sagemaker-modelcard-function.md)

_Allowed values_: `Maximize | Minimize`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationDetail

InferenceEnvironment

All content copied from https://docs.aws.amazon.com/.
