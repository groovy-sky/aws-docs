This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service ImageRepository

Describes a source image repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageConfiguration" : ImageConfiguration,
  "ImageIdentifier" : String,
  "ImageRepositoryType" : String
}

```

### YAML

```yaml

  ImageConfiguration:
    ImageConfiguration
  ImageIdentifier: String
  ImageRepositoryType: String

```

## Properties

`ImageConfiguration`

Configuration for running the identified image.

_Required_: No

_Type_: [ImageConfiguration](aws-properties-apprunner-service-imageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageIdentifier`

The identifier of an image.

For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](../../../amazonecr/latest/userguide/docker-pull-ecr-image.md) in the _Amazon ECR User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/.*)|(^public\.ecr\.aws\/.+\/.+)`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageRepositoryType`

The type of the image repository. This reflects the repository provider and whether the repository is private or public.

_Required_: Yes

_Type_: String

_Allowed values_: `ECR | ECR_PUBLIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageConfiguration

IngressConfiguration

All content copied from https://docs.aws.amazon.com/.
