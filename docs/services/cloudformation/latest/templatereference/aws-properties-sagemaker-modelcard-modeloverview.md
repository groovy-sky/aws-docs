This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard ModelOverview

An overview about the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlgorithmType" : String,
  "InferenceEnvironment" : InferenceEnvironment,
  "ModelArtifact" : [ String, ... ],
  "ModelCreator" : String,
  "ModelDescription" : String,
  "ModelId" : String,
  "ModelName" : String,
  "ModelOwner" : String,
  "ModelVersion" : Number,
  "ProblemType" : String
}

```

### YAML

```yaml

  AlgorithmType: String
  InferenceEnvironment:
    InferenceEnvironment
  ModelArtifact:
    - String
  ModelCreator: String
  ModelDescription: String
  ModelId: String
  ModelName: String
  ModelOwner: String
  ModelVersion: Number
  ProblemType: String

```

## Properties

`AlgorithmType`

The algorithm used to solve the problem.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceEnvironment`

An overview about model inference.

_Required_: No

_Type_: [InferenceEnvironment](aws-properties-sagemaker-modelcard-inferenceenvironment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelArtifact`

The location of the model artifact.

_Required_: No

_Type_: Array of String

_Maximum_: `1024 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelCreator`

The creator of the model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDescription`

A description of the model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The SageMaker AI Model ARN or non-SageMaker AI Model ID.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelName`

The name of the model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelOwner`

The owner of the model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelVersion`

The version of the model.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProblemType`

The problem being solved with the model.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricGroup

ModelPackageCreator

All content copied from https://docs.aws.amazon.com/.
