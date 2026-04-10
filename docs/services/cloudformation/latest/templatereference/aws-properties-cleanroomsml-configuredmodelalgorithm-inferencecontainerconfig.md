This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithm InferenceContainerConfig

Provides configuration information for the inference container that is used when you run an inference job on a configured model algorithm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageUri" : String
}

```

### YAML

```yaml

  ImageUri: String

```

## Properties

`ImageUri`

The registry path of the docker image that contains the inference algorithm. Clean Rooms ML
currently only
supports
the
`registry/repository[:tag]`
image path
format. For
more information about using images in Clean Rooms ML, see the [Sagemaker API reference](../../../../reference/sagemaker/latest/apireference/api-algorithmspecification.md#sagemaker-Type-AlgorithmSpecification-TrainingImage).

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerConfig

MetricDefinition

All content copied from https://docs.aws.amazon.com/.
