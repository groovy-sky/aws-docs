This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage

A container for your trained model that can be deployed for SageMaker inference. This can
include inference code, artifacts, and metadata. The model package type can be one of
the following.

- Versioned model: A part of a model package group in Model Registry.

- Unversioned model: Not part of a model package group and used in AWS Marketplace.

For more information, see [`CreateModelPackage`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModelPackage.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::ModelPackage",
  "Properties" : {
      "AdditionalInferenceSpecifications" : [ AdditionalInferenceSpecificationDefinition, ... ],
      "AdditionalInferenceSpecificationsToAdd" : [ AdditionalInferenceSpecificationDefinition, ... ],
      "ApprovalDescription" : String,
      "CertifyForMarketplace" : Boolean,
      "ClientToken" : String,
      "CustomerMetadataProperties" : {Key: Value, ...},
      "Domain" : String,
      "DriftCheckBaselines" : DriftCheckBaselines,
      "InferenceSpecification" : InferenceSpecification,
      "MetadataProperties" : MetadataProperties,
      "ModelApprovalStatus" : String,
      "ModelCard" : ModelCard,
      "ModelMetrics" : ModelMetrics,
      "ModelPackageDescription" : String,
      "ModelPackageGroupName" : String,
      "ModelPackageName" : String,
      "ModelPackageStatusDetails" : ModelPackageStatusDetails,
      "ModelPackageVersion" : Integer,
      "SamplePayloadUrl" : String,
      "SecurityConfig" : SecurityConfig,
      "SkipModelValidation" : String,
      "SourceAlgorithmSpecification" : SourceAlgorithmSpecification,
      "SourceUri" : String,
      "Tags" : [ Tag, ... ],
      "Task" : String,
      "ValidationSpecification" : ValidationSpecification
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::ModelPackage
Properties:
  AdditionalInferenceSpecifications:
    - AdditionalInferenceSpecificationDefinition
  AdditionalInferenceSpecificationsToAdd:
    - AdditionalInferenceSpecificationDefinition
  ApprovalDescription: String
  CertifyForMarketplace: Boolean
  ClientToken: String
  CustomerMetadataProperties:
    Key: Value
  Domain: String
  DriftCheckBaselines:
    DriftCheckBaselines
  InferenceSpecification:
    InferenceSpecification
  MetadataProperties:
    MetadataProperties
  ModelApprovalStatus: String
  ModelCard:
    ModelCard
  ModelMetrics:
    ModelMetrics
  ModelPackageDescription: String
  ModelPackageGroupName: String
  ModelPackageName: String
  ModelPackageStatusDetails:
    ModelPackageStatusDetails
  ModelPackageVersion: Integer
  SamplePayloadUrl: String
  SecurityConfig:
    SecurityConfig
  SkipModelValidation: String
  SourceAlgorithmSpecification:
    SourceAlgorithmSpecification
  SourceUri: String
  Tags:
    - Tag
  Task: String
  ValidationSpecification:
    ValidationSpecification

```

## Properties

`AdditionalInferenceSpecifications`

An array of additional Inference Specification objects.

_Required_: No

_Type_: Array of [AdditionalInferenceSpecificationDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalInferenceSpecificationsToAdd`

An array of additional Inference Specification objects to be added to the existing array. The total number of
additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies
artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to
store the compiled artifacts.

_Required_: No

_Type_: Array of [AdditionalInferenceSpecificationDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApprovalDescription`

A description provided when the model approval is set.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertifyForMarketplace`

Whether the model package is to be certified to be listed on AWS
Marketplace. For information about listing model packages on AWS
Marketplace, see [List Your Algorithm or Model\
Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientToken`

A unique token that guarantees that the call to this API is idempotent.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerMetadataProperties`

The metadata properties for the model package.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:\/=+\-@]*)${1,128}`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The machine learning domain of your model package and its components. Common machine
learning domains include computer vision and natural language processing.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DriftCheckBaselines`

Represents the drift check baselines that can be used when the model monitor is set
using the model package.

_Required_: No

_Type_: [DriftCheckBaselines](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceSpecification`

Defines how to perform inference generation after a training job is run.

_Required_: No

_Type_: [InferenceSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-inferencespecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataProperties`

Metadata properties of the tracking entity, trial, or trial component.

_Required_: No

_Type_: [MetadataProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-metadataproperties.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelApprovalStatus`

The approval status of the model. This can be one of the following values.

- `APPROVED` \- The model is approved

- `REJECTED` \- The model is rejected.

- `PENDING_MANUAL_APPROVAL` \- The model is waiting for manual
approval.

_Required_: No

_Type_: String

_Allowed values_: `Approved | Rejected | PendingManualApproval`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelCard`

An Amazon SageMaker Model Card.

_Required_: No

_Type_: [ModelCard](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modelcard.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ModelMetrics`

Metrics for the model.

_Required_: No

_Type_: [ModelMetrics](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modelmetrics.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageDescription`

The description of the model package.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageGroupName`

The model group to which the model belongs.

_Required_: No

_Type_: String

_Pattern_: `(arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:[a-z\-]*\/)?([a-zA-Z0-9]([a-zA-Z0-9-]){0,62})(?<!-)$`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageName`

The name of the model package. The name can be as follows:

- For a versioned model, the name is automatically generated by SageMaker Model Registry and
follows the format
' `ModelPackageGroupName/ModelPackageVersion`'.

- For an unversioned model, you must provide the name.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageStatusDetails`

Specifies the validation and image scan statuses of the model package.

_Required_: No

_Type_: [ModelPackageStatusDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modelpackagestatusdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelPackageVersion`

The version number of a versioned model.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplePayloadUrl`

The Amazon Simple Storage Service path where the sample payload are stored. This path must point to a
single gzip compressed tar archive (.tar.gz suffix).

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfig`

Security configuration for the model package, including encryption settings.

_Required_: No

_Type_: [SecurityConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-securityconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SkipModelValidation`

Indicates if you want to skip model validation.

_Required_: No

_Type_: String

_Allowed values_: `None | All`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceAlgorithmSpecification`

A list of algorithms that were used to create a model package.

_Required_: No

_Type_: [SourceAlgorithmSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-sourcealgorithmspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceUri`

The URI of the source for the model package.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{Z}\p{N}\p{P}]{0,1024}`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

A list of the tags associated with the model package. For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the _AWS General_
_Reference Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Task`

The machine learning task your model package accomplishes. Common machine learning
tasks include object detection and image classification.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationSpecification`

Specifies batch transform jobs that SageMaker runs to validate your model package.

_Required_: No

_Type_: [ValidationSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-validationspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the model package group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time that the model package was created.

`LastModifiedTime`

The last time the model package was modified.

`ModelPackageArn`

The Amazon Resource Name (ARN) of the model package.

`ModelPackageStatus`

The status of the model package. This can be one of the following values.

- `PENDING` \- The model package creation is pending.

- `IN_PROGRESS` \- The model package is in the process of being created.

- `COMPLETED` \- The model package was successfully created.

- `FAILED` \- The model package creation failed.

- `DELETING` \- The model package is in the process of being deleted.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

AdditionalInferenceSpecificationDefinition
