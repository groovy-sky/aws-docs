This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent DeployedImage

Gets the Amazon EC2 Container Registry path of the docker image of the model that is hosted in this [ProductionVariant](../../../../reference/sagemaker/latest/apireference/api-productionvariant.md).

If you used the `registry/repository[:tag]` form to specify the image path
of the primary container when you created the model hosted in this
`ProductionVariant`, the path resolves to a path of the form
`registry/repository[@digest]`. A digest is a hash value that identifies
a specific version of an image. For information about Amazon ECR paths, see [Pulling an Image](../../../amazonecr/latest/userguide/docker-pull-ecr-image.md) in the _Amazon ECR User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResolutionTime" : String,
  "ResolvedImage" : String,
  "SpecifiedImage" : String
}

```

### YAML

```yaml

  ResolutionTime: String
  ResolvedImage: String
  SpecifiedImage: String

```

## Properties

`ResolutionTime`

The date and time when the image path for the model resolved to the
`ResolvedImage`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolvedImage`

The specific digest path of the image hosted in this
`ProductionVariant`.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpecifiedImage`

The image path you specified when you created the model.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoRollbackConfiguration

InferenceComponentCapacitySize

All content copied from https://docs.aws.amazon.com/.
