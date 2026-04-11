This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream RedshiftRetryOptions

Configures retry behavior in case Firehose is unable to deliver
documents to Amazon Redshift.

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

The length of time during which Firehose retries delivery after a
failure, starting from the initial request and including the first attempt. The default
value is 3600 seconds (60 minutes). Firehose does not retry if the value of
`DurationInSeconds` is 0 (zero) or if the first delivery attempt takes longer
than the current value.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftDestinationConfiguration

RetryOptions

All content copied from https://docs.aws.amazon.com/.
