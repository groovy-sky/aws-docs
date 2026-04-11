This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration MLPaymentConfig

An object representing the collaboration member's machine learning payment responsibilities set by the
collaboration creator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelInference" : ModelInferencePaymentConfig,
  "ModelTraining" : ModelTrainingPaymentConfig,
  "SyntheticDataGeneration" : SyntheticDataGenerationPaymentConfig
}

```

### YAML

```yaml

  ModelInference:
    ModelInferencePaymentConfig
  ModelTraining:
    ModelTrainingPaymentConfig
  SyntheticDataGeneration:
    SyntheticDataGenerationPaymentConfig

```

## Properties

`ModelInference`

The payment responsibilities accepted by the member for model inference.

_Required_: No

_Type_: [ModelInferencePaymentConfig](aws-properties-cleanrooms-collaboration-modelinferencepaymentconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelTraining`

The payment responsibilities accepted by the member for model training.

_Required_: No

_Type_: [ModelTrainingPaymentConfig](aws-properties-cleanrooms-collaboration-modeltrainingpaymentconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyntheticDataGeneration`

The payment configuration for machine learning synthetic data generation.

_Required_: No

_Type_: [SyntheticDataGenerationPaymentConfig](aws-properties-cleanrooms-collaboration-syntheticdatagenerationpaymentconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MLMemberAbilities

ModelInferencePaymentConfig

All content copied from https://docs.aws.amazon.com/.
