This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream HttpEndpointDestinationConfiguration

Describes the configuration of the HTTP endpoint destination. Kinesis Firehose
supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service
providers, including Datadog, MongoDB, and New Relic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferingHints" : BufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "EndpointConfiguration" : HttpEndpointConfiguration,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RequestConfiguration" : HttpEndpointRequestConfiguration,
  "RetryOptions" : RetryOptions,
  "RoleARN" : String,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "SecretsManagerConfiguration" : SecretsManagerConfiguration
}

```

### YAML

```yaml

  BufferingHints:
    BufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  EndpointConfiguration:
    HttpEndpointConfiguration
  ProcessingConfiguration:
    ProcessingConfiguration
  RequestConfiguration:
    HttpEndpointRequestConfiguration
  RetryOptions:
    RetryOptions
  RoleARN: String
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  SecretsManagerConfiguration:
    SecretsManagerConfiguration

```

## Properties

`BufferingHints`

The buffering options that can be used before data is delivered to the specified
destination. Kinesis Data Firehose treats these options as hints, and it might choose to
use more optimal values. The SizeInMBs and IntervalInSeconds parameters are optional.
However, if you specify a value for one of them, you must also provide a value for the
other.

_Required_: No

_Type_: [BufferingHints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-bufferinghints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

Describes the Amazon CloudWatch logging options for your delivery stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfiguration`

The configuration of the HTTP endpoint selected as the destination.

_Required_: Yes

_Type_: [HttpEndpointConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Describes the data processing configuration.

_Required_: No

_Type_: [ProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestConfiguration`

The configuration of the request sent to the HTTP endpoint specified as the
destination.

_Required_: No

_Type_: [HttpEndpointRequestConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-httpendpointrequestconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data
to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment
of receipt from the specified HTTP endpoint destination.

_Required_: No

_Type_: [RetryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-retryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

Kinesis Data Firehose uses this IAM role for all the permissions that the delivery
stream needs.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

Describes the S3 bucket backup options for the data that Kinesis Data Firehose
delivers to the HTTP endpoint destination. You can back up all documents (AllData) or only
the documents that Kinesis Data Firehose could not deliver to the specified HTTP endpoint
destination (FailedDataOnly).

_Required_: No

_Type_: String

_Allowed values_: `FailedDataOnly | AllData`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Describes the configuration of a destination in Amazon S3.

_Required_: Yes

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerConfiguration`

The configuration that defines how you access secrets for HTTP Endpoint destination.

_Required_: No

_Type_: [SecretsManagerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HttpEndpointConfiguration

HttpEndpointRequestConfiguration
