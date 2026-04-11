This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping OnFailure

A destination for events that failed processing. For more information, see [Adding a destination](../../../lambda/latest/dg/invocation-async-retain-records.md#invocation-async-destinations).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String
}

```

### YAML

```yaml

  Destination: String

```

## Properties

`Destination`

The Amazon Resource Name (ARN) of the destination resource.

To retain records of failed invocations from [Kinesis](../../../lambda/latest/dg/with-kinesis.md), [DynamoDB](../../../lambda/latest/dg/with-ddb.md), [self-managed Apache Kafka](../../../lambda/latest/dg/kafka-on-failure.md), or [Amazon MSK](../../../lambda/latest/dg/kafka-on-failure.md), you can configure an Amazon SNS topic, Amazon SQS queue, Amazon S3 bucket, or Kafka topic as the destination.

###### Note

Amazon SNS destinations have a message size limit of 256 KB. If the combined size of the function request and response payload exceeds the limit, Lambda will drop the payload when sending `OnFailure` event to the destination. For details on this behavior, refer to [Retaining records of asynchronous invocations](../../../lambda/latest/dg/invocation-async-retain-records.md).

To retain records of failed invocations from [Kinesis](../../../lambda/latest/dg/with-kinesis.md),
[DynamoDB](../../../lambda/latest/dg/with-ddb.md), [self-managed Kafka](../../../lambda/latest/dg/with-kafka.md#services-smaa-onfailure-destination) or
[Amazon MSK](../../../lambda/latest/dg/with-msk.md#services-msk-onfailure-destination),
you can configure an Amazon SNS topic, Amazon SQS queue, or Amazon S3 bucket as the destination.

_Required_: No

_Type_: String

_Pattern_: `^$|kafka://([^.]([a-zA-Z0-9\-_.]{0,248}))|arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.*)`

_Minimum_: `12`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### On-Failure Destination Configuration

Configure a function to send a record of failed batches to an SQS queue.

#### YAML

```yaml

          OnFailure:
            Destination: arn:aws:sqs:us-east-2:123456789012:dlq
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricsConfig

ProvisionedPollerConfig

All content copied from https://docs.aws.amazon.com/.
