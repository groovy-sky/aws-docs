This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard TrainingJobDetails

The overview of a training job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HyperParameters" : [ TrainingHyperParameter, ... ],
  "TrainingArn" : String,
  "TrainingDatasets" : [ String, ... ],
  "TrainingEnvironment" : TrainingEnvironment,
  "TrainingMetrics" : [ TrainingMetric, ... ],
  "UserProvidedHyperParameters" : [ TrainingHyperParameter, ... ],
  "UserProvidedTrainingMetrics" : [ TrainingMetric, ... ]
}

```

### YAML

```yaml

  HyperParameters:
    - TrainingHyperParameter
  TrainingArn: String
  TrainingDatasets:
    - String
  TrainingEnvironment:
    TrainingEnvironment
  TrainingMetrics:
    - TrainingMetric
  UserProvidedHyperParameters:
    - TrainingHyperParameter
  UserProvidedTrainingMetrics:
    - TrainingMetric

```

## Properties

`HyperParameters`

The hyper parameters used in the training job.

_Required_: No

_Type_: Array of [TrainingHyperParameter](aws-properties-sagemaker-modelcard-traininghyperparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingArn`

The SageMaker AI training job Amazon Resource Name (ARN)

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingDatasets`

The location of the datasets used to train the model.

_Required_: No

_Type_: Array of String

_Maximum_: `1024 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingEnvironment`

The SageMaker AI training job image URI.

_Required_: No

_Type_: [TrainingEnvironment](aws-properties-sagemaker-modelcard-trainingenvironment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingMetrics`

The SageMaker AI training job results.

_Required_: No

_Type_: Array of [TrainingMetric](aws-properties-sagemaker-modelcard-trainingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProvidedHyperParameters`

Additional hyper parameters that you've specified when training the model.

_Required_: No

_Type_: Array of [TrainingHyperParameter](aws-properties-sagemaker-modelcard-traininghyperparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProvidedTrainingMetrics`

Custom training job results.

_Required_: No

_Type_: Array of [TrainingMetric](aws-properties-sagemaker-modelcard-trainingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainingHyperParameter

TrainingMetric

All content copied from https://docs.aws.amazon.com/.
