This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream CatalogConfiguration

Describes the containers where the destination Apache Iceberg Tables are persisted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogArn" : String,
  "WarehouseLocation" : String
}

```

### YAML

```yaml

  CatalogArn: String
  WarehouseLocation: String

```

## Properties

`CatalogArn`

Specifies the Glue catalog ARN identifier of the destination Apache Iceberg Tables. You must specify the ARN in the format `arn:aws:glue:region:account-id:catalog`.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WarehouseLocation`

The warehouse location for Apache Iceberg tables. You must configure this when schema
evolution and table creation is enabled.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: String

_Pattern_: `s3:\/\/.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BufferingHints

CloudWatchLoggingOptions

All content copied from https://docs.aws.amazon.com/.
