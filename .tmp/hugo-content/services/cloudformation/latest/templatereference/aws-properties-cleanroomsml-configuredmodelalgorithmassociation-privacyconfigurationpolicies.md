This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation PrivacyConfigurationPolicies

Information about the privacy configuration policies for a configured model algorithm
association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TrainedModelExports" : TrainedModelExportsConfigurationPolicy,
  "TrainedModelInferenceJobs" : TrainedModelInferenceJobsConfigurationPolicy,
  "TrainedModels" : TrainedModelsConfigurationPolicy
}

```

### YAML

```yaml

  TrainedModelExports:
    TrainedModelExportsConfigurationPolicy
  TrainedModelInferenceJobs:
    TrainedModelInferenceJobsConfigurationPolicy
  TrainedModels:
    TrainedModelsConfigurationPolicy

```

## Properties

`TrainedModelExports`

Specifies who will receive the trained model export.

_Required_: No

_Type_: [TrainedModelExportsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsconfigurationpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrainedModelInferenceJobs`

Specifies who will receive the trained model inference jobs.

_Required_: No

_Type_: [TrainedModelInferenceJobsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrainedModels`

Specifies who will receive the trained models.

_Required_: No

_Type_: [TrainedModelsConfigurationPolicy](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivacyConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
