This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function DeadLetterConfig

The [dead-letter queue](../../../lambda/latest/dg/invocation-async-retain-records.md#invocation-dlq) for
failed asynchronous invocations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetArn" : String
}

```

### YAML

```yaml

  TargetArn: String

```

## Properties

`TargetArn`

The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Dead-letter Queue

Add a dead-letter queue to a function.

#### YAML

```yaml

      DeadLetterConfig:
        TargetArn: arn:aws:sqs:us-east-2:123456789012:dlq
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Code

DurableConfig

All content copied from https://docs.aws.amazon.com/.
