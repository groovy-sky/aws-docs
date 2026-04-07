This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::TrainingDataset Dataset

Defines where the training dataset is located, what type of data it contains, and how to
access the data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputConfig" : DatasetInputConfig,
  "Type" : String
}

```

### YAML

```yaml

  InputConfig:
    DatasetInputConfig
  Type: String

```

## Properties

`InputConfig`

A DatasetInputConfig object that defines the data source and schema mapping.

_Required_: Yes

_Type_: [DatasetInputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-trainingdataset-datasetinputconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

What type of information is found in the dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `INTERACTIONS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ColumnSchema

DatasetInputConfig
