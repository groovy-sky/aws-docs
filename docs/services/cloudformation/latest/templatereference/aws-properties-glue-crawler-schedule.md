This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler Schedule

A scheduling object using a `cron` statement to schedule an event.

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

A `cron` expression used to specify the schedule. For more information, see
[Time-Based Schedules for\
Jobs and Crawlers](../../../glue/latest/dg/monitor-data-warehouse-schedule.md). For example, to run something every day at 12:15 UTC,
specify `cron(15 12 * * ? *)`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Target

SchemaChangePolicy

All content copied from https://docs.aws.amazon.com/.
