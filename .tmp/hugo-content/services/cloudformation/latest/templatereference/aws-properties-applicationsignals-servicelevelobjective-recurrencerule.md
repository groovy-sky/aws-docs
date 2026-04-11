This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective RecurrenceRule

The recurrence rule for the time exclusion window.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String
}

```

### YAML

```yaml

  Expression: String

```

## Properties

`Expression`

The following two rules are supported:

- rate(value unit) - The value must be a positive integer and the unit can be hour\|day\|month.

- cron - An expression which consists of six fields separated by white spaces: (minutes hours day\_of\_month month day\_of\_week year).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoredRequestCountMetric

RequestBasedSli

All content copied from https://docs.aws.amazon.com/.
