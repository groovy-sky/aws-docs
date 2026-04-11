This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign S3Config

The Amazon S3 bucket where the AWS IoT FleetWise campaign sends data. Amazon S3 is an object storage service that stores data as objects within buckets. For more information, see [Creating, configuring, and working with Amazon S3 buckets](../../../s3/latest/userguide/creating-buckets-s3.md) in the _Amazon Simple Storage Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketArn" : String,
  "DataFormat" : String,
  "Prefix" : String,
  "StorageCompressionFormat" : String
}

```

### YAML

```yaml

  BucketArn: String
  DataFormat: String
  Prefix: String
  StorageCompressionFormat: String

```

## Properties

`BucketArn`

The Amazon Resource Name (ARN) of the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):s3:::.+$`

_Minimum_: `16`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataFormat`

Specify the format that files are saved in the Amazon S3 bucket. You can save files in an Apache Parquet or JSON format.

- Parquet - Store data in a columnar storage file format. Parquet is optimal for
fast data retrieval and can reduce costs. This option is selected by
default.

- JSON - Store data in a standard text-based JSON file format.

_Required_: No

_Type_: String

_Allowed values_: `JSON | PARQUET`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Prefix`

Enter an S3 bucket prefix. The prefix is the string of characters after the bucket name and before the object name. You can use the prefix to organize data stored in Amazon S3 buckets. For more information, see [Organizing objects using prefixes](../../../s3/latest/userguide/using-prefixes.md) in the _Amazon Simple Storage Service User Guide_.

By default, AWS IoT FleetWise sets the prefix `processed-data/year=YY/month=MM/date=DD/hour=HH/` (in UTC) to data it delivers to Amazon S3. You can enter a prefix to append it to this default prefix. For example, if you enter the prefix `vehicles`, the prefix will be `vehicles/processed-data/year=YY/month=MM/date=DD/hour=HH/`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_:./!*'()]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageCompressionFormat`

By default, stored data is compressed as a .gzip file. Compressed files have a reduced
file size, which can optimize the cost of data storage.

_Required_: No

_Type_: String

_Allowed values_: `NONE | GZIP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MqttTopicConfig

SignalFetchConfig

All content copied from https://docs.aws.amazon.com/.
