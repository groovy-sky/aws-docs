This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ElasticsearchRetryOptions

The `ElasticsearchRetryOptions` property type configures the retry behavior
for when Amazon Kinesis Data Firehose (Kinesis Data Firehose) can't deliver data to Amazon
Elasticsearch Service (Amazon ES).

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

After an initial failure to deliver to Amazon ES, the total amount of time during which
Kinesis Data Firehose re-attempts delivery (including the first attempt). If Kinesis Data
Firehose can't deliver the data within the specified time, it writes the data to the backup
S3 bucket. For valid values, see the `DurationInSeconds` content for the [ElasticsearchRetryOptions](../../../../reference/firehose/latest/apireference/api-elasticsearchretryoptions.md) data type in the _Amazon Kinesis Data_
_Firehose API Reference_.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticsearchDestinationConfiguration

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
