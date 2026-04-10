This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard EvaluationDetail

The evaluation details of the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Datasets" : [ String, ... ],
  "EvaluationJobArn" : String,
  "EvaluationObservation" : String,
  "Metadata" : {Key: Value, ...},
  "MetricGroups" : [ MetricGroup, ... ],
  "Name" : String
}

```

### YAML

```yaml

  Datasets:
    - String
  EvaluationJobArn: String
  EvaluationObservation: String
  Metadata:
    Key: Value
  MetricGroups:
    - MetricGroup
  Name: String

```

## Properties

`Datasets`

The location of the datasets used to evaluate the model.

_Required_: No

_Type_: Array of String

_Maximum_: `1024 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationJobArn`

The Amazon Resource Name (ARN) of the evaluation job.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationObservation`

Any observations made during the model evaluation.

_Required_: No

_Type_: String

_Maximum_: `2096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

Additional attributes associated with the evaluation results.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricGroups`

An evaluation Metric Group object.

_Required_: No

_Type_: Array of [MetricGroup](aws-properties-sagemaker-modelcard-metricgroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The evaluation job name.

_Required_: Yes

_Type_: String

_Pattern_: `.{1,63}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Content

Function

All content copied from https://docs.aws.amazon.com/.
