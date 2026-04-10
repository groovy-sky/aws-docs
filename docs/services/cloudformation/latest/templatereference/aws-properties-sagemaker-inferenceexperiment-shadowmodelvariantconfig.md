This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment ShadowModelVariantConfig

The name and sampling percentage of a shadow variant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SamplingPercentage" : Integer,
  "ShadowModelVariantName" : String
}

```

### YAML

```yaml

  SamplingPercentage: Integer
  ShadowModelVariantName: String

```

## Properties

`SamplingPercentage`

The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.

_Required_: Yes

_Type_: Integer

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShadowModelVariantName`

The name of the shadow variant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([\-a-zA-Z0-9]*[a-zA-Z0-9])?`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ShadowModeConfig

Tag

All content copied from https://docs.aws.amazon.com/.
