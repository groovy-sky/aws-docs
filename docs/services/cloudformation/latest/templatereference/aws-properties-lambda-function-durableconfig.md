This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function DurableConfig

Configuration settings for [durable functions](https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html), including execution timeout and retention period for execution history.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionTimeout" : Integer,
  "RetentionPeriodInDays" : Integer
}

```

### YAML

```yaml

  ExecutionTimeout: Integer
  RetentionPeriodInDays: Integer

```

## Properties

`ExecutionTimeout`

The maximum time (in seconds) that a durable execution can run before timing out. This timeout applies to the entire durable execution, not individual function invocations.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `31622400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriodInDays`

The number of days to retain execution history after a durable execution completes. After this period, execution history is no longer available through the GetDurableExecutionHistory API.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `90`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeadLetterConfig

Environment
