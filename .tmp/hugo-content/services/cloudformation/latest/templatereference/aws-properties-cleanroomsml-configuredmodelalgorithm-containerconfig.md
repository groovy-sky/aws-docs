This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithm ContainerConfig

Provides configuration information for the training container, including entrypoints and arguments.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arguments" : [ String, ... ],
  "Entrypoint" : [ String, ... ],
  "ImageUri" : String,
  "MetricDefinitions" : [ MetricDefinition, ... ]
}

```

### YAML

```yaml

  Arguments:
    - String
  Entrypoint:
    - String
  ImageUri: String
  MetricDefinitions:
    - MetricDefinition

```

## Properties

`Arguments`

The arguments for a container used to run a training job. See How Amazon SageMaker Runs
Your Training Image for additional information. For more information, see [How Sagemaker\
runs your training image](../../../sagemaker/latest/dg/your-algorithms-training-algo-dockerfile.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Entrypoint`

The entrypoint script for a Docker container used to run a training job. This script
takes precedence over the default train processing instructions. See How Amazon SageMaker
Runs Your Training Image for additional information. For more information, see [How Sagemaker\
runs your training image](../../../sagemaker/latest/dg/your-algorithms-training-algo-dockerfile.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageUri`

The registry path of the docker image that contains the algorithm. Clean Rooms ML currently only
supports the `registry/repository[:tag]` image path
format. For
more information about using images in Clean Rooms ML, see the [Sagemaker API reference](../../../../reference/sagemaker/latest/apireference/api-algorithmspecification.md#sagemaker-Type-AlgorithmSpecification-TrainingImage).

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricDefinitions`

A list of metric definition objects. Each object specifies the metric name and regular
expressions used to parse algorithm logs. AWS Clean Rooms ML publishes each metric to all members'
Amazon CloudWatch using IAM role configured in [PutMLConfiguration](../../../../reference/cleanrooms-ml/latest/apireference/api-putmlconfiguration.md).

_Required_: No

_Type_: Array of [MetricDefinition](aws-properties-cleanroomsml-configuredmodelalgorithm-metricdefinition.md)

_Minimum_: `0`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRoomsML::ConfiguredModelAlgorithm

InferenceContainerConfig

All content copied from https://docs.aws.amazon.com/.
