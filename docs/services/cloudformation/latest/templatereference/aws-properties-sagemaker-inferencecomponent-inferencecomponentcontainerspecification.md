This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentContainerSpecification

Defines a container that provides the runtime environment for a model that you deploy
with an inference component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArtifactUrl" : String,
  "DeployedImage" : DeployedImage,
  "Environment" : {Key: Value, ...},
  "Image" : String
}

```

### YAML

```yaml

  ArtifactUrl: String
  DeployedImage:
    DeployedImage
  Environment:
    Key: Value
  Image: String

```

## Properties

`ArtifactUrl`

The Amazon S3 path where the model artifacts, which result from model training,
are stored. This path must point to a single gzip compressed tar archive (.tar.gz
suffix).

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeployedImage`

The deployed container image for the inference component.

_Required_: No

_Type_: [DeployedImage](aws-properties-sagemaker-inferencecomponent-deployedimage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to set in the Docker container. Each key and value in the
Environment string-to-string map can have length of up to 1024. We support up to 16 entries
in the map.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z_][a-zA-Z0-9_]{1,1024}$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The Amazon Elastic Container Registry (Amazon ECR) path where the Docker image for the model is stored.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceComponentComputeResourceRequirements

InferenceComponentDeploymentConfig

All content copied from https://docs.aws.amazon.com/.
