This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream AmazonopensearchserviceDestinationConfiguration

Describes the configuration of a destination in Amazon OpenSearch Service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferingHints" : AmazonopensearchserviceBufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "ClusterEndpoint" : String,
  "DocumentIdOptions" : DocumentIdOptions,
  "DomainARN" : String,
  "IndexName" : String,
  "IndexRotationPeriod" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : AmazonopensearchserviceRetryOptions,
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
    AmazonopensearchserviceBufferingHints
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
    AmazonopensearchserviceRetryOptions
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

The buffering options. If no value is specified, the default values for
AmazonopensearchserviceBufferingHints are used.

_Required_: No

_Type_: [AmazonopensearchserviceBufferingHints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-amazonopensearchservicebufferinghints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

Describes the Amazon CloudWatch logging options for your delivery stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterEndpoint`

The endpoint to use when communicating with the cluster. Specify either this
ClusterEndpoint or the DomainARN field.

_Required_: No

_Type_: String

_Pattern_: `https:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentIdOptions`

Indicates the method for setting up document ID. The supported methods are Firehose generated document ID and OpenSearch Service generated document ID.

_Required_: No

_Type_: [DocumentIdOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-documentidoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainARN`

The ARN of the Amazon OpenSearch Service domain.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The Amazon OpenSearch Service index name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexRotationPeriod`

The Amazon OpenSearch Service index rotation period. Index rotation appends a timestamp
to the IndexName to facilitate the expiration of old data.

_Required_: No

_Type_: String

_Allowed values_: `NoRotation | OneHour | OneDay | OneWeek | OneMonth`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Describes a data processing configuration.

_Required_: No

_Type_: [ProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The retry behavior in case Kinesis Data Firehose is unable to deliver documents to
Amazon OpenSearch Service. The default value is 300 (5 minutes).

_Required_: No

_Type_: [AmazonopensearchserviceRetryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserviceretryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose
for calling the Amazon OpenSearch Service Configuration API and for indexing
documents.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

Defines how documents should be delivered to Amazon S3.

_Required_: No

_Type_: String

_Allowed values_: `FailedDocumentsOnly | AllDocuments`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Describes the configuration of a destination in Amazon S3.

_Required_: Yes

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The Amazon OpenSearch Service type name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The details of the VPC of the Amazon OpenSearch Service destination.

_Required_: No

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AmazonopensearchserviceBufferingHints

AmazonopensearchserviceRetryOptions
