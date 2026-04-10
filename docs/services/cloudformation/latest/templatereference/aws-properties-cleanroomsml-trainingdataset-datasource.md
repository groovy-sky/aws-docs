This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::TrainingDataset DataSource

Defines information about the Glue data source that contains the training data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueDataSource" : GlueDataSource
}

```

### YAML

```yaml

  GlueDataSource:
    GlueDataSource

```

## Properties

`GlueDataSource`

A GlueDataSource object that defines the catalog ID, database name, and table name for
the training data.

_Required_: Yes

_Type_: [GlueDataSource](aws-properties-cleanroomsml-trainingdataset-gluedatasource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetInputConfig

GlueDataSource

All content copied from https://docs.aws.amazon.com/.
