This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventInvokeConfig DestinationConfig

A configuration object that specifies the destination of an event after Lambda processes it. For more information, see [Adding a destination](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations).

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

_Type_: [OnFailure](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventinvokeconfig-onfailure.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnSuccess`

The destination configuration for successful invocations.

###### Note

When using an Amazon SQS queue as a destination, FIFO queues cannot be used.

_Required_: No

_Type_: [OnSuccess](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventinvokeconfig-onsuccess.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lambda::EventInvokeConfig

OnFailure
