This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset Input

Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or
Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseInputDefinition" : DatabaseInputDefinition,
  "DataCatalogInputDefinition" : DataCatalogInputDefinition,
  "Metadata" : Metadata,
  "S3InputDefinition" : S3Location
}

```

### YAML

```yaml

  DatabaseInputDefinition:
    DatabaseInputDefinition
  DataCatalogInputDefinition:
    DataCatalogInputDefinition
  Metadata:
    Metadata
  S3InputDefinition:
    S3Location

```

## Properties

`DatabaseInputDefinition`

Connection information for dataset input files stored in a database.

_Required_: No

_Type_: [DatabaseInputDefinition](aws-properties-databrew-dataset-databaseinputdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataCatalogInputDefinition`

The AWS Glue Data Catalog parameters for the data.

_Required_: No

_Type_: [DataCatalogInputDefinition](aws-properties-databrew-dataset-datacataloginputdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

Contains additional resource information needed for specific datasets.

_Required_: No

_Type_: [Metadata](aws-properties-databrew-dataset-metadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3InputDefinition`

The Amazon S3 location where the data is stored.

_Required_: No

_Type_: [S3Location](aws-properties-databrew-dataset-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormatOptions

JsonOptions

All content copied from https://docs.aws.amazon.com/.
