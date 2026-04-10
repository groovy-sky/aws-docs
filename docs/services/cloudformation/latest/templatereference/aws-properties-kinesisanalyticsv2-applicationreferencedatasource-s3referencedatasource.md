This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource S3ReferenceDataSource

For an SQL-based Amazon Kinesis Data Analytics application, identifies the Amazon S3
bucket and object that contains the reference data.

A Kinesis Data Analytics application loads reference data only once. If the data
changes, you call the [UpdateApplication](../../../managed-flink/latest/apiv2/api-updateapplication.md) operation to trigger reloading of data into your
application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketARN" : String,
  "FileKey" : String
}

```

### YAML

```yaml

  BucketARN: String
  FileKey: String

```

## Properties

`BucketARN`

The Amazon Resource Name (ARN) of the S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileKey`

The object key name containing the reference data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3ReferenceDataSource](../../../managed-flink/latest/apiv2/api-s3referencedatasource.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceSchema

Next

All content copied from https://docs.aws.amazon.com/.
