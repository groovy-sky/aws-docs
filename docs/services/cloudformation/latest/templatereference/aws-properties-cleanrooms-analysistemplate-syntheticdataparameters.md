This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate SyntheticDataParameters

The parameters that control how synthetic data is generated, including privacy settings, column classifications, and other configuration options that affect the data synthesis process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MlSyntheticDataParameters" : MLSyntheticDataParameters
}

```

### YAML

```yaml

  MlSyntheticDataParameters:
    MLSyntheticDataParameters

```

## Properties

`MlSyntheticDataParameters`

The machine learning-specific parameters for synthetic data generation.

_Required_: Yes

_Type_: [MLSyntheticDataParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-analysistemplate-mlsyntheticdataparameters.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SyntheticDataColumnProperties

Tag
