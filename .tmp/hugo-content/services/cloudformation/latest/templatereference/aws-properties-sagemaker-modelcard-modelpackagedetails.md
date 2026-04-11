This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard ModelPackageDetails

Details about the model package associated with the model card.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApprovalDescription" : String,
  "CreatedBy" : ModelPackageCreator,
  "Domain" : String,
  "InferenceSpecification" : InferenceSpecification,
  "ModelApprovalStatus" : String,
  "ModelPackageArn" : String,
  "ModelPackageDescription" : String,
  "ModelPackageGroupName" : String,
  "ModelPackageName" : String,
  "ModelPackageStatus" : String,
  "ModelPackageVersion" : Number,
  "SourceAlgorithms" : [ SourceAlgorithm, ... ],
  "Task" : String
}

```

### YAML

```yaml

  ApprovalDescription: String
  CreatedBy:
    ModelPackageCreator
  Domain: String
  InferenceSpecification:
    InferenceSpecification
  ModelApprovalStatus: String
  ModelPackageArn: String
  ModelPackageDescription: String
  ModelPackageGroupName: String
  ModelPackageName: String
  ModelPackageStatus: String
  ModelPackageVersion: Number
  SourceAlgorithms:
    - SourceAlgorithm
  Task: String

```

## Properties

`ApprovalDescription`

A description of the approval status of the model package.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedBy`

Information about the user who created the model package.

_Required_: No

_Type_: [ModelPackageCreator](aws-properties-sagemaker-modelcard-modelpackagecreator.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The machine learning domain of the model package.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceSpecification`

Details about the inference specification for the model package.

_Required_: No

_Type_: [InferenceSpecification](aws-properties-sagemaker-modelcard-inferencespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelApprovalStatus`

The approval status of the model package.

_Required_: No

_Type_: String

_Allowed values_: `Approved | Rejected | PendingManualApproval`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageArn`

The Amazon Resource Name (ARN) of the model package.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageDescription`

A description of the model package.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageGroupName`

The name of the model package group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageName`

The name of the model package.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageStatus`

The status of the model package.

_Required_: No

_Type_: String

_Allowed values_: `Pending | InProgress | Completed | Failed | Deleting`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageVersion`

The version of the model package.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceAlgorithms`

A list of algorithms that were used to create the model package.

_Required_: No

_Type_: Array of [SourceAlgorithm](aws-properties-sagemaker-modelcard-sourcealgorithm.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Task`

The machine learning task performed by the model package.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelPackageCreator

ObjectiveFunction

All content copied from https://docs.aws.amazon.com/.
