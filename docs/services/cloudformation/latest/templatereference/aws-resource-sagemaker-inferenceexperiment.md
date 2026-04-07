This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment

Creates an inference experiment using the configurations specified in the request.

Use this API to setup and schedule an experiment to compare model variants on a Amazon SageMaker inference endpoint. For
more information about inference experiments, see [Shadow tests](https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests.html).

Amazon SageMaker begins your experiment at the scheduled time and routes traffic to your endpoint's model variants based
on your specified configuration.

While the experiment is in progress or after it has concluded, you can view metrics that compare your model
variants. For more information, see [View, monitor, and edit shadow tests](https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests-view-monitor-edit.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::InferenceExperiment",
  "Properties" : {
      "DataStorageConfig" : DataStorageConfig,
      "Description" : String,
      "DesiredState" : String,
      "EndpointName" : String,
      "KmsKey" : String,
      "ModelVariants" : [ ModelVariantConfig, ... ],
      "Name" : String,
      "RoleArn" : String,
      "Schedule" : InferenceExperimentSchedule,
      "ShadowModeConfig" : ShadowModeConfig,
      "StatusReason" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::InferenceExperiment
Properties:
  DataStorageConfig:
    DataStorageConfig
  Description: String
  DesiredState: String
  EndpointName: String
  KmsKey: String
  ModelVariants:
    - ModelVariantConfig
  Name: String
  RoleArn: String
  Schedule:
    InferenceExperimentSchedule
  ShadowModeConfig:
    ShadowModeConfig
  StatusReason: String
  Tags:
    - Tag
  Type: String

```

## Properties

`DataStorageConfig`

The Amazon S3 location and configuration for storing inference request and response data.

_Required_: No

_Type_: [DataStorageConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-datastorageconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the inference experiment.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredState`

The desired state of the experiment after stopping. The possible states are the following:

- `Completed`: The experiment completed successfully

- `Cancelled`: The experiment was canceled

_Required_: No

_Type_: String

_Allowed values_: `Running | Completed | Cancelled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointName`

The name of the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKey`

The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3
server-side encryption.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelVariants`

An array of `ModelVariantConfigSummary` objects. There is one for each variant in the inference
experiment. Each `ModelVariantConfigSummary` object in the array describes the infrastructure
configuration for deploying the corresponding variant.

_Required_: Yes

_Type_: Array of [ModelVariantConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-modelvariantconfig.html)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the inference experiment.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `120`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage
Amazon SageMaker Inference endpoints for model deployment.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schedule`

The duration for which the inference experiment ran or will run.

The maximum duration that you can set for an inference experiment is 30 days.

_Required_: No

_Type_: [InferenceExperimentSchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-inferenceexperimentschedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShadowModeConfig`

The configuration of `ShadowMode` inference experiment type, which shows the production variant
that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the
inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.

_Required_: No

_Type_: [ShadowModeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusReason`

The error message for the inference experiment status result.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the inference experiment.

_Required_: Yes

_Type_: String

_Allowed values_: `ShadowMode`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the inference experiment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN for your inference experiment.

`CreationTime`

The timestamp at which the inference experiment was created.

`LastModifiedTime`

The timestamp at which you last modified the inference experiment.

`Status`

The status of the inference experiment. The following are the possible statuses for an inference
experiment:

- `Creating` \- Amazon SageMaker is creating your experiment.

- `Created` \- Amazon SageMaker has finished the creation of your experiment and will begin the
experiment at the scheduled time.

- `Updating` \- When you make changes to your experiment, your experiment shows as updating.

- `Starting` \- Amazon SageMaker is beginning your experiment.

- `Running` \- Your experiment is in progress.

- `Stopping` \- Amazon SageMaker is stopping your experiment.

- `Completed` \- Your experiment has completed.

- `Cancelled` \- When you conclude your experiment early using the [StopInferenceExperiment](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopInferenceExperiment.html) API, or if any operation fails with an unexpected error, it shows
as cancelled.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CaptureContentTypeHeader
