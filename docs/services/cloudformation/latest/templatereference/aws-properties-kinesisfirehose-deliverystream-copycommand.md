This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream CopyCommand

The `CopyCommand` property type configures the Amazon Redshift
`COPY` command that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses
to load data into an Amazon Redshift cluster from an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyOptions" : String,
  "DataTableColumns" : String,
  "DataTableName" : String
}

```

### YAML

```yaml

  CopyOptions: String
  DataTableColumns: String
  DataTableName: String

```

## Properties

`CopyOptions`

Parameters to use with the Amazon Redshift `COPY` command. For examples, see
the `CopyOptions` content for the [CopyCommand](../../../../reference/firehose/latest/apireference/api-copycommand.md) data type in
the _Amazon Kinesis Data Firehose API Reference_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `204800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTableColumns`

A comma-separated list of column names.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `204800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTableName`

The name of the target table. The table must already exist in the database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLoggingOptions

DatabaseColumns

All content copied from https://docs.aws.amazon.com/.
