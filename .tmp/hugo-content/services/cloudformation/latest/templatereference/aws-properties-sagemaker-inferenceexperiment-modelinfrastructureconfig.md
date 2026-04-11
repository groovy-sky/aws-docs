This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment ModelInfrastructureConfig

The configuration for the infrastructure that the model will be deployed to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfrastructureType" : String,
  "RealTimeInferenceConfig" : RealTimeInferenceConfig
}

```

### YAML

```yaml

  InfrastructureType: String
  RealTimeInferenceConfig:
    RealTimeInferenceConfig

```

## Properties

`InfrastructureType`

The inference option to which to deploy your model. Possible values are the following:

- `RealTime`: Deploy to real-time inference.

_Required_: Yes

_Type_: String

_Allowed values_: `RealTimeInference`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RealTimeInferenceConfig`

The infrastructure configuration for deploying the model to real-time inference.

_Required_: Yes

_Type_: [RealTimeInferenceConfig](aws-properties-sagemaker-inferenceexperiment-realtimeinferenceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceExperimentSchedule

ModelVariantConfig

All content copied from https://docs.aws.amazon.com/.
