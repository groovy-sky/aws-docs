This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SchemaConfiguration

Specifies the schema to which you want Firehose to configure your data
before it writes it to Amazon S3. This parameter is required if `Enabled` is set
to true.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "DatabaseName" : String,
  "Region" : String,
  "RoleARN" : String,
  "TableName" : String,
  "VersionId" : String
}

```

### YAML

```yaml

  CatalogId: String
  DatabaseName: String
  Region: String
  RoleARN: String
  TableName: String
  VersionId: String

```

## Properties

`CatalogId`

The ID of the AWS Glue Data Catalog. If you don't supply this, the
AWS account ID is used by default.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

Specifies the name of the AWS Glue database that contains the schema for
the output data.

###### Important

If the `SchemaConfiguration` request parameter is used as part of invoking
the `CreateDeliveryStream` API, then the `DatabaseName` property
is required and its value must be specified.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

If you don't specify an AWS Region, the default is the current
Region.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The role that Firehose can use to access AWS Glue. This
role must be in the same account you use for Firehose. Cross-account roles
aren't allowed.

###### Important

If the `SchemaConfiguration` request parameter is used as part of invoking
the `CreateDeliveryStream` API, then the `RoleARN` property is
required and its value must be specified.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

Specifies the AWS Glue table that contains the column information that
constitutes your data schema.

###### Important

If the `SchemaConfiguration` request parameter is used as part of invoking
the `CreateDeliveryStream` API, then the `TableName` property is
required and its value must be specified.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

Specifies the table version for the output data schema. If you don't specify this
version ID, or if you set it to `LATEST`, Firehose uses the most
recent version. This means that any updates to the table are automatically picked
up.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3DestinationConfiguration

SchemaEvolutionConfiguration

All content copied from https://docs.aws.amazon.com/.
