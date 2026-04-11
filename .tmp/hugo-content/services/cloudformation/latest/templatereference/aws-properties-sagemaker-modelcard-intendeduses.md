This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard IntendedUses

The intended uses of a model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExplanationsForRiskRating" : String,
  "FactorsAffectingModelEfficiency" : String,
  "IntendedUses" : String,
  "PurposeOfModel" : String,
  "RiskRating" : String
}

```

### YAML

```yaml

  ExplanationsForRiskRating: String
  FactorsAffectingModelEfficiency: String
  IntendedUses: String
  PurposeOfModel: String
  RiskRating: String

```

## Properties

`ExplanationsForRiskRating`

An explanation of why your organization categorizes the model with its risk rating.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FactorsAffectingModelEfficiency`

Factors affecting model efficacy.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntendedUses`

The intended use cases for the model.

_Required_: No

_Type_: [String](aws-properties-sagemaker-modelcard-intendeduses.md)

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PurposeOfModel`

The general purpose of the model.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RiskRating`

Your organization's risk rating. You can specify one the following values as the risk rating:

- High

- Medium

- Low

- Unknown

_Required_: No

_Type_: String

_Allowed values_: `High | Medium | Low | Unknown`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceSpecification

MetricDataItems

All content copied from https://docs.aws.amazon.com/.
