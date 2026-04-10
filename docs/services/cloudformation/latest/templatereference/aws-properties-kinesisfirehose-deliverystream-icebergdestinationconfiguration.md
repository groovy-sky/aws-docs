This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream IcebergDestinationConfiguration

Specifies the destination configure settings for Apache Iceberg Table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppendOnly" : Boolean,
  "BufferingHints" : BufferingHints,
  "CatalogConfiguration" : CatalogConfiguration,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "DestinationTableConfigurationList" : [ DestinationTableConfiguration, ... ],
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : RetryOptions,
  "RoleARN" : String,
  "s3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "SchemaEvolutionConfiguration" : SchemaEvolutionConfiguration,
  "TableCreationConfiguration" : TableCreationConfiguration
}

```

### YAML

```yaml

  AppendOnly: Boolean
  BufferingHints:
    BufferingHints
  CatalogConfiguration:
    CatalogConfiguration
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  DestinationTableConfigurationList:
    - DestinationTableConfiguration
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    RetryOptions
  RoleARN: String
  s3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  SchemaEvolutionConfiguration:
    SchemaEvolutionConfiguration
  TableCreationConfiguration:
    TableCreationConfiguration

```

## Properties

`AppendOnly`

Describes whether all incoming data for this delivery stream will be append only
(inserts only and not for updates and deletes) for Iceberg delivery. This feature is only
applicable for Apache Iceberg Tables.

The default value is false. If you set this value to true, Firehose automatically
increases the throughput limit of a stream based on the throttling levels of the stream. If
you set this parameter to true for a stream with updates and deletes, you will see out of
order delivery.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferingHints`

Property description not available.

_Required_: No

_Type_: [BufferingHints](aws-properties-kinesisfirehose-deliverystream-bufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CatalogConfiguration`

Configuration describing where the destination Apache Iceberg Tables are persisted.

_Required_: Yes

_Type_: [CatalogConfiguration](aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CloudWatchLoggingOptions`

Property description not available.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationTableConfigurationList`

Provides a list of `DestinationTableConfigurations` which Firehose uses
to deliver data to Apache Iceberg Tables. Firehose will write data with insert if table specific configuration is not provided here.

_Required_: No

_Type_: Array of [DestinationTableConfiguration](aws-properties-kinesisfirehose-deliverystream-destinationtableconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Property description not available.

_Required_: No

_Type_: [ProcessingConfiguration](aws-properties-kinesisfirehose-deliverystream-processingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

Property description not available.

_Required_: No

_Type_: [RetryOptions](aws-properties-kinesisfirehose-deliverystream-retryoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the IAM role to be assumed by Firehose for calling Apache Iceberg Tables.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`s3BackupMode`

Describes how Firehose will backup records. Currently,S3 backup only supports
`FailedDataOnly`.

_Required_: No

_Type_: String

_Allowed values_: `AllData | FailedDataOnly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Property description not available.

_Required_: Yes

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaEvolutionConfiguration`

The configuration to enable automatic schema evolution.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: [SchemaEvolutionConfiguration](aws-properties-kinesisfirehose-deliverystream-schemaevolutionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableCreationConfiguration`

The configuration to enable automatic table creation.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: [TableCreationConfiguration](aws-properties-kinesisfirehose-deliverystream-tablecreationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpEndpointRequestConfiguration

InputFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
