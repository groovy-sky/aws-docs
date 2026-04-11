This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe Input

Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or
Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataCatalogInputDefinition" : DataCatalogInputDefinition,
  "S3InputDefinition" : S3Location
}

```

### YAML

```yaml

  DataCatalogInputDefinition:
    DataCatalogInputDefinition
  S3InputDefinition:
    S3Location

```

## Properties

`DataCatalogInputDefinition`

The AWS Glue Data Catalog parameters for the data.

_Required_: No

_Type_: [DataCatalogInputDefinition](aws-properties-databrew-recipe-datacataloginputdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3InputDefinition`

The Amazon S3 location where the data is stored.

_Required_: No

_Type_: [S3Location](aws-properties-databrew-recipe-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataCatalogInputDefinition

Parameters

All content copied from https://docs.aws.amazon.com/.
