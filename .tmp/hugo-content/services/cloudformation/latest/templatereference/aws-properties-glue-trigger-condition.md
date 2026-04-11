This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Trigger Condition

Defines a condition under which a trigger fires.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrawlerName" : String,
  "CrawlState" : String,
  "JobName" : String,
  "LogicalOperator" : String,
  "State" : String
}

```

### YAML

```yaml

  CrawlerName: String
  CrawlState: String
  JobName: String
  LogicalOperator: String
  State: String

```

## Properties

`CrawlerName`

The name of the crawler to which this condition applies.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlState`

The state of the crawler to which this condition applies.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobName`

The name of the job whose `JobRuns` this condition applies to, and on which
this trigger waits.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalOperator`

A logical operator.

_Required_: No

_Type_: String

_Allowed values_: `EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The condition state. Currently, the values supported are `SUCCEEDED`,
`STOPPED`, `TIMEOUT`, and `FAILED`.

_Required_: No

_Type_: String

_Allowed values_: `STARTING | RUNNING | STOPPING | STOPPED | SUCCEEDED | FAILED | TIMEOUT | ERROR | WAITING | EXPIRED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Condition Structure](../../../glue/latest/dg/aws-glue-api-jobs-trigger.md#aws-glue-api-jobs-trigger-Condition) in the _AWS Glue Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

EventBatchingCondition

All content copied from https://docs.aws.amazon.com/.
