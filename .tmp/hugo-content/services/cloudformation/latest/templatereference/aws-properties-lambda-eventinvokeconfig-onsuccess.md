This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventInvokeConfig OnSuccess

A destination for events that were processed successfully.

To retain records of successful [asynchronous invocations](../../../lambda/latest/dg/invocation-async.md#invocation-async-destinations),
you can configure an Amazon SNS topic, Amazon SQS queue, Lambda function,
or Amazon EventBridge event bus as the destination.

###### Note

`OnSuccess` is not supported in `CreateEventSourceMapping` or `UpdateEventSourceMapping` requests.

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

###### Note

Amazon SNS destinations have a message size limit of 256 KB. If the combined size of the function request and response payload exceeds the limit, Lambda will drop the payload when sending `OnFailure` event to the destination. For details on this behavior, refer to [Retaining records of asynchronous invocations](../../../lambda/latest/dg/invocation-async-retain-records.md).

_Required_: Yes

_Type_: String

_Pattern_: `^$|arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:([a-z]+(-[a-z]+)+-\d{1})?:(\d{12})?:(.*)`

_Minimum_: `0`

_Maximum_: `350`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### On-Success Destination Configuration

Configure a function to send a record of successful asynchronous invocations to an SQS queue.

#### YAML

```yaml

          OnSuccess:
            Destination: arn:aws:sqs:us-east-2:123456789012:destination
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnFailure

AWS::Lambda::EventSourceMapping

All content copied from https://docs.aws.amazon.com/.
