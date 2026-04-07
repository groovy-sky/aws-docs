This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SplunkDestinationConfiguration

The `SplunkDestinationConfiguration` property type specifies the
configuration of a destination in Splunk for a Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferingHints" : SplunkBufferingHints,
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "HECAcknowledgmentTimeoutInSeconds" : Integer,
  "HECEndpoint" : String,
  "HECEndpointType" : String,
  "HECToken" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : SplunkRetryOptions,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "SecretsManagerConfiguration" : SecretsManagerConfiguration
}

```

### YAML

```yaml

  BufferingHints:
    SplunkBufferingHints
  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  HECAcknowledgmentTimeoutInSeconds: Integer
  HECEndpoint: String
  HECEndpointType: String
  HECToken: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    SplunkRetryOptions
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  SecretsManagerConfiguration:
    SecretsManagerConfiguration

```

## Properties

`BufferingHints`

The buffering options. If no value is specified, the default values for Splunk are used.

_Required_: No

_Type_: [SplunkBufferingHints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-splunkbufferinghints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLoggingOptions`

The Amazon CloudWatch logging options for your Firehose stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HECAcknowledgmentTimeoutInSeconds`

The amount of time that Firehose waits to receive an acknowledgment from
Splunk after it sends it data. At the end of the timeout period, Firehose
either tries to send the data again or considers it an error, based on your retry
settings.

_Required_: No

_Type_: Integer

_Minimum_: `180`

_Maximum_: `600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HECEndpoint`

The HTTP Event Collector (HEC) endpoint to which Firehose sends your
data.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HECEndpointType`

This type can be either `Raw` or `Event`.

_Required_: Yes

_Type_: String

_Allowed values_: `Raw | Event`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HECToken`

This is a GUID that you obtain from your Splunk cluster when you create a new HEC
endpoint.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

The data processing configuration.

_Required_: No

_Type_: [ProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The retry behavior in case Firehose is unable to deliver data to Splunk,
or if it doesn't receive an acknowledgment of receipt from Splunk.

_Required_: No

_Type_: [SplunkRetryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-splunkretryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

Defines how documents should be delivered to Amazon S3. When set to
`FailedEventsOnly`, Firehose writes any data that could not be
indexed to the configured Amazon S3 destination. When set to `AllEvents`,
Firehose delivers all incoming records to Amazon S3, and also writes failed
documents to Amazon S3. The default value is `FailedEventsOnly`.

You can update this backup mode from `FailedEventsOnly` to
`AllEvents`. You can't update it from `AllEvents` to
`FailedEventsOnly`.

_Required_: No

_Type_: String

_Allowed values_: `FailedEventsOnly | AllEvents`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

The configuration for the backup Amazon S3 location.

_Required_: Yes

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerConfiguration`

The configuration that defines how you access secrets for Splunk.

_Required_: No

_Type_: [SecretsManagerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SplunkDestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_SplunkDestinationConfiguration.html) in the _Amazon Kinesis Data_
_Firehose API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SplunkBufferingHints

SplunkRetryOptions
