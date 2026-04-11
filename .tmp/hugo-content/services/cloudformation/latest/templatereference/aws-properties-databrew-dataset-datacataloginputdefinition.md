This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset DataCatalogInputDefinition

Represents how metadata stored in the AWS Glue Data Catalog is defined in a DataBrew
dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "DatabaseName" : String,
  "TableName" : String,
  "TempDirectory" : S3Location
}

```

### YAML

```yaml

  CatalogId: String
  DatabaseName: String
  TableName: String
  TempDirectory:
    S3Location

```

## Properties

`CatalogId`

The unique identifier of the AWS account that holds the Data Catalog that stores the
data.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of a database in the Data Catalog.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of a database table in the Data Catalog. This table corresponds to a DataBrew
dataset.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TempDirectory`

An Amazon location that AWS Glue Data Catalog can use as a temporary
directory.

_Required_: No

_Type_: [S3Location](aws-properties-databrew-dataset-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseInputDefinition

DatasetParameter

All content copied from https://docs.aws.amazon.com/.
