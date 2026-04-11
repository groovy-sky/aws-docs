This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelsConfigurationPolicy

The configuration policy for the trained models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerLogs" : [ LogsConfigurationPolicy, ... ],
  "ContainerMetrics" : MetricsConfigurationPolicy,
  "MaxArtifactSize" : TrainedModelArtifactMaxSize
}

```

### YAML

```yaml

  ContainerLogs:
    - LogsConfigurationPolicy
  ContainerMetrics:
    MetricsConfigurationPolicy
  MaxArtifactSize:
    TrainedModelArtifactMaxSize

```

## Properties

`ContainerLogs`

The container for the logs of the trained model.

_Required_: No

_Type_: Array of [LogsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerMetrics`

The container for the metrics of the trained model.

_Required_: No

_Type_: [MetricsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-metricsconfigurationpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxArtifactSize`

The maximum size limit for trained model artifacts as defined in the configuration
policy. This setting helps enforce consistent size limits across trained models in the
collaboration.

_Required_: No

_Type_: [TrainedModelArtifactMaxSize](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelartifactmaxsize.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainedModelInferenceMaxOutputSize

AWS::CleanRoomsML::TrainingDataset

All content copied from https://docs.aws.amazon.com/.
