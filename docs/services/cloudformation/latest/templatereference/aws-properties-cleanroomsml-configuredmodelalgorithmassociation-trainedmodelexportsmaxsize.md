This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelExportsMaxSize

The maximum size of the trained model metrics that can be exported. If the trained model
metrics dataset is larger than this value, it will not be exported.

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

The unit of measurement for the data size.

_Required_: Yes

_Type_: String

_Allowed values_: `GB`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The maximum size of the dataset to export.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainedModelExportsConfigurationPolicy

TrainedModelInferenceJobsConfigurationPolicy

All content copied from https://docs.aws.amazon.com/.
