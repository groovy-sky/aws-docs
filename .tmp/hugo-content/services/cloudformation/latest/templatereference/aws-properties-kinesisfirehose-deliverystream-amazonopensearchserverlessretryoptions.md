This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream AmazonOpenSearchServerlessRetryOptions

Configures retry behavior in case Firehose is unable to deliver documents
to the Serverless offering for Amazon OpenSearch Service.

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

After an initial failure to deliver to the Serverless offering for Amazon OpenSearch
Service, the total amount of time during which Firehose retries delivery
(including the first attempt). After this time has elapsed, the failed documents are
written to Amazon S3. Default value is 300 seconds (5 minutes). A value of 0 (zero) results
in no retries.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `7200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonOpenSearchServerlessDestinationConfiguration

AmazonopensearchserviceBufferingHints

All content copied from https://docs.aws.amazon.com/.
