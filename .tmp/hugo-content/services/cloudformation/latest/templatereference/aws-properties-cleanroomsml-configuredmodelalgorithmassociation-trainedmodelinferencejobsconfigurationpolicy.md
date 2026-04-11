This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelInferenceJobsConfigurationPolicy

Provides configuration information for the trained model inference job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerLogs" : [ LogsConfigurationPolicy, ... ],
  "MaxOutputSize" : TrainedModelInferenceMaxOutputSize
}

```

### YAML

```yaml

  ContainerLogs:
    - LogsConfigurationPolicy
  MaxOutputSize:
    TrainedModelInferenceMaxOutputSize

```

## Properties

`ContainerLogs`

The logs container for the trained model inference job.

_Required_: No

_Type_: Array of [LogsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxOutputSize`

The maximum allowed size of the output of the trained model inference job.

_Required_: No

_Type_: [TrainedModelInferenceMaxOutputSize](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencemaxoutputsize.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainedModelExportsMaxSize

TrainedModelInferenceMaxOutputSize

All content copied from https://docs.aws.amazon.com/.
