This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelInferenceMaxOutputSize

Information about the maximum output size for a trained model inference job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Unit: String
  Value: Number

```

## Properties

`Unit`

The measurement unit to use.

_Required_: Yes

_Type_: String

_Allowed values_: `GB`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The maximum output size value.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainedModelInferenceJobsConfigurationPolicy

TrainedModelsConfigurationPolicy

All content copied from https://docs.aws.amazon.com/.
