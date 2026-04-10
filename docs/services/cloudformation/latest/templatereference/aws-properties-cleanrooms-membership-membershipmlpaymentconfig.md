This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipMLPaymentConfig

An object representing the collaboration member's machine learning payment responsibilities set by the
collaboration creator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelInference" : MembershipModelInferencePaymentConfig,
  "ModelTraining" : MembershipModelTrainingPaymentConfig,
  "SyntheticDataGeneration" : MembershipSyntheticDataGenerationPaymentConfig
}

```

### YAML

```yaml

  ModelInference:
    MembershipModelInferencePaymentConfig
  ModelTraining:
    MembershipModelTrainingPaymentConfig
  SyntheticDataGeneration:
    MembershipSyntheticDataGenerationPaymentConfig

```

## Properties

`ModelInference`

The payment responsibilities accepted by the member for model inference.

_Required_: No

_Type_: [MembershipModelInferencePaymentConfig](aws-properties-cleanrooms-membership-membershipmodelinferencepaymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelTraining`

The payment responsibilities accepted by the member for model training.

_Required_: No

_Type_: [MembershipModelTrainingPaymentConfig](aws-properties-cleanrooms-membership-membershipmodeltrainingpaymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyntheticDataGeneration`

The payment configuration for synthetic data generation for this machine learning membership.

_Required_: No

_Type_: [MembershipSyntheticDataGenerationPaymentConfig](aws-properties-cleanrooms-membership-membershipsyntheticdatagenerationpaymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipJobComputePaymentConfig

MembershipModelInferencePaymentConfig

All content copied from https://docs.aws.amazon.com/.
