This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery TargetConfiguration

Configuration used for writing the output of a query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimestreamConfiguration" : TimestreamConfiguration
}

```

### YAML

```yaml

  TimestreamConfiguration:
    TimestreamConfiguration

```

## Properties

`TimestreamConfiguration`

Configuration needed to write data into the Timestream database and table.

_Required_: Yes

_Type_: [TimestreamConfiguration](aws-properties-timestream-scheduledquery-timestreamconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TimestreamConfiguration

All content copied from https://docs.aws.amazon.com/.
