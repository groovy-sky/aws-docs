This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment ModelVariantConfig

Contains information about the deployment options of a model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfrastructureConfig" : ModelInfrastructureConfig,
  "ModelName" : String,
  "VariantName" : String
}

```

### YAML

```yaml

  InfrastructureConfig:
    ModelInfrastructureConfig
  ModelName: String
  VariantName: String

```

## Properties

`InfrastructureConfig`

The configuration for the infrastructure that the model will be deployed to.

_Required_: Yes

_Type_: [ModelInfrastructureConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelName`

The name of the Amazon SageMaker Model entity.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariantName`

The name of the variant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([\-a-zA-Z0-9]*[a-zA-Z0-9])?`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelInfrastructureConfig

RealTimeInferenceConfig
