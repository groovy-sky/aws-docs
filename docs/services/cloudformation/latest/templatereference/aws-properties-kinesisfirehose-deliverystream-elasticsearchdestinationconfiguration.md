This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ElasticsearchDestinationConfiguration

The `ElasticsearchDestinationConfiguration` property type specifies an Amazon
Elasticsearch Service (Amazon ES) domain that Amazon Kinesis Data Firehose (Kinesis Data
Firehose) delivers data to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferingHints" : ElasticsearchBufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "ClusterEndpoint" : String,
  "DocumentIdOptions" : DocumentIdOptions,
  "DomainARN" : String,
  "IndexName" : String,
  "IndexRotationPeriod" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : ElasticsearchRetryOptions,
  "RoleARN" : String,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "TypeName" : String,
  "VpcConfiguration" : VpcConfiguration
}

```

### YAML

```yaml

  BufferingHints:
    ElasticsearchBufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  ClusterEndpoint: String
  DocumentIdOptions:
    DocumentIdOptions
  DomainARN: String
  IndexName: String
  IndexRotationPeriod: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    ElasticsearchRetryOptions
  RoleARN: String
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  TypeName: String
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`BufferingHints`

Configures how Kinesis Data Firehose buffers incoming data while delivering it to the
Amazon ES domain.

_Required_: No

_Type_: [ElasticsearchBufferingHints](aws-properties-kinesisfirehose-deliverystream-elasticsearchbufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

The Amazon CloudWatch Logs logging options for the delivery stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterEndpoint`

The endpoint to use when communicating with the cluster. Specify either this
`ClusterEndpoint` or the `DomainARN` field.

_Required_: No

_Type_: String

_Pattern_: `https:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentIdOptions`

Indicates the method for setting up document ID. The supported methods are Firehose generated document ID and OpenSearch Service generated document ID.

_Required_: No

_Type_: [DocumentIdOptions](aws-properties-kinesisfirehose-deliverystream-documentidoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainARN`

The ARN of the Amazon ES domain. The IAM role must have permissions for
`DescribeElasticsearchDomain`, `DescribeElasticsearchDomains`, and
`DescribeElasticsearchDomainConfig` after assuming the role specified in
**RoleARN**.

Specify either `ClusterEndpoint` or `DomainARN`.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The name of the Elasticsearch index to which Kinesis Data Firehose adds data for
indexing.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexRotationPeriod`

The frequency of Elasticsearch index rotation. If you enable index rotation, Kinesis
Data Firehose appends a portion of the UTC arrival timestamp to the specified index name,
and rotates the appended timestamp accordingly. For more information, see [Index Rotation for the Amazon ES Destination](../../../firehose/latest/dev/basic-deliver.md#es-index-rotation) in the _Amazon Kinesis_
_Data Firehose Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `NoRotation | OneHour | OneDay | OneWeek | OneMonth`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

The data processing configuration for the Kinesis Data Firehose delivery
stream.

_Required_: No

_Type_: [ProcessingConfiguration](aws-properties-kinesisfirehose-deliverystream-processingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The retry behavior when Kinesis Data Firehose is unable to deliver data to Amazon
ES.

_Required_: No

_Type_: [ElasticsearchRetryOptions](aws-properties-kinesisfirehose-deliverystream-elasticsearchretryoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose
for calling the Amazon ES Configuration API and for indexing documents. For more
information, see [Controlling Access with Amazon Kinesis\
Data Firehose](../../../firehose/latest/dev/controlling-access.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

The condition under which Kinesis Data Firehose delivers data to Amazon Simple Storage
Service (Amazon S3). You can send Amazon S3 all documents (all data) or only the documents
that Kinesis Data Firehose could not deliver to the Amazon ES destination. For more
information and valid values, see the `S3BackupMode` content for the [ElasticsearchDestinationConfiguration](../../../../reference/firehose/latest/apireference/api-elasticsearchdestinationconfiguration.md) data type in the _Amazon Kinesis_
_Data Firehose API Reference_.

_Required_: No

_Type_: String

_Allowed values_: `FailedDocumentsOnly | AllDocuments`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

The S3 bucket where Kinesis Data Firehose backs up incoming data.

_Required_: Yes

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The Elasticsearch type name that Amazon ES adds to documents when indexing
data.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The details of the VPC of the Amazon ES destination.

_Required_: No

_Type_: [VpcConfiguration](aws-properties-kinesisfirehose-deliverystream-vpcconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticsearchBufferingHints

ElasticsearchRetryOptions

All content copied from https://docs.aws.amazon.com/.
