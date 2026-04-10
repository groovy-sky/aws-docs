This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream RetryOptions

Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data
to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment
of receipt from the specified HTTP endpoint destination. Kinesis Firehose supports any
custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers,
including Datadog, MongoDB, and New Relic.

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

The total amount of time that Kinesis Data Firehose spends on retries. This duration
starts after the initial attempt to send data to the custom destination via HTTPS endpoint
fails. It doesn't include the periods during which Kinesis Data Firehose waits for
acknowledgment from the specified destination after each attempt.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftRetryOptions

S3DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
