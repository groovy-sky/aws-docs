This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment ShadowModeConfig

The configuration of `ShadowMode` inference experiment type, which specifies a production variant
to take all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the
inference requests. For the shadow variant it also specifies the percentage of requests that Amazon SageMaker replicates.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ShadowModelVariants" : [ ShadowModelVariantConfig, ... ],
  "SourceModelVariantName" : String
}

```

### YAML

```yaml

  ShadowModelVariants:
    - ShadowModelVariantConfig
  SourceModelVariantName: String

```

## Properties

`ShadowModelVariants`

List of shadow variant configurations.

_Required_: Yes

_Type_: Array of [ShadowModelVariantConfig](aws-properties-sagemaker-inferenceexperiment-shadowmodelvariantconfig.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceModelVariantName`

The name of the production variant, which takes all the inference requests.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([\-a-zA-Z0-9]*[a-zA-Z0-9])?`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RealTimeInferenceConfig

ShadowModelVariantConfig

All content copied from https://docs.aws.amazon.com/.
