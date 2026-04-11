This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application GlueDataCatalogConfiguration

The configuration of the Glue Data Catalog that you use for Apache Flink SQL queries
and table API transforms that you write in an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseARN" : String
}

```

### YAML

```yaml

  DatabaseARN: String

```

## Properties

`DatabaseARN`

The Amazon Resource Name (ARN) of the database.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlinkRunConfiguration

Input

All content copied from https://docs.aws.amazon.com/.
