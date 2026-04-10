This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard TrainingDetails

The training details of the model

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectiveFunction" : ObjectiveFunction,
  "TrainingJobDetails" : TrainingJobDetails,
  "TrainingObservations" : String
}

```

### YAML

```yaml

  ObjectiveFunction:
    ObjectiveFunction
  TrainingJobDetails:
    TrainingJobDetails
  TrainingObservations: String

```

## Properties

`ObjectiveFunction`

The function that is optimized during model training.

_Required_: No

_Type_: [ObjectiveFunction](aws-properties-sagemaker-modelcard-objectivefunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingJobDetails`

Details about any associated training jobs.

_Required_: No

_Type_: [TrainingJobDetails](aws-properties-sagemaker-modelcard-trainingjobdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingObservations`

Any observations about training.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TrainingEnvironment

All content copied from https://docs.aws.amazon.com/.
