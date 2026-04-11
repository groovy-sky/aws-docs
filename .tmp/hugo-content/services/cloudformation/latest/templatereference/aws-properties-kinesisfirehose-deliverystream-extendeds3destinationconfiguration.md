This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ExtendedS3DestinationConfiguration

The `ExtendedS3DestinationConfiguration` property type configures an
Amazon S3 destination for an Amazon Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketARN" : String,
  "BufferingHints" : BufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "CompressionFormat" : String,
  "CustomTimeZone" : String,
  "DataFormatConversionConfiguration" : DataFormatConversionConfiguration,
  "DynamicPartitioningConfiguration" : DynamicPartitioningConfiguration,
  "EncryptionConfiguration" : EncryptionConfiguration,
  "ErrorOutputPrefix" : String,
  "FileExtension" : String,
  "Prefix" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RoleARN" : String,
  "S3BackupConfiguration" : S3DestinationConfiguration,
  "S3BackupMode" : String
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
  CustomTimeZone: String
  DataFormatConversionConfiguration:
    DataFormatConversionConfiguration
  DynamicPartitioningConfiguration:
    DynamicPartitioningConfiguration
  EncryptionConfiguration:
    EncryptionConfiguration
  ErrorOutputPrefix: String
  FileExtension: String
  Prefix: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RoleARN: String
  S3BackupConfiguration:
    S3DestinationConfiguration
  S3BackupMode: String

```

## Properties

`BucketARN`

The Amazon Resource Name (ARN) of the Amazon S3 bucket. For constraints, see [ExtendedS3DestinationConfiguration](../../../../reference/firehose/latest/apireference/api-extendeds3destinationconfiguration.md) in the _Amazon Kinesis Data_
_Firehose API Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferingHints`

The buffering option.

_Required_: No

_Type_: [BufferingHints](aws-properties-kinesisfirehose-deliverystream-bufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

The Amazon CloudWatch logging options for your Firehose stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompressionFormat`

The compression format. If no value is specified, the default is
`UNCOMPRESSED`.

_Required_: No

_Type_: String

_Allowed values_: `UNCOMPRESSED | GZIP | ZIP | Snappy | HADOOP_SNAPPY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomTimeZone`

The time zone you prefer. UTC is the default.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataFormatConversionConfiguration`

The serializer, deserializer, and schema for converting data from the JSON format to
the Parquet or ORC format before writing it to Amazon S3.

_Required_: No

_Type_: [DataFormatConversionConfiguration](aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamicPartitioningConfiguration`

The configuration of the dynamic partitioning mechanism that creates targeted data
sets from the streaming data by partitioning it based on partition keys.

_Required_: No

_Type_: [DynamicPartitioningConfiguration](aws-properties-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

The encryption configuration for the Kinesis Data Firehose delivery stream. The default
value is `NoEncryption`.

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

`FileExtension`

Specify a file extension. It will override the default file extension

_Required_: No

_Type_: String

_Pattern_: `^$|\.[0-9a-z!\-_.*'()]+`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The `YYYY/MM/DD/HH` time format prefix is automatically used for delivered
Amazon S3 files. For more information, see [ExtendedS3DestinationConfiguration](../../../../reference/firehose/latest/apireference/api-extendeds3destinationconfiguration.md) in the _Amazon Kinesis Data_
_Firehose API Reference_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

The data processing configuration for the Kinesis Data Firehose delivery
stream.

_Required_: No

_Type_: [ProcessingConfiguration](aws-properties-kinesisfirehose-deliverystream-processingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the AWS credentials. For constraints,
see [ExtendedS3DestinationConfiguration](../../../../reference/firehose/latest/apireference/api-extendeds3destinationconfiguration.md) in the _Amazon Kinesis Data_
_Firehose API Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupConfiguration`

The configuration for backup in Amazon S3.

_Required_: No

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

The Amazon S3 backup mode. After you create a Firehose stream, you can update it to
enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the
Firehose stream to disable it.

_Required_: No

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

HiveJsonSerDe

All content copied from https://docs.aws.amazon.com/.
