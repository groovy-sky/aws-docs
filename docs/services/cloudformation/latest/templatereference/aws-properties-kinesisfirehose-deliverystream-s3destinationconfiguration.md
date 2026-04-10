This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream S3DestinationConfiguration

The `S3DestinationConfiguration` property type specifies an Amazon Simple
Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data
Firehose) delivers data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketARN" : String,
  "BufferingHints" : BufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "CompressionFormat" : String,
  "EncryptionConfiguration" : EncryptionConfiguration,
  "ErrorOutputPrefix" : String,
  "Prefix" : String,
  "RoleARN" : String
}

```

### YAML

```yaml

  BucketARN: String
  BufferingHints:
    BufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  CompressionFormat: String
  EncryptionConfiguration:
    EncryptionConfiguration
  ErrorOutputPrefix: String
  Prefix: String
  RoleARN: String

```

## Properties

`BucketARN`

The Amazon Resource Name (ARN) of the Amazon S3 bucket to send data to.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferingHints`

Configures how Kinesis Data Firehose buffers incoming data while delivering it to the
Amazon S3 bucket.

_Required_: No

_Type_: [BufferingHints](aws-properties-kinesisfirehose-deliverystream-bufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

The CloudWatch logging options for your Firehose stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompressionFormat`

The type of compression that Kinesis Data Firehose uses to compress the data that it
delivers to the Amazon S3 bucket. For valid values, see the `CompressionFormat`
content for the [S3DestinationConfiguration](../../../../reference/firehose/latest/apireference/api-s3destinationconfiguration.md) data type in the _Amazon Kinesis Data_
_Firehose API Reference_.

_Required_: No

_Type_: String

_Allowed values_: `UNCOMPRESSED | GZIP | ZIP | Snappy | HADOOP_SNAPPY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

Configures Amazon Simple Storage Service (Amazon S3) server-side encryption. Kinesis
Data Firehose uses AWS Key Management Service (AWS KMS)
to encrypt the data that it delivers to your Amazon S3 bucket.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-kinesisfirehose-deliverystream-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorOutputPrefix`

A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing
them to S3. This prefix appears immediately following the bucket name. For information
about how to specify this prefix, see [Custom Prefixes for Amazon S3\
Objects](../../../firehose/latest/dev/s3-prefixes.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

A prefix that Kinesis Data Firehose adds to the files that it delivers to the Amazon S3
bucket. The prefix helps you identify the files that Kinesis Data Firehose
delivered.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The ARN of an AWS Identity and Access Management (IAM) role that grants
Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS (if you
enable data encryption). For more information, see [Grant Kinesis Data\
Firehose Access to an Amazon S3 Destination](../../../firehose/latest/dev/controlling-access.md#using-iam-s3) in the _Amazon Kinesis Data_
_Firehose Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetryOptions

SchemaConfiguration

All content copied from https://docs.aws.amazon.com/.
