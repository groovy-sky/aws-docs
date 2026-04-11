This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset DatabaseInputDefinition

Connection information for dataset input files stored in a database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseTableName" : String,
  "GlueConnectionName" : String,
  "QueryString" : String,
  "TempDirectory" : S3Location
}

```

### YAML

```yaml

  DatabaseTableName: String
  GlueConnectionName: String
  QueryString:
    String
  TempDirectory:
    S3Location

```

## Properties

`DatabaseTableName`

The table within the target database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueConnectionName`

The AWS Glue Connection that stores the connection information for
the target database.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Custom SQL to run against the provided AWS Glue connection. This SQL will be used as
the input for DataBrew projects and jobs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TempDirectory`

An Amazon location that AWS Glue Data Catalog can use as a temporary
directory.

_Required_: No

_Type_: [S3Location](aws-properties-databrew-dataset-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CsvOptions

DataCatalogInputDefinition

All content copied from https://docs.aws.amazon.com/.
