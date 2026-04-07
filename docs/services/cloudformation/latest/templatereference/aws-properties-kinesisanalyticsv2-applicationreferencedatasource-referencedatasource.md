This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource ReferenceDataSource

For a SQL-based Kinesis Data Analytics application, describes the reference data
source by providing the source information (Amazon S3 bucket name and object key name), the
resulting in-application table name that is created, and the necessary schema to map the data
elements in the Amazon S3 object to the in-application table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReferenceSchema" : ReferenceSchema,
  "S3ReferenceDataSource" : S3ReferenceDataSource,
  "TableName" : String
}

```

### YAML

```yaml

  ReferenceSchema:
    ReferenceSchema
  S3ReferenceDataSource:
    S3ReferenceDataSource
  TableName: String

```

## Properties

`ReferenceSchema`

Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.

_Required_: Yes

_Type_: [ReferenceSchema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referenceschema.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ReferenceDataSource`

Identifies the S3 bucket and object that contains the reference data. A Kinesis Data
Analytics application loads reference data only once. If the data changes, you call the
[UpdateApplication](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your
application.

_Required_: No

_Type_: [S3ReferenceDataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the in-application table to create.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ReferenceDataSource](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_ReferenceDataSource.html) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecordFormat

ReferenceSchema
