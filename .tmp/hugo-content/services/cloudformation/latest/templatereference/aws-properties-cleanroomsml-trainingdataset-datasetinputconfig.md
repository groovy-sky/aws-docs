This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::TrainingDataset DatasetInputConfig

Defines the Glue data source and schema mapping information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSource" : DataSource,
  "Schema" : [ ColumnSchema, ... ]
}

```

### YAML

```yaml

  DataSource:
    DataSource
  Schema:
    - ColumnSchema

```

## Properties

`DataSource`

A DataSource object that specifies the Glue data source for the training data.

_Required_: Yes

_Type_: [DataSource](aws-properties-cleanroomsml-trainingdataset-datasource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schema`

The schema information for the training data.

_Required_: Yes

_Type_: Array of [ColumnSchema](aws-properties-cleanroomsml-trainingdataset-columnschema.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dataset

DataSource

All content copied from https://docs.aws.amazon.com/.
