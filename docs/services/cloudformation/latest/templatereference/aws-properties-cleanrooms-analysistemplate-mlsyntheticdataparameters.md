This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate MLSyntheticDataParameters

Parameters that control the generation of synthetic data for machine learning, including
privacy settings and column classification details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnClassification" : ColumnClassificationDetails,
  "Epsilon" : Number,
  "MaxMembershipInferenceAttackScore" : Number
}

```

### YAML

```yaml

  ColumnClassification:
    ColumnClassificationDetails
  Epsilon: Number
  MaxMembershipInferenceAttackScore: Number

```

## Properties

`ColumnClassification`

Classification details for data columns that specify how each column should be treated during synthetic data generation.

_Required_: Yes

_Type_: [ColumnClassificationDetails](aws-properties-cleanrooms-analysistemplate-columnclassificationdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Epsilon`

The epsilon value for differential privacy when generating synthetic data. Lower values provide stronger privacy guarantees but may reduce data utility.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxMembershipInferenceAttackScore`

The maximum acceptable score for membership inference attack vulnerability. Synthetic data generation fails if the score for the resulting data exceeds this threshold.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Hash

S3Location

All content copied from https://docs.aws.amazon.com/.
