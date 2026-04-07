This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream AmazonOpenSearchServerlessDestinationConfiguration

Describes the configuration of a destination in the Serverless offering for Amazon
OpenSearch Service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferingHints" : AmazonOpenSearchServerlessBufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "CollectionEndpoint" : String,
  "IndexName" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : AmazonOpenSearchServerlessRetryOptions,
  "RoleARN" : String,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "VpcConfiguration" : VpcConfiguration
}

```

### YAML

```yaml

  BufferingHints:
    AmazonOpenSearchServerlessBufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  CollectionEndpoint: String
  IndexName: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    AmazonOpenSearchServerlessRetryOptions
  RoleARN: String
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`BufferingHints`

The buffering options. If no value is specified, the default values for
AmazonopensearchserviceBufferingHints are used.

_Required_: No

_Type_: [AmazonOpenSearchServerlessBufferingHints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessbufferinghints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

Property description not available.

_Required_: No

_Type_: [CloudWatchLoggingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollectionEndpoint`

The endpoint to use when communicating with the collection in the Serverless offering
for Amazon OpenSearch Service.

_Required_: No

_Type_: String

_Pattern_: `https:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The Serverless offering for Amazon OpenSearch Service index name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Property description not available.

_Required_: No

_Type_: [ProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The retry behavior in case Firehose is unable to deliver documents to the
Serverless offering for Amazon OpenSearch Service. The default value is 300 (5
minutes).

_Required_: No

_Type_: [AmazonOpenSearchServerlessRetryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessretryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the IAM role to be assumed by Firehose
for calling the Serverless offering for Amazon OpenSearch Service Configuration API and for
indexing documents.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

Defines how documents should be delivered to Amazon S3. When it is set to
FailedDocumentsOnly, Firehose writes any documents that could not be indexed
to the configured Amazon S3 destination, with AmazonOpenSearchService-failed/ appended to
the key prefix. When set to AllDocuments, Firehose delivers all incoming
records to Amazon S3, and also writes failed documents with AmazonOpenSearchService-failed/
appended to the prefix.

_Required_: No

_Type_: String

_Allowed values_: `FailedDocumentsOnly | AllDocuments`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Property description not available.

_Required_: Yes

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

Property description not available.

_Required_: No

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AmazonOpenSearchServerlessBufferingHints

AmazonOpenSearchServerlessRetryOptions
