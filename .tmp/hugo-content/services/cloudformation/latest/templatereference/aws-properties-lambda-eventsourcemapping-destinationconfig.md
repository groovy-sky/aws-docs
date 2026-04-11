This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping DestinationConfig

A configuration object that specifies the destination of an event after Lambda processes it. For more information, see [Adding a destination](../../../lambda/latest/dg/invocation-async-retain-records.md#invocation-async-destinations).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnFailure" : OnFailure
}

```

### YAML

```yaml

  OnFailure:
    OnFailure

```

## Properties

`OnFailure`

The destination configuration for failed invocations.

_Required_: No

_Type_: [OnFailure](aws-properties-lambda-eventsourcemapping-onfailure.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### On-Failure Destination Configuration

Configure a function to send a record of failed batches to an SQS queue.

#### YAML

```yaml

      DestinationConfig:
          OnFailure:
            Destination: arn:aws:sqs:us-east-2:123456789012:dlq
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonManagedKafkaEventSourceConfig

DocumentDBEventSourceConfig

All content copied from https://docs.aws.amazon.com/.
