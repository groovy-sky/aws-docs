This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetCloudWatchLogsParameters

The parameters for using an CloudWatch Logs log stream as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogStreamName" : String,
  "Timestamp" : String
}

```

### YAML

```yaml

  LogStreamName: String
  Timestamp: String

```

## Properties

`LogStreamName`

The name of the log stream.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

A [dynamic path parameter](../../../eventbridge/latest/userguide/eb-pipes-event-target.md) to a field in the payload containing the time the event
occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.

The value cannot be a static timestamp as the provided timestamp would be applied to all
events delivered by the Pipe, regardless of when they are actually delivered.

If no dynamic path parameter is provided, the default value is the time the invocation is
processed by the Pipe.

_Required_: No

_Type_: String

_Pattern_: `^\$(\.[\w_-]+(\[(\d+|\*)\])*)*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetBatchJobParameters

PipeTargetEcsTaskParameters

All content copied from https://docs.aws.amazon.com/.
