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

_Type_: [BufferingHints](aws-properties-kinesisfirehose-deliverystream-bufferinghints.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

Describes the Amazon CloudWatch logging options for your delivery stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfiguration`

The configuration of the HTTP endpoint selected as the destination.

_Required_: Yes

_Type_: [HttpEndpointConfiguration](aws-properties-kinesisfirehose-deliverystream-httpendpointconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

Describes the data processing configuration.

_Required_: No

_Type_: [ProcessingConfiguration](aws-properties-kinesisfirehose-deliverystream-processingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestConfiguration`

The configuration of the request sent to the HTTP endpoint specified as the
destination.

_Required_: No

_Type_: [HttpEndpointRequestConfiguration](aws-properties-kinesisfirehose-deliverystream-httpendpointrequestconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data
to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment
of receipt from the specified HTTP endpoint destination.

_Required_: No

_Type_: [RetryOptions](aws-properties-kinesisfirehose-deliverystream-retryoptions.md)

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

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerConfiguration`

The configuration that defines how you access secrets for HTTP Endpoint destination.

_Required_: No

_Type_: [SecretsManagerConfiguration](aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpEndpointConfiguration

HttpEndpointRequestConfiguration

All content copied from https://docs.aws.amazon.com/.
