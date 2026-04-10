This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application CatalogConfiguration

The configuration parameters for the default Amazon Glue database. You use this
database for SQL queries that you write in a Kinesis Data Analytics Studio
notebook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueDataCatalogConfiguration" : GlueDataCatalogConfiguration
}

```

### YAML

```yaml

  GlueDataCatalogConfiguration:
    GlueDataCatalogConfiguration

```

## Properties

`GlueDataCatalogConfiguration`

The configuration parameters for the default Amazon Glue database. You use this
database for Apache Flink SQL queries and table API transforms that you write in a
Kinesis Data Analytics Studio notebook.

_Required_: No

_Type_: [GlueDataCatalogConfiguration](aws-properties-kinesisanalyticsv2-application-gluedatacatalogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationSystemRollbackConfiguration

CheckpointConfiguration

All content copied from https://docs.aws.amazon.com/.
