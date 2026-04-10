This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelPackageContainerDefinition

Describes the Docker container for the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerHostname" : String,
  "Environment" : {Key: Value, ...},
  "Framework" : String,
  "FrameworkVersion" : String,
  "Image" : String,
  "ImageDigest" : String,
  "ModelDataSource" : ModelDataSource,
  "ModelDataUrl" : String,
  "ModelInput" : ModelInput,
  "NearestModelName" : String
}

```

### YAML

```yaml

  ContainerHostname: String
  Environment:
    Key: Value
  Framework: String
  FrameworkVersion: String
  Image: String
  ImageDigest: String
  ModelDataSource:
    ModelDataSource
  ModelDataUrl: String
  ModelInput:
    ModelInput
  NearestModelName: String

```

## Properties

`ContainerHostname`

The DNS host name for the Docker container.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to set in the Docker container. Each key and value in the
`Environment` string to string map can have length of up to 1024. We
support up to 16 entries in the map.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Framework`

The machine learning framework of the model package container image.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkVersion`

The framework version of the Model Package Container Image.

_Required_: No

_Type_: String

_Pattern_: `[0-9]\.[A-Za-z0-9.]+`

_Minimum_: `3`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The Amazon Elastic Container Registry (Amazon ECR) path where inference code is stored.

If you are using your own custom algorithm instead of an algorithm provided by SageMaker,
the inference code must meet SageMaker requirements. SageMaker supports both
`registry/repository[:tag]` and `registry/repository[@digest]`
image path formats. For more information, see [Using Your Own Algorithms with Amazon\
SageMaker](../../../sagemaker/latest/dg/your-algorithms.md).

_Required_: Yes

_Type_: String

_Pattern_: `[\S]{1,255}`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageDigest`

An MD5 hash of the training algorithm that identifies the Docker image used for
training.

_Required_: No

_Type_: String

_Pattern_: `^[Ss][Hh][Aa]256:[0-9a-fA-F]{64}$`

_Maximum_: `72`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDataSource`

Specifies the location of ML model data to deploy during endpoint creation.

_Required_: No

_Type_: [ModelDataSource](aws-properties-sagemaker-modelpackage-modeldatasource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDataUrl`

The Amazon S3 path where the model artifacts, which result from model training, are stored.
This path must point to a single `gzip` compressed tar archive
( `.tar.gz` suffix).

###### Note

The model artifacts must be in an S3 bucket that is in the same region as the
model package.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelInput`

A structure with Model Input details.

_Required_: No

_Type_: [ModelInput](aws-properties-sagemaker-modelpackage-modelinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NearestModelName`

The name of a pre-trained machine learning benchmarked by
Amazon SageMaker Inference Recommender model that matches your model.
You can find a list of benchmarked models by calling `ListModelMetadata`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelMetrics

ModelPackageStatusDetails

All content copied from https://docs.aws.amazon.com/.
