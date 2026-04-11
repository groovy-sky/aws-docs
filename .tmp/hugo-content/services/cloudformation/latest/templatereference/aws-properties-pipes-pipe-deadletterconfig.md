This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe DeadLetterConfig

A `DeadLetterConfig` object that contains information about a dead-letter
queue configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String
}

```

### YAML

```yaml

  Arn: String

```

## Properties

`Arn`

The ARN of the specified target for the dead-letter queue.

For Amazon Kinesis stream and Amazon DynamoDB stream sources, specify
either an Amazon SNS topic or Amazon SQS queue ARN.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudwatchLogsLogDestination

DimensionMapping

All content copied from https://docs.aws.amazon.com/.
