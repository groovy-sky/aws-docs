This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model ContainerDefinition

Describes the container, as part of model definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerHostname" : String,
  "Environment" : Json,
  "Image" : String,
  "ImageConfig" : ImageConfig,
  "InferenceSpecificationName" : String,
  "Mode" : String,
  "ModelDataSource" : ModelDataSource,
  "ModelDataUrl" : String,
  "ModelPackageName" : String,
  "MultiModelConfig" : MultiModelConfig
}

```

### YAML

```yaml

  ContainerHostname: String
  Environment: Json
  Image: String
  ImageConfig:
    ImageConfig
  InferenceSpecificationName: String
  Mode: String
  ModelDataSource:
    ModelDataSource
  ModelDataUrl: String
  ModelPackageName: String
  MultiModelConfig:
    MultiModelConfig

```

## Properties

`ContainerHostname`

This parameter is ignored for models that contain only a
`PrimaryContainer`.

When a `ContainerDefinition` is part of an inference pipeline, the value of
the parameter uniquely identifies the container for the purposes of logging and metrics.
For information, see [Use Logs and Metrics\
to Monitor an Inference Pipeline](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipeline-logs-metrics.html). If you don't specify a value for this
parameter for a `ContainerDefinition` that is part of an inference pipeline,
a unique name is automatically assigned based on the position of the
`ContainerDefinition` in the pipeline. If you specify a value for the
`ContainerHostName` for any `ContainerDefinition` that is part
of an inference pipeline, you must specify a value for the
`ContainerHostName` parameter of every `ContainerDefinition`
in that pipeline.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

The environment variables to set in the Docker container. Don't include any
sensitive data in your environment variables.

The maximum length of each key and value in the `Environment` map is
1024 bytes. The maximum length of all keys and values in the map, combined, is 32 KB. If
you pass multiple containers to a `CreateModel` request, then the maximum
length of all of their maps, combined, is also 32 KB.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Image`

The path where inference code is stored. This can be either in Amazon EC2 Container Registry or in a
Docker registry that is accessible from the same VPC that you configure for your
endpoint. If you are using your own custom algorithm instead of an algorithm provided by
SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both
`registry/repository[:tag]` and `registry/repository[@digest]`
image path formats. For more information, see [Using Your Own Algorithms with\
Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).

###### Note

The model artifacts in an Amazon S3 bucket and the Docker image for inference container
in Amazon EC2 Container Registry must be in the same region as the model or endpoint you are
creating.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageConfig`

Specifies whether the model container is in Amazon ECR or a private Docker registry
accessible from your Amazon Virtual Private Cloud (VPC). For information about storing containers in a
private Docker registry, see [Use a\
Private Docker Registry for Real-Time Inference Containers](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-containers-inference-private.html).

###### Note

The model artifacts in an Amazon S3 bucket and the Docker image for inference container
in Amazon EC2 Container Registry must be in the same region as the model or endpoint you are
creating.

_Required_: No

_Type_: [ImageConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-model-imageconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceSpecificationName`

The inference specification name in the model package version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mode`

Whether the container hosts a single model or multiple models.

_Required_: No

_Type_: String

_Allowed values_: `SingleModel | MultiModel`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataSource`

Specifies the location of ML model data to deploy.

###### Note

Currently you cannot use `ModelDataSource` in conjunction with SageMaker
batch transform, SageMaker serverless endpoints, SageMaker multi-model endpoints, and SageMaker
Marketplace.

_Required_: No

_Type_: [ModelDataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-model-modeldatasource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataUrl`

The S3 path where the model artifacts, which result from model training, are stored.
This path must point to a single gzip compressed tar archive (.tar.gz suffix). The S3
path is required for SageMaker built-in algorithms, but not if you use your own algorithms.
For more information on built-in algorithms, see [Common\
Parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html).

###### Note

The model artifacts must be in an S3 bucket that is in the same region as the
model or endpoint you are creating.

If you provide a value for this parameter, SageMaker uses AWS Security Token
Service to download model artifacts from the S3 path you provide. AWS STS
is activated in your AWS account by default. If you previously
deactivated AWS STS for a region, you need to reactivate AWS STS for that region. For more information, see [Activating and\
Deactivating AWS STS in an AWS Region](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html) in the
_AWS Identity and Access Management User_
_Guide_.

###### Important

If you use a built-in algorithm to create a model, SageMaker requires that you provide
a S3 path to the model artifacts in `ModelDataUrl`.

_Required_: No

_Type_: String

_Pattern_: `(https|s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageName`

The name or Amazon Resource Name (ARN) of the model package to use to create the
model.

_Required_: No

_Type_: String

_Pattern_: `(arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:[a-z\-]*\/)?([a-zA-Z0-9]([a-zA-Z0-9-]){0,62})(?<!-)(\/[0-9]{1,9})?`

_Minimum_: `1`

_Maximum_: `176`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiModelConfig`

Specifies additional configuration for multi-model endpoints.

_Required_: No

_Type_: [MultiModelConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-model-multimodelconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::Model

HubAccessConfig
