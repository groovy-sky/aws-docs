This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery ScheduleConfiguration

Configuration of the schedule of the query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScheduleExpression" : String
}

```

### YAML

```yaml

  ScheduleExpression: String

```

## Properties

`ScheduleExpression`

An expression that denotes when to trigger the scheduled query run. This can be a cron
expression or a rate expression.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Configuration

SnsConfiguration
