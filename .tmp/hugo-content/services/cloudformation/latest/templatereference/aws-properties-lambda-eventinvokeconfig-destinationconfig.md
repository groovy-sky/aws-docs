This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventInvokeConfig DestinationConfig

A configuration object that specifies the destination of an event after Lambda processes it. For more information, see [Adding a destination](../../../lambda/latest/dg/invocation-async-retain-records.md#invocation-async-destinations).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnFailure" : OnFailure,
  "OnSuccess" : OnSuccess
}

```

### YAML

```yaml

  OnFailure:
    OnFailure
  OnSuccess:
    OnSuccess

```

## Properties

`OnFailure`

The destination configuration for failed invocations.

###### Note

When using an Amazon SQS queue as a destination, FIFO queues cannot be used.

_Required_: No

_Type_: [OnFailure](aws-properties-lambda-eventinvokeconfig-onfailure.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnSuccess`

The destination configuration for successful invocations.

###### Note

When using an Amazon SQS queue as a destination, FIFO queues cannot be used.

_Required_: No

_Type_: [OnSuccess](aws-properties-lambda-eventinvokeconfig-onsuccess.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### On-Failure Destination Configuration

Configure a function to send a record of failed asynchronous invocations to an SQS queue.

#### YAML

```yaml

      DestinationConfig:
          OnFailure:
            Destination: arn:aws:sqs:us-east-2:123456789012:dlq
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::EventInvokeConfig

OnFailure

All content copied from https://docs.aws.amazon.com/.
