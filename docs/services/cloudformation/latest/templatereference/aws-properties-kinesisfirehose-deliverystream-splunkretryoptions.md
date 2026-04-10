This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SplunkRetryOptions

The `SplunkRetryOptions` property type specifies retry behavior in case
Kinesis Data Firehose is unable to deliver documents to Splunk or if it doesn't receive an
acknowledgment from Splunk.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DurationInSeconds" : Integer
}

```

### YAML

```yaml

  DurationInSeconds: Integer

```

## Properties

`DurationInSeconds`

The total amount of time that Firehose spends on retries. This duration
starts after the initial attempt to send data to Splunk fails. It doesn't include the
periods during which Firehose waits for acknowledgment from Splunk after each
attempt.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SplunkRetryOptions](../../../../reference/firehose/latest/apireference/api-splunkretryoptions.md) in the _Amazon Kinesis Data Firehose API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SplunkDestinationConfiguration

TableCreationConfiguration

All content copied from https://docs.aws.amazon.com/.
