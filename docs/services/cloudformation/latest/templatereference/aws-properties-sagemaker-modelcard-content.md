This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard Content

The content of the model card. It follows the [model card json schema](../../../sagemaker/latest/dg/model-cards.md#model-cards-json-schema).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalInformation" : AdditionalInformation,
  "BusinessDetails" : BusinessDetails,
  "EvaluationDetails" : [ EvaluationDetail, ... ],
  "IntendedUses" : IntendedUses,
  "ModelOverview" : ModelOverview,
  "ModelPackageDetails" : ModelPackageDetails,
  "TrainingDetails" : TrainingDetails
}

```

### YAML

```yaml

  AdditionalInformation:
    AdditionalInformation
  BusinessDetails:
    BusinessDetails
  EvaluationDetails:
    - EvaluationDetail
  IntendedUses:
    IntendedUses
  ModelOverview:
    ModelOverview
  ModelPackageDetails:
    ModelPackageDetails
  TrainingDetails:
    TrainingDetails

```

## Properties

`AdditionalInformation`

Additional information about the model.

_Required_: No

_Type_: [AdditionalInformation](aws-properties-sagemaker-modelcard-additionalinformation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessDetails`

Information about how the model supports business goals.

_Required_: No

_Type_: [BusinessDetails](aws-properties-sagemaker-modelcard-businessdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationDetails`

An overview about the model's evaluation.

_Required_: No

_Type_: Array of [EvaluationDetail](aws-properties-sagemaker-modelcard-evaluationdetail.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntendedUses`

The intended usage of the model.

_Required_: No

_Type_: [IntendedUses](aws-properties-sagemaker-modelcard-intendeduses.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelOverview`

An overview about the model

_Required_: No

_Type_: [ModelOverview](aws-properties-sagemaker-modelcard-modeloverview.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageDetails`

Details about the model package.

_Required_: No

_Type_: [ModelPackageDetails](aws-properties-sagemaker-modelcard-modelpackagedetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingDetails`

An overview about model training.

_Required_: No

_Type_: [TrainingDetails](aws-properties-sagemaker-modelcard-trainingdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Container

EvaluationDetail

All content copied from https://docs.aws.amazon.com/.
