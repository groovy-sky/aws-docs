This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset GlueConfiguration

Configuration information for coordination with AWS Glue, a fully managed extract, transform
and load (ETL) service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "TableName" : String
}

```

### YAML

```yaml

  DatabaseName: String
  TableName: String

```

## Properties

`DatabaseName`

The name of the database in your AWS Glue Data Catalog in which the table is located. An
AWS Glue Data Catalog database contains metadata tables.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the table in your AWS Glue Data Catalog that is used to perform the ETL
operations. An AWS Glue Data Catalog table contains partitioned data and descriptions of data
sources and targets.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

IotEventsDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
