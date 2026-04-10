This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelQualityJobDefinition DatasetFormat

The format of the dataset used for the model quality monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Csv" : Csv,
  "Json" : Json,
  "Parquet" : Boolean
}

```

### YAML

```yaml

  Csv:
    Csv
  Json:
    Json
  Parquet: Boolean

```

## Properties

`Csv`

The CSV format configuration for the dataset.

_Required_: No

_Type_: [Csv](aws-properties-sagemaker-modelqualityjobdefinition-csv.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Json`

The JSON format configuration for the dataset.

_Required_: No

_Type_: [Json](aws-properties-sagemaker-modelqualityjobdefinition-json.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parquet`

Indicates that the dataset is in Parquet format.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Csv

EndpointInput

All content copied from https://docs.aws.amazon.com/.
