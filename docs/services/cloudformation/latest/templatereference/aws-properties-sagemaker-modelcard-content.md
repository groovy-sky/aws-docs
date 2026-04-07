This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard Content

The content of the model card. It follows the [model card json schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema).

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

_Type_: [AdditionalInformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-additionalinformation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessDetails`

Information about how the model supports business goals.

_Required_: No

_Type_: [BusinessDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-businessdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationDetails`

An overview about the model's evaluation.

_Required_: No

_Type_: Array of [EvaluationDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-evaluationdetail.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntendedUses`

The intended usage of the model.

_Required_: No

_Type_: [IntendedUses](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-intendeduses.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelOverview`

An overview about the model

_Required_: No

_Type_: [ModelOverview](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-modeloverview.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageDetails`

Details about the model package.

_Required_: No

_Type_: [ModelPackageDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-modelpackagedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingDetails`

An overview about model training.

_Required_: No

_Type_: [TrainingDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelcard-trainingdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Container

EvaluationDetail
